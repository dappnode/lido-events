package api

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIHandler holds the necessary dependencies for API endpoints
type APIHandler struct {
	ctx               context.Context
	StoragePort       ports.StoragePort
	NotifierPort      ports.NotifierPort
	RelaysUsedPort    ports.RelaysUsedPort
	RelaysAllowedPort ports.RelaysAllowedPort
	Router            *mux.Router
	adapterPrefix     string
}

// Ensure APIHandler implements the ports.API interface
var _ ports.API = (*APIHandler)(nil)

// GetRouter implements the ports.API interface
func (h *APIHandler) GetRouter() http.Handler {
	return h.Router
}

// NewAPIAdapter initializes the APIHandler and sets up routes with CORS enabled
func NewAPIAdapter(ctx context.Context, storagePort ports.StoragePort, notifierPort ports.NotifierPort, relaysUsedPort ports.RelaysUsedPort, relaysAllowedPort ports.RelaysAllowedPort, allowedOrigins []string) *APIHandler {
	h := &APIHandler{
		ctx:               ctx,
		StoragePort:       storagePort,
		NotifierPort:      notifierPort,
		RelaysUsedPort:    relaysUsedPort,
		RelaysAllowedPort: relaysAllowedPort,
		Router:            mux.NewRouter(),
		adapterPrefix:     "API",
	}

	// Set up API routes
	h.SetupRoutes()

	// Define CORS configuration
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Add CORS middleware globally
	h.Router.Use(handlers.CORS(
		corsAllowedOrigins,
		corsAllowedMethods,
		corsAllowedHeaders,
	))

	return h
}

// SetupRoutes sets up all the routes for the API
func (h *APIHandler) SetupRoutes() {
	// Define API endpoints
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.UpdateTelegramConfig).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.GetTelegramConfig).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.AddOperator).Methods("POST", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.DeleteOperator).Methods("DELETE", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/operator_performance", h.GetOperatorPerformance).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/exit_requests", h.GetExitRequests).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/relays_allowed", h.GetRelaysAllowed).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/relays_used", h.GetRelaysUsed).Methods("GET", "OPTIONS")
	h.Router.HandleFunc("/api/v0/events_indexer/node_operator_events", h.GetNodeOperatorEvents).Methods("GET", "OPTIONS")

	// Add a generic OPTIONS handler to ensure preflight requests are handled
	h.Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond to OPTIONS requests with 200 OK
	})
}

// GetNodeOperatorEvents retrieves node operator events
func (h *APIHandler) GetNodeOperatorEvents(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetNodeOperatorEvents request received")
	operatorId := r.URL.Query().Get("operatorId")

	if operatorId == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in GetNodeOperatorEvents request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	// returns
	nodeOperatorEvents, err := h.StoragePort.GetNodeOperatorEvents(operatorId)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching node operator events: %v", err)
		writeErrorResponse(w, "Error fetching node operator events", http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := json.Marshal(nodeOperatorEvents)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetNodeOperatorEvents: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// GetRelaysAllowed retrieves the list of allowed relays
func (h *APIHandler) GetRelaysAllowed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetRelaysAllowed request received")
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

// GetRelaysUsed retrieves the list of used relays
func (h *APIHandler) GetRelaysUsed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetRelaysUsed request received")
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

// GetTelegramConfig retrieves the Telegram configuration
func (h *APIHandler) GetTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetTelegramConfig request received")
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

// UpdateTelegramConfig handles updates to the Telegram configuration
func (h *APIHandler) UpdateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "UpdateTelegramConfig request received")
	var req struct {
		Token  string `json:"token"`
		UserID int64  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.WarnWithPrefix("API", "Invalid request body in UpdateTelegramConfig: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		logger.DebugWithPrefix("API", "Missing token and userId in UpdateTelegramConfig request")
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

// DeleteOperator handles deletion of an operator ID
func (h *APIHandler) DeleteOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "DeleteOperator request received")
	operatorId := r.URL.Query().Get("operatorId")

	if operatorId == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in DeleteOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(operatorId, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in DeleteOperator")
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

// AddOperator handles updates to the operator ID
func (h *APIHandler) AddOperator(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "AddOperator request received")
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.DebugWithPrefix("API", "Invalid request body in AddOperator: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	if req.OperatorID == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in AddOperator request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest, nil)
		return
	}

	if _, ok := new(big.Int).SetString(req.OperatorID, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in AddOperator")
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
	// TODO: Consider moving this logic to the services layer
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
	if err := h.StoragePort.SaveCsModuletLastProcessedBlock(0); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update CsModuleLastProcessedBlock: %v", err)
		writeErrorResponse(w, "Failed to reset CsModuleLastProcessedBlock", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetOperatorPerformance retrieves operator performance data
func (h *APIHandler) GetOperatorPerformance(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetOperatorPerformance request received")
	operatorId := r.URL.Query().Get("operatorId")

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in GetOperatorPerformance")
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
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetOperatorPerformance: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// GetExitRequests retrieves exit requests for a given operator ID
func (h *APIHandler) GetExitRequests(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetExitRequests request received")
	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in GetExitRequests request")
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
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetExitRequests: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Utility function for consistent error responses
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err != nil {
		message += ": " + err.Error()
	}
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
