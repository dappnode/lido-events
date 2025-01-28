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
	server                *http.Server
	servicePrefix         string
	ctx                   context.Context
	StoragePort           ports.StoragePort
	NotifierPort          ports.NotifierPort
	RelaysUsedPort        ports.RelaysUsedPort
	RelaysAllowedPort     ports.RelaysAllowedPort
	CsModuleEventsScanner *CsModuleEventsScanner
	Router                *mux.Router
	processingAddresses   sync.Map // To track addresses being processed
}

// NewAPIServerService initializes the API server
func NewAPIServerService(ctx context.Context, port uint64, storagePort ports.StoragePort, notifierPort ports.NotifierPort, relaysUsedPort ports.RelaysUsedPort, relaysAllowedPort ports.RelaysAllowedPort, CsModuleEventsScanner *CsModuleEventsScanner, allowedOrigins []string) *APIServerService {
	router := mux.NewRouter()
	apiServer := &APIServerService{
		server:                &http.Server{Addr: ":" + strconv.FormatUint(port, 10)},
		servicePrefix:         "API",
		ctx:                   ctx,
		StoragePort:           storagePort,
		NotifierPort:          notifierPort,
		RelaysUsedPort:        relaysUsedPort,
		RelaysAllowedPort:     relaysAllowedPort,
		CsModuleEventsScanner: CsModuleEventsScanner,
		Router:                router,
	}

	apiServer.SetupRoutes()
	apiServer.server.Handler = apiServer.Router

	// Define CORS configuration
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Add CORS middleware globally
	apiServer.Router.Use(handlers.CORS(
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
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.updateTelegramConfig).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.getTelegramConfig).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.addOperator).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.deleteOperator).Methods("DELETE", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operator_performance", h.getOperatorPerformance).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/exit_requests", h.getExitRequests).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/relays_allowed", h.getRelaysAllowed).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/relays_used", h.getRelaysUsed).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/address_events", h.getAddressEvents).Methods("GET", "OPTIONS")

	// Add a generic OPTIONS handler to ensure preflight requests are handled
	h.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond to OPTIONS requests with 200 OK
	})
}

// Handlers

// getAddressEvents retrieves events for a given address
func (h *APIServerService) getAddressEvents(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getAddressEvents request received")

	address := r.URL.Query().Get("address")
	if address == "" {
		writeErrorResponse(w, "address is required", http.StatusBadRequest, nil)
		return
	}

	// Validate address
	if !common.IsHexAddress(address) {
		logger.ErrorWithPrefix("API", "Invalid address format in getAddressEvents: %s", address)
		writeErrorResponse(w, "Invalid address format", http.StatusBadRequest, nil)
		return
	}
	addressValidated := common.HexToAddress(address)

	// Check if the address is already being processed
	if _, exists := h.processingAddresses.Load(addressValidated.Hex()); exists {
		logger.DebugWithPrefix("API", "Address %s is already being processed", addressValidated.Hex())
		// If processing, return a 202 response
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Request is being processed"))
		return
	}

	// Mark the address as being processed
	h.processingAddresses.Store(addressValidated.Hex(), struct{}{})
	defer h.processingAddresses.Delete(addressValidated.Hex()) // Clear address when done

	// Perform the scanning (synchronously)
	if err := h.CsModuleEventsScanner.ScanAddressEvents(h.ctx, addressValidated); err != nil {
		logger.ErrorWithPrefix("API", "Error scanning address events: %v", err)
		writeErrorResponse(w, "Error scanning address events", http.StatusInternalServerError, err)
		return
	}

	// Fetch the updated events for the address
	addressEvents, err := h.StoragePort.GetAddressEvents(addressValidated)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching address events: %v", err)
		writeErrorResponse(w, "Error fetching address events", http.StatusInternalServerError, err)
		return
	}

	// Respond with the scanned and fetched data
	jsonResponse, err := json.Marshal(addressEvents)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in getAddressEvents: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// getRelaysAllowed retrieves the list of allowed relays
