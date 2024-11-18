package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"

	"github.com/gorilla/mux"
)

// APIHandler holds the necessary dependencies for API endpoints
type APIHandler struct {
	StoragePort ports.StoragePort
	Router      *mux.Router
}

// NewAPIAdapter initializes the APIHandler with a logger and sets up routes
func NewAPIAdapter(storagePort ports.StoragePort) *APIHandler {

	h := &APIHandler{
		StoragePort: storagePort,
		Router:      mux.NewRouter(),
	}

	h.SetupRoutes()
	return h
}

// SetupRoutes sets up all the routes for the API
func (h *APIHandler) SetupRoutes() {
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.UpdateTelegramConfig).Methods("POST")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.UpdateOperatorID).Methods("POST")
	h.Router.HandleFunc("/api/v0/event_indexer/operator_performance", h.GetOperatorPerformance).Methods("GET")
	h.Router.HandleFunc("/api/v0/event_indexer/exit_requests", h.GetExitRequests).Methods("GET")
}

// UpdateTelegramConfig handles updates to the Telegram configuration
func (h *APIHandler) UpdateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string `json:"token"`
		UserID int64  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid request body in UpdateTelegramConfig: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		log.Println("Missing token and userId in UpdateTelegramConfig request")
		writeErrorResponse(w, "At least one of token or userId must be provided", http.StatusBadRequest)
		return
	}

	if err := h.StoragePort.SaveTelegramConfig(domain.TelegramConfig(req)); err != nil {
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateOperatorID handles updates to the operator ID
func (h *APIHandler) UpdateOperatorID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid request body in UpdateOperatorID: %v", err)
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OperatorID == "" {
		log.Println("Missing operatorId in UpdateOperatorID request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	if _, ok := new(big.Int).SetString(req.OperatorID, 10); !ok {
		log.Println("Invalid operatorId format in UpdateOperatorID")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	if err := h.StoragePort.SaveOperatorId(req.OperatorID); err != nil {
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetOperatorPerformance retrieves operator performance data
func (h *APIHandler) GetOperatorPerformance(w http.ResponseWriter, r *http.Request) {
	operatorId := r.URL.Query().Get("operatorId")

	operatorIdNum := new(big.Int)
	if _, ok := operatorIdNum.SetString(operatorId, 10); !ok {
		log.Println("Invalid operatorId in GetOperatorPerformance request")
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	report, err := h.StoragePort.GetReports(operatorIdNum)
	if err != nil {
		writeErrorResponse(w, "Error fetching Lido report", http.StatusInternalServerError)
		return
	}

	if len(report) == 0 {
		log.Printf("No report found for operator ID %s", operatorId)
		writeErrorResponse(w, "No report found for the given epoch range", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		log.Printf("Error generating JSON response in GetOperatorPerformance: %v", err)
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// GetExitRequests retrieves exit requests for a given operator ID
func (h *APIHandler) GetExitRequests(w http.ResponseWriter, r *http.Request) {
	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		log.Println("Missing operatorId in GetExitRequests request")
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	exitRequests, err := h.StoragePort.GetExitRequests(operatorId)
	if err != nil {
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError)
		return
	}

	if exitRequests == nil {
		log.Printf("No exit requests found for operator ID %s", operatorId)
		writeErrorResponse(w, "No exit requests found for the given operator ID", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
		log.Printf("Error generating JSON response in GetExitRequests: %v", err)
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
