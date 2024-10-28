package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/services"

	"github.com/gorilla/mux"
)

type APIHandler struct {
	StorageService *services.StorageService
	Notifier       *services.NotifierService
	Router         *mux.Router
}

// NewAPIAdapter initializes the APIHandler and sets up the routes
func NewAPIAdapter(storageService *services.StorageService, notifier *services.NotifierService) *APIHandler {
	h := &APIHandler{
		StorageService: storageService,
		Notifier:       notifier,
		Router:         mux.NewRouter(),
	}

	h.SetupRoutes()
	return h
}

// SetupRoutes sets up all the routes for the API
func (h *APIHandler) SetupRoutes() {
	h.Router.HandleFunc("/api/v0/events_indexer/telegramConfig", h.UpdateTelegramConfig).Methods("POST")
	h.Router.HandleFunc("/api/v0/events_indexer/operatorId", h.UpdateOperatorID).Methods("POST")
	h.Router.HandleFunc("/api/v0/event_indexer/lido_report", h.GetLidoReport).Methods("GET")
	h.Router.HandleFunc("/api/v0/event_indexer/exit_requests", h.GetExitRequests).Methods("GET")
}

// Handler to update the Telegram token
func (h *APIHandler) UpdateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token  string `json:"token"`
		ChatID int64  `json:"chatId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Token == "" && req.ChatID == 0 {
		writeErrorResponse(w, "At least one of token or chatId must be provided", http.StatusBadRequest)
		return
	}

	err := h.StorageService.SetTelegramConfig(domain.TelegramConfig(req))
	if err != nil {
		writeErrorResponse(w, "Failed to update Telegram configuration", http.StatusInternalServerError)
		return
	}

	err = h.Notifier.SendNotification("Telegram configuration updated")
	if err != nil {
		writeErrorResponse(w, "Failed to send notification", http.StatusInternalServerError)
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

	// Convert the string to a big.Int
	operatorIdNum := new(big.Int)
	_, ok := operatorIdNum.SetString(req.OperatorID, 10)
	if !ok {
		fmt.Println("Error: invalid number string")
		return
	}
	err := h.StorageService.SetOperatorId(domain.OperatorId(operatorIdNum))
	if err != nil {
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError)
		return
	}

	err = h.Notifier.SendNotification("Operator ID updated")
	if err != nil {
		writeErrorResponse(w, "Failed to send notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler to get Lido report within a range of epochs
func (h *APIHandler) GetLidoReport(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

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

	report, err := h.StorageService.GetLidoReport(strconv.Itoa(startEpoch), strconv.Itoa(endEpoch))
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
	exitRequests, err := h.StorageService.GetExitRequests()
	if err != nil {
		writeErrorResponse(w, "Error fetching exit requests", http.StatusInternalServerError)
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
