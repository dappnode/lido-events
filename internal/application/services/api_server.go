package services

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"sync"

	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIServerService starts and manages the API server
type APIServerService struct {
	server                             *http.Server
	servicePrefix                      string
	ctx                                context.Context
	storagePort                        ports.StoragePort
	notifierPort                       ports.NotifierPort
	relaysUsedPort                     ports.RelaysUsedPort
	relaysAllowedPort                  ports.RelaysAllowedPort
	csFeeOracleEventsScanner           *CsFeeOracleEventsScanner
	csModuleEventsScanner              *CsModuleEventsScanner
	distributionLogUpdatedEventScanner *DistributionLogUpdatedEventScanner
	validatorExitRequestEventScanner   *ValidatorExitRequestEventScanner
	pendingHashesLoader                *PendingHashesLoader
	router                             *mux.Router
	processingAddresses                sync.Map // To track NO events being processed by address
	processingWithdrawals              sync.Map // To track withdrawals being processed by operator ID
	processingExitRequests             sync.Map // To track exit requests being processed by operator ID
	processingOperatorPerformance      sync.Map // To track operator performance being processed by operator ID
	processingProcessingStarted        bool     // bool to track ProcessingStarted events being processed
	processingPenalties                bool     // bool to track EL rewards stealing penalties being processed
}

// NewAPIServerService initializes the API server
func NewAPIServerService(ctx context.Context, port uint64, storagePort ports.StoragePort, notifierPort ports.NotifierPort, relaysUsedPort ports.RelaysUsedPort, relaysAllowedPort ports.RelaysAllowedPort, csFeeOracleEventsScanner *CsFeeOracleEventsScanner, csModuleEventsScanner *CsModuleEventsScanner, distributionLogUpdatedEventScanner *DistributionLogUpdatedEventScanner, validatorExitRequestEventScanner *ValidatorExitRequestEventScanner, pendingHashesLoader *PendingHashesLoader, allowedOrigins []string) *APIServerService {
	apiServer := &APIServerService{
		server:                             &http.Server{Addr: ":" + strconv.FormatUint(port, 10)},
		servicePrefix:                      "API",
		ctx:                                ctx,
		storagePort:                        storagePort,
		notifierPort:                       notifierPort,
		relaysUsedPort:                     relaysUsedPort,
		relaysAllowedPort:                  relaysAllowedPort,
		csFeeOracleEventsScanner:           csFeeOracleEventsScanner,
		csModuleEventsScanner:              csModuleEventsScanner,
		distributionLogUpdatedEventScanner: distributionLogUpdatedEventScanner,
		validatorExitRequestEventScanner:   validatorExitRequestEventScanner,
		pendingHashesLoader:                pendingHashesLoader,
		router:                             mux.NewRouter(),
	}

	apiServer.SetupRoutes()
	apiServer.server.Handler = apiServer.router

	// Define CORS configuration
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Add CORS middleware globally
	apiServer.router.Use(handlers.CORS(
		corsAllowedOrigins,
		corsAllowedMethods,
		corsAllowedHeaders,
	))

	return apiServer
}

func (s *APIServerService) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			logger.FatalWithPrefix(s.servicePrefix, "Server error: %v", err)
		}
	}()
}

func (s *APIServerService) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}

// SetupRoutes defines API endpoints and routes
func (h *APIServerService) SetupRoutes() {
	h.router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.updateTelegramConfig).Methods("POST", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.getTelegramConfig).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/operatorId", h.addOperator).Methods("POST", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/operatorId", h.deleteOperator).Methods("DELETE", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/operator_performance", h.getOperatorPerformance).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/exit_requests", h.getExitRequests).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/relays_allowed", h.getRelaysAllowed).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/relays_used", h.getRelaysUsed).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/address_events", h.getAddressEvents).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/withdrawals_submitted", h.getWithdrawalsSubmitted).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/el_rewards_stealing_penalties_reported", h.getElRewardsStealingPenaltiesReported).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/pending_hashes", h.getPendingHashes).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/processing_started", h.getProcessingStarted).Methods("GET", "OPTIONS")

	// Add a generic OPTIONS handler to ensure preflight requests are handled
	h.router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond to OPTIONS requests with 200 OK
	})
}

// Handlers

