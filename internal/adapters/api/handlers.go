package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"lido-events/internal/services"
)

type APIHandler struct {
	IndexerService *services.Service
}

// Handler to update the Telegram token
func (h *APIHandler) UpdateTelegramToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TelegramToken string `json:"telegramToken"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.IndexerService.UpdateTelegramToken(req.TelegramToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	err := h.IndexerService.UpdateOperatorID(req.OperatorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler to get Lido report within a range of epochs
func (h *APIHandler) GetLidoReport(w http.ResponseWriter, r *http.Request) {
	startEpoch := r.URL.Query().Get("start")
	endEpoch := r.URL.Query().Get("end")

	start, err := strconv.Atoi(startEpoch)
	if err != nil {
		http.Error(w, "Invalid start epoch", http.StatusBadRequest)
		return
	}

	end, err := strconv.Atoi(endEpoch)
	if err != nil {
		http.Error(w, "Invalid end epoch", http.StatusBadRequest)
		return
	}

	report, err := h.IndexerService.GetLidoReport(start, end)
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
	exitRequests, err := h.IndexerService.GetExitRequests()
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
