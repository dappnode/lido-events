package api

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"

	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"

	"github.com/gorilla/mux"
)

// TODO: add cors middleware

type APIHandler struct {
	StoragePort ports.StoragePort
	Router      *mux.Router
}

// NewAPIAdapter initializes the APIHandler and sets up the routes
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

// Handler to update the Telegram token
func (h *APIHandler) UpdateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string `json:"token"`
		UserID int64  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Token == "" && req.UserID == 0 {
		writeErrorResponse(w, "At least one of token or userId must be provided", http.StatusBadRequest)
		return
	}

	err := h.StoragePort.SaveTelegramConfig(domain.TelegramConfig(req))
	if err != nil {
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler to update the node operator ID
func (h *APIHandler) UpdateOperatorID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.OperatorID == "" {
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	// Check if the operator ID is a valid number
	_, ok := new(big.Int).SetString(req.OperatorID, 10)
	if !ok {
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	// Update the operator ID in the storage service
	err := h.StoragePort.SaveOperatorId(req.OperatorID)
	if err != nil {
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler to get Lido report within a range of epochs
func (h *APIHandler) GetOperatorPerformance(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	operatorId := r.URL.Query().Get("operatorId")

	operatorIdNum := new(big.Int)
	_, ok := operatorIdNum.SetString(operatorId, 10)
	if !ok {
		writeErrorResponse(w, "Invalid operator ID", http.StatusBadRequest)
		return
	}

	startEpoch, err := strconv.Atoi(start)
	if err != nil {
		writeErrorResponse(w, "Invalid start epoch", http.StatusBadRequest)
		return
	}

	endEpoch, err := strconv.Atoi(end)
	if err != nil {
		writeErrorResponse(w, "Invalid end epoch", http.StatusBadRequest)
		return
	}

	report, err := h.StoragePort.GetOperatorPerformance(operatorIdNum, strconv.Itoa(startEpoch), strconv.Itoa(endEpoch))
	if err != nil {
		writeErrorResponse(w, "Error fetching Lido report", http.StatusInternalServerError)
		return
	}

	if len(report) == 0 {
		writeErrorResponse(w, "No report found for the given epoch range", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		writeErrorResponse(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Handler to get validator exit requests
func (h *APIHandler) GetExitRequests(w http.ResponseWriter, r *http.Request) {
	operatorId := r.URL.Query().Get("operatorId")
	if operatorId == "" {
		writeErrorResponse(w, "operatorId is required", http.StatusBadRequest)
		return
	}

	exitRequests, err := h.StoragePort.GetExitRequests(operatorId)
	if err != nil {
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError)
		return
	}

	// not found
	if exitRequests == nil {
		writeErrorResponse(w, "No exit requests found for the given operator ID", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
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