// getProcessingStarted retrieves the processing started events
func (h *APIServerService) getProcessingStarted(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getProcessingStarted request received")
	// Check if the processing started events are already being processed
	if h.processingProcessingStarted {
		logger.DebugWithPrefix(h.servicePrefix, "ProcessingStarted events are already being processed")
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Perform the scanning (synchronously)
	if err := h.csFeeOracleEventsScanner.RunScan(h.ctx); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning ProcessingStarted events: %v", err)
		writeErrorResponse(w, "Error scanning ProcessingStarted events", http.StatusInternalServerError, err)
		return
	}

	// Fetch the updated events
	processingStartedEvents, err := h.storagePort.GetProcessingStartedEvents()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching ProcessingStarted events: %v", err)
		writeErrorResponse(w, "Error fetching ProcessingStarted events", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(processingStartedEvents)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getProcessingStarted: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getPendingHashes retrieves the list of pending hashes for a given operator ID
func (h *APIServerService) getPendingHashes(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getPendingHashes request received")

	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		logger.ErrorWithPrefix(h.servicePrefix, "Missing operatorId in getPendingHashes request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.ErrorWithPrefix(h.servicePrefix, "Invalid operatorId format in getPendingHashes")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if the operator ID exists
	operatorIds, err := h.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching operator IDs: %v", err)
		writeErrorResponse(w, "Error fetching operator IDs", http.StatusInternalServerError, err)
		return
	}

	found := false
	for _, id := range operatorIds {
		if id.String() == operatorId {
			found = true
			break
		}
	}

	if !found {
		logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s not found", operatorId)
		writeErrorResponse(w, "Operator ID not found", http.StatusNotFound, err)
		return
	}

	pendingHashes, err := h.storagePort.GetPendingHashes(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching pending hashes: %v", err)
		writeErrorResponse(w, "Error fetching pending hashes", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(pendingHashes)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getPendingHashes: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getElRewardsStealingPenaltiesReported retrieves the list of EL rewards stealing penalties reported
func (h *APIServerService) getElRewardsStealingPenaltiesReported(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getElRewardsStealingPenaltiesReported request received")

	// Check if the penalties are already being processed
	if h.processingPenalties {
		logger.DebugWithPrefix(h.servicePrefix, "EL rewards stealing penalties are already being processed")
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the penalties as being processed
	h.processingPenalties = true
	defer func() {
		h.processingPenalties = false // Clear when done
	}()

	// Perform the scanning (synchronously)
	if err := h.csModuleEventsScanner.ScanElRewardsStealingPenaltyReported(h.ctx); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning EL rewards stealing penalties reported: %v", err)
		writeErrorResponse(w, "Error scanning EL rewards stealing penalties reported", http.StatusInternalServerError, err)
		return
	}

	penalties, err := h.storagePort.GetElRewardsStealingPenaltiesReported()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching EL rewards stealing penalties reported: %v", err)
		writeErrorResponse(w, "Error fetching EL rewards stealing penalties reported", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(penalties)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getElRewardsStealingPenaltiesReported: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getWithdrawalsSubmitted retrieves the list of withdrawals submitted
func (h *APIServerService) getWithdrawalsSubmitted(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getWithdrawalsSubmitted request received")

	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		logger.ErrorWithPrefix(h.servicePrefix, "Missing operatorId in getWithdrawalsSubmitted request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.ErrorWithPrefix(h.servicePrefix, "Invalid operatorId format in getWithdrawalsSubmitted")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if the operator ID is already being processed
	if _, exists := h.processingWithdrawals.Load(operatorIdNum.String()); exists {
		logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s is already being processed", operatorIdNum.String())
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the operator ID as being processed
	h.processingWithdrawals.Store(operatorIdNum.String(), struct{}{})
	defer h.processingWithdrawals.Delete(operatorIdNum.String()) // Clear operator ID when done

	// Perform the scanning (synchronously)
	if err := h.csModuleEventsScanner.ScanWithdrawalsSubmittedEvents(h.ctx, operatorIdNum); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning withdrawals submitted: %v", err)
		writeErrorResponse(w, "Error scanning withdrawals submitted", http.StatusInternalServerError, err)
		return
	}

	withdrawals, err := h.storagePort.GetWithdrawals(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching withdrawals submitted: %v", err)
		writeErrorResponse(w, "Error fetching withdrawals submitted", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(withdrawals)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getWithdrawalsSubmitted: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getAddressEvents retrieves events for a given address
func (h *APIServerService) getAddressEvents(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getAddressEvents request received")

	// Check for optional boolean parameter "force"
	force := r.URL.Query().Get("force")
	// Default value
	forceBool := false
	// Convert to boolean only if force is not empty
	if force != "" {
		var err error
		forceBool, err = strconv.ParseBool(force)
		if err != nil {
			writeErrorResponse(w, "Invalid force parameter", http.StatusBadRequest, err)
			return
		}
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		writeErrorResponse(w, "address is required", http.StatusBadRequest, nil)
		return
	}

	// Validate address
	if !common.IsHexAddress(address) {
		logger.ErrorWithPrefix(h.servicePrefix, "Invalid address format in getAddressEvents: %s", address)
		writeErrorResponse(w, "Invalid address format", http.StatusBadRequest, nil)
		return
	}
	addressValidated := common.HexToAddress(address)

	// Check if the address is already being processed
	if _, exists := h.processingAddresses.Load(addressValidated.Hex()); exists {
		logger.DebugWithPrefix(h.servicePrefix, "Address %s is already being processed", addressValidated.Hex())
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the address as being processed
	h.processingAddresses.Store(addressValidated.Hex(), struct{}{})
	defer h.processingAddresses.Delete(addressValidated.Hex()) // Clear address when done

	// Perform the scanning (synchronously)
	if err := h.csModuleEventsScanner.ScanAddressEvents(h.ctx, addressValidated, forceBool); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning address events: %v", err)
		writeErrorResponse(w, "Error scanning address events", http.StatusInternalServerError, err)
		return
	}

	// Fetch the updated events for the address
	addressEvents, err := h.storagePort.GetAddressEvents(addressValidated)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching address events: %v", err)
		writeErrorResponse(w, "Error fetching address events", http.StatusInternalServerError, err)
		return
	}

	// Respond with the scanned and fetched data
	jsonResponse, err := json.Marshal(addressEvents)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getAddressEvents: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// getRelaysAllowed retrieves the list of allowed relays
func (h *APIServerService) getRelaysAllowed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getRelaysAllowed request received")
	relays, err := h.relaysAllowedPort.GetRelaysAllowList(h.ctx)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching allowed relays: %v", err)
		writeErrorResponse(w, "Error fetching allowed relays", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(relays)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in GetRelaysAllowed: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getRelaysUsed retrieves the list of used relays
func (h *APIServerService) getRelaysUsed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getRelaysUsed request received")
	relays, err := h.relaysUsedPort.GetRelaysUsed(h.ctx)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching used relays: %v", err)
		writeErrorResponse(w, "Error fetching used relays", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(relays)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in GetRelaysUsed: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getTelegramConfig retrieves the Telegram configuration
func (h *APIServerService) getTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getTelegramConfig request received")
	config, err := h.storagePort.GetTelegramConfig()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching Telegram configuration: %v", err)
		writeErrorResponse(w, "Error fetching Telegram configuration", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(config)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in GetTelegramConfig: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// updateTelegramConfig handles updates to the Telegram configuration
func (h *APIServerService) updateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "updateTelegramConfig request received")
	var req struct {
		Token  string `json:"token"`
		UserID int64  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.WarnWithPrefix(h.servicePrefix, "Invalid request body in updateTelegramConfig: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		logger.DebugWithPrefix(h.servicePrefix, "Missing token and userId in updateTelegramConfig request")
		writeErrorResponse(w, "At least one of token or userId must be provided", http.StatusBadRequest, nil)
		return
	}

	// Update storage
	if err := h.storagePort.SaveTelegramConfig(domain.TelegramConfig(req)); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to update Telegram configuration: %v", err)
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError, err)
		return
	}

	// Synchronously update the Telegram bot configuration
	if err := h.notifierPort.UpdateBotConfig(); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to update Telegram bot configuration: %v", err)
		writeErrorResponse(w, "Failed to update Telegram bot configuration", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteOperator handles deletion of an operator ID
func (h *APIServerService) deleteOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "deleteOperator request received")
	operatorId := r.URL.Query().Get("operatorId")

	if operatorId == "" {
		logger.DebugWithPrefix(h.servicePrefix, "Missing operatorId in deleteOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(operatorId, 10); !ok {
		logger.DebugWithPrefix(h.servicePrefix, "Invalid operatorId format in deleteOperator")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if operator ID exists
	operatorIds, err := h.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to fetch operator IDs: %v", err)
		writeErrorResponse(w, "Failed to fetch operator IDs", http.StatusInternalServerError, err)
		return
	}

	found := false
	for _, id := range operatorIds {
		if id.String() == operatorId {
			found = true
			break
		}
	}

	if !found {
		logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s not found", operatorId)
		writeErrorResponse(w, "Operator ID not found", http.StatusNotFound, err)
		return
	}

	if err := h.storagePort.DeleteOperator(operatorId); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to delete Operator ID: %v", err)
		writeErrorResponse(w, "Failed to delete Operator ID", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// addOperator handles updates to the operator ID
func (h *APIServerService) addOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "addOperator request received")
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.DebugWithPrefix(h.servicePrefix, "Invalid request body in addOperator: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.OperatorID == "" {
		logger.DebugWithPrefix(h.servicePrefix, "Missing operatorId in addOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(req.OperatorID, 10); !ok {
		logger.DebugWithPrefix(h.servicePrefix, "Invalid operatorId format in addOperator")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if operator ID already exists
	operatorIds, err := h.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to fetch operator IDs: %v", err)
		writeErrorResponse(w, "Failed to fetch operator IDs", http.StatusInternalServerError, err)
		return
	}

	for _, id := range operatorIds {
		if id.String() == req.OperatorID {
			logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s already exists", req.OperatorID)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if err := h.storagePort.SaveOperatorId(req.OperatorID); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to update Operator ID: %v", err)
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// getOperatorPerformance retrieves operator performance data
func (h *APIServerService) getOperatorPerformance(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getOperatorPerformance request received")
	operatorId := r.URL.Query().Get("operatorId")

	if operatorId == "" {
		logger.ErrorWithPrefix(h.servicePrefix, "Missing operatorId in getOperatorPerformance request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.ErrorWithPrefix(h.servicePrefix, "Invalid operatorId format in getOperatorPerformance")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if the operator ID is already being processed
	if _, exists := h.processingOperatorPerformance.Load(operatorIdNum.String()); exists {
		logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s is already being processed", operatorIdNum.String())
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the operator ID as being processed
	h.processingOperatorPerformance.Store(operatorIdNum.String(), struct{}{})
	defer h.processingOperatorPerformance.Delete(operatorIdNum.String()) // Clear operator ID when done

	// Perform distribution log scanning (synchronously)
	if err := h.distributionLogUpdatedEventScanner.RunScan(h.ctx); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning operator performance events: %v", err)
		writeErrorResponse(w, "Error scanning operator performance events", http.StatusInternalServerError, err)
		return
	}

	// Load the pending hashes
	if err := h.pendingHashesLoader.LoadPendingHashes(true); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error loading pending hashes: %v", err)
		writeErrorResponse(w, "Error loading pending hashes", http.StatusInternalServerError, err)
		return
	}

	// Fetch operator performance data
	report, err := h.storagePort.GetReports(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching Lido report: %v", err)
		writeErrorResponse(w, "Error fetching Lido report", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getOperatorPerformance: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getExitRequests retrieves exit requests for a given operator ID
func (h *APIServerService) getExitRequests(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getExitRequests request received")

	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		logger.ErrorWithPrefix(h.servicePrefix, "Missing operatorId in getExitRequests request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.ErrorWithPrefix(h.servicePrefix, "Invalid operatorId format in getExitRequests")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if the operator ID is already being processed
	if _, exists := h.processingExitRequests.Load(operatorIdNum.String()); exists {
		logger.DebugWithPrefix(h.servicePrefix, "Operator ID %s is already being processed", operatorIdNum.String())
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the operator ID as being processed
	h.processingExitRequests.Store(operatorIdNum.String(), struct{}{})
	defer h.processingExitRequests.Delete(operatorIdNum.String()) // Clear operator ID when done

	// Perform the scanning (synchronously)
	if err := h.validatorExitRequestEventScanner.RunScan(h.ctx); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error scanning ValidatorExitRequest events: %v", err)
		writeErrorResponse(w, "Error scanning ValidatorExitRequest events", http.StatusInternalServerError, err)
		return
	}

	exitRequests, err := h.storagePort.GetExitRequests(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error fetching exit requests: %v", err)
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Error generating JSON response in getExitRequests: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Utility function for error responses
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err != nil {
		message += ": " + err.Error()
	}
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