func (h *APIServerService) getRelaysAllowed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getRelaysAllowed request received")
	relays, err := h.RelaysAllowedPort.GetRelaysAllowList(h.ctx)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching allowed relays: %v", err)
		writeErrorResponse(w, "Error fetching allowed relays", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(relays)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetRelaysAllowed: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getRelaysUsed retrieves the list of used relays
func (h *APIServerService) getRelaysUsed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getRelaysUsed request received")
	relays, err := h.RelaysUsedPort.GetRelaysUsed(h.ctx)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching used relays: %v", err)
		writeErrorResponse(w, "Error fetching used relays", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(relays)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetRelaysUsed: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getTelegramConfig retrieves the Telegram configuration
func (h *APIServerService) getTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getTelegramConfig request received")
	config, err := h.StoragePort.GetTelegramConfig()
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching Telegram configuration: %v", err)
		writeErrorResponse(w, "Error fetching Telegram configuration", http.StatusInternalServerError, err)
		return
	}

	if config.Token == "" && config.UserID == 0 {
		logger.DebugWithPrefix("API", "No Telegram configuration found")
		writeErrorResponse(w, "No Telegram configuration found", http.StatusNotFound, err)
		return
	}

	jsonResponse, err := json.Marshal(config)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetTelegramConfig: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// updateTelegramConfig handles updates to the Telegram configuration
func (h *APIServerService) updateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "updateTelegramConfig request received")
	var req struct {
		Token  string `json:"token"`
		UserID int64  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.WarnWithPrefix("API", "Invalid request body in updateTelegramConfig: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		logger.DebugWithPrefix("API", "Missing token and userId in updateTelegramConfig request")
		writeErrorResponse(w, "At least one of token or userId must be provided", http.StatusBadRequest, nil)
		return
	}

	// Update storage
	if err := h.StoragePort.SaveTelegramConfig(domain.TelegramConfig(req)); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update Telegram configuration: %v", err)
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError, err)
		return
	}

	// Synchronously update the Telegram bot configuration
	if err := h.NotifierPort.UpdateBotConfig(); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update Telegram bot configuration: %v", err)
		writeErrorResponse(w, "Failed to update Telegram bot configuration", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// deleteOperator handles deletion of an operator ID
func (h *APIServerService) deleteOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "deleteOperator request received")
	operatorId := r.URL.Query().Get("operatorId")

	if operatorId == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in deleteOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(operatorId, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in deleteOperator")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if operator ID exists
	operatorIds, err := h.StoragePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix("API", "Failed to fetch operator IDs: %v", err)
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
		logger.DebugWithPrefix("API", "Operator ID %s not found", operatorId)
		writeErrorResponse(w, "Operator ID not found", http.StatusNotFound, err)
		return
	}

	if err := h.StoragePort.DeleteOperator(operatorId); err != nil {
		logger.ErrorWithPrefix("API", "Failed to delete Operator ID: %v", err)
		writeErrorResponse(w, "Failed to delete Operator ID", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// addOperator handles updates to the operator ID
func (h *APIServerService) addOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "addOperator request received")
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.DebugWithPrefix("API", "Invalid request body in addOperator: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.OperatorID == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in addOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(req.OperatorID, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in addOperator")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	// Check if operator ID already exists
	operatorIds, err := h.StoragePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix("API", "Failed to fetch operator IDs: %v", err)
		writeErrorResponse(w, "Failed to fetch operator IDs", http.StatusInternalServerError, err)
		return
	}

	for _, id := range operatorIds {
		if id.String() == req.OperatorID {
			logger.DebugWithPrefix("API", "Operator ID %s already exists", req.OperatorID)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	if err := h.StoragePort.SaveOperatorId(req.OperatorID); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update Operator ID: %v", err)
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError, err)
		return
	}

	// Set last block processed to 0, this will trigger the events scanner to start from the beginning
	// and look for events for the new operator ID
	if err := h.StoragePort.SaveDistributionLogLastProcessedBlock(0); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update DistributionLogLastProcessedBlock: %v", err)
		writeErrorResponse(w, "Failed to reset DistributionLogLastProcessedBlock", http.StatusInternalServerError, err)
		return
	}
	if err := h.StoragePort.SaveValidatorExitRequestLastProcessedBlock(0); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update ValidatorExitRequestLastProcessedBlock: %v", err)
		writeErrorResponse(w, "Failed to reset ValidatorExitRequestLastProcessedBlock", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// getOperatorPerformance retrieves operator performance data
func (h *APIServerService) getOperatorPerformance(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getOperatorPerformance request received")
	operatorId := r.URL.Query().Get("operatorId")

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in getOperatorPerformance")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest, nil)
		return
	}

	report, err := h.StoragePort.GetReports(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching Lido report: %v", err)
		writeErrorResponse(w, "Error fetching Lido report", http.StatusInternalServerError, err)
		return
	}

	if len(report) == 0 {
		logger.WarnWithPrefix("API", "No report found for operator ID %s", operatorId)
		writeErrorResponse(w, "No report found for the given epoch range", http.StatusNotFound, err)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in getOperatorPerformance: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getExitRequests retrieves exit requests for a given operator ID
func (h *APIServerService) getExitRequests(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "getExitRequests request received")
	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in getExitRequests request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	exitRequests, err := h.StoragePort.GetExitRequests(operatorId)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching exit requests: %v", err)
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError, err)
		return
	}

	if exitRequests == nil {
		logger.DebugWithPrefix("API", "No exit requests found for operator ID %s", operatorId)
		writeErrorResponse(w, "No exit requests found for the given operator ID", http.StatusNotFound, err)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in getExitRequests: %v", err)
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
