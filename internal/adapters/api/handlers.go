package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"lido-events/internal/domain/entities"
	service "lido-events/internal/services"
)

type APIHandler struct {
	StorageService *service.StorageService
	Notifier       *service.NotifierService
}

// Handler to update the Telegram token
func (h *APIHandler) UpdateTelegramConfig(w http.ResponseWriter, r *http.Request) {
	// token or chatId. Only 1 is required at least
	var req struct {
		Token  string `json:"token"`
		ChatID int64  `json:"chatId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.StorageService.SetTelegramConfig(entities.TelegramConfig(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.Notifier.Send("Telegram configuration updated")
	w.WriteHeader(http.StatusOK)
}

// Handler to update the node operator ID
func (h *APIHandler) UpdateOperatorID(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OperatorID string `json:"operatorId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convierte req.OperatorID de string a entities.OperatorId
	err := h.StorageService.SetOperatorId(entities.OperatorId(req.OperatorID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Notifier.Send("Operator ID updated")
	w.WriteHeader(http.StatusOK)
}

// Handler to get Lido report within a range of epochs
func (h *APIHandler) GetLidoReport(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	_, err := strconv.Atoi(start)
	if err != nil {
		http.Error(w, "Invalid start epoch", http.StatusBadRequest)
		return
	}

	_, err = strconv.Atoi(end)
	if err != nil {
		http.Error(w, "Invalid end epoch", http.StatusBadRequest)
		return
	}

	report, err := h.StorageService.GetLidoReport(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// Handler to get validator exit requests
func (h *APIHandler) GetExitRequests(w http.ResponseWriter, r *http.Request) {
	exitRequests, err := h.StorageService.GetExitRequests()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(exitRequests)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
