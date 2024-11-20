package api

import (
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
	StoragePort   ports.StoragePort
	Router        *mux.Router
	adapterPrefix string
}

// NewAPIAdapter initializes the APIHandler and sets up routes with CORS enabled
func NewAPIAdapter(storagePort ports.StoragePort, allowedOrigins []string) *APIHandler {
	h := &APIHandler{
		storagePort,
		mux.NewRouter(),
		"API",
	}

	h.SetupRoutes()

	// Add CORS middleware to the router using the provided allowed origins
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	h.Router.Use(func(next http.Handler) http.Handler {
		return handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(next)
	})

	return h
}

// SetupRoutes sets up all the routes for the API
func (h *APIHandler) SetupRoutes() {
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.UpdateTelegramConfig).Methods("POST")
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.GetTelegramConfig).Methods("GET")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.UpdateOperatorID).Methods("POST")
	h.Router.HandleFunc("/api/v0/events_indexer/operator_performance", h.GetOperatorPerformance).Methods("GET")
	h.Router.HandleFunc("/api/v0/events_indexer/exit_requests", h.GetExitRequests).Methods("GET")
}

// GetTelegramConfig retrieves the Telegram configuration
func (h *APIHandler) GetTelegramConfig(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "GetTelegramConfig request received")
	config, err := h.StoragePort.GetTelegramConfig()
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching Telegram configuration: %v", err)
		writeErrorResponse(w, "Error fetching Telegram configuration", http.StatusInternalServerError)
		return
	}

	if config.Token == "" && config.UserID == 0 {
		logger.DebugWithPrefix("API", "No Telegram configuration found")
		writeErrorResponse(w, "No Telegram configuration found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(config)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetTelegramConfig: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError)
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
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		logger.DebugWithPrefix("API", "Missing token and userId in UpdateTelegramConfig request")
		writeErrorResponse(w, "At least one of token or userId must be provided", http.StatusBadRequest)
		return
	}

	if err := h.StoragePort.SaveTelegramConfig(domain.TelegramConfig(req)); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update Telegram configuration: %v", err)
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateOperatorID handles updates to the operator ID
func (h *APIHandler) UpdateOperatorID(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix("API", "UpdateOperatorID request received")
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.DebugWithPrefix("API", "Invalid request body in UpdateOperatorID: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OperatorID == "" {
		logger.DebugWithPrefix("API", "Missing operatorId in UpdateOperatorID request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	if _, ok := new(big.Int).SetString(req.OperatorID, 10); !ok {
		logger.DebugWithPrefix("API", "Invalid operatorId format in UpdateOperatorID")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	if err := h.StoragePort.SaveOperatorId(req.OperatorID); err != nil {
		logger.ErrorWithPrefix("API", "Failed to update Operator ID: %v", err)
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError)
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
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	report, err := h.StoragePort.GetReports(operatorIdNum)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching Lido report: %v", err)
		writeErrorResponse(w, "Error fetching Lido report", http.StatusInternalServerError)
		return
	}

	if len(report) == 0 {
		logger.WarnWithPrefix("API", "No report found for operator ID %s", operatorId)
		writeErrorResponse(w, "No report found for the given epoch range", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetOperatorPerformance: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError)
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
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	exitRequests, err := h.StoragePort.GetExitRequests(operatorId)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error fetching exit requests: %v", err)
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError)
		return
	}

	if exitRequests == nil {
		logger.DebugWithPrefix("API", "No exit requests found for operator ID %s", operatorId)
		writeErrorResponse(w, "No exit requests found for the given operator ID", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
		logger.ErrorWithPrefix("API", "Error generating JSON response in GetExitRequests: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Utility function for consistent error responses
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
