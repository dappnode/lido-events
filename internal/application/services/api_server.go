package services

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"sync"

	"lido-events/internal/application/ports"
	"lido-events/internal/logger"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIServerService starts and manages the API server
type APIServerService struct {
	server                           *http.Server
	servicePrefix                    string
	ctx                              context.Context
	exitsStorage                     ports.ExitsStorage
	performanceStorage               ports.PerformanceStorage
	relays                           ports.Relays
	validatorExitRequestEventScanner *ValidatorExitRequestEventScanner
	router                           *mux.Router
}

// NewAPIServerService initializes the API server
func NewAPIServerService(ctx context.Context, port uint64, exitsStorage ports.ExitsStorage, performanceStorage ports.PerformanceStorage, relays ports.Relays, validatorExitRequestEventScanner *ValidatorExitRequestEventScanner, allowedOrigins []string) *APIServerService {
	apiServer := &APIServerService{
		server:                           &http.Server{Addr: ":" + strconv.FormatUint(port, 10)},
		servicePrefix:                    "API",
		ctx:                              ctx,
		exitsStorage:                     exitsStorage,
		performanceStorage:               performanceStorage,
		relays:                           relays,
		validatorExitRequestEventScanner: validatorExitRequestEventScanner,
		router:                           mux.NewRouter(),
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
	h.router.HandleFunc("/api/v0/events_indexer/operatorId", h.addOperator).Methods("POST", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/operatorId", h.deleteOperator).Methods("DELETE", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/operator_performance", h.getOperatorPerformance).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/exit_requests", h.getExitRequests).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/relays_allowed", h.getRelaysAllowed).Methods("GET", "OPTIONS")
	h.router.HandleFunc("/api/v0/events_indexer/relays_used", h.getRelaysUsed).Methods("GET", "OPTIONS")

	// Add a generic OPTIONS handler to ensure preflight requests are handled
	h.router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond to OPTIONS requests with 200 OK
	})
}

// Handlers

// getRelaysAllowed retrieves the list of allowed relays
func (h *APIServerService) getRelaysAllowed(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.servicePrefix, "getRelaysAllowed request received")
	relays, err := h.relays.GetRelaysAllowList(h.ctx)
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
	relays, err := h.relays.GetRelaysUsed(h.ctx)
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
	operatorIds, err := h.exitsStorage.GetOperatorIds()
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

	if err := h.exitsStorage.DeleteOperator(operatorId); err != nil {
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
	operatorIds, err := h.exitsStorage.GetOperatorIds()
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

	if err := h.exitsStorage.SaveOperatorId(req.OperatorID); err != nil {
		logger.ErrorWithPrefix(h.servicePrefix, "Failed to update Operator ID: %v", err)
		writeErrorResponse(w, "Failed to update Operator ID", http.StatusInternalServerError, err)
		return
	}

	// Trigger exit scanner asynchronously after successfully saving the operator ID
	go func() {
		if err := h.validatorExitRequestEventScanner.RunScan(h.ctx); err != nil {
			logger.ErrorWithPrefix(h.servicePrefix, "Async exit events scanner after addOperator failed: %v", err)
		}
	}()

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

	// Fetch operator performance data
	report, err := h.performanceStorage.GetNoPerformance(operatorId)
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

	exitRequests, err := h.exitsStorage.GetExitRequests(operatorIdNum)
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
