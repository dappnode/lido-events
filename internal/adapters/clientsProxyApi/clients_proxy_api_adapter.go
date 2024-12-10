package clientsproxyapi

import (
	"bytes"
	"io"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"net/http"
	"net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ClientsProxyAPIHandler provides an HTTP handler for the proxy API.
type ClientsProxyAPIHandler struct {
	Router        *mux.Router
	adapterPrefix string
}

// Ensure ClientsProxyAPIHandler implements the ports.ClientsProxyApi interface.
var _ ports.ClientsProxyApi = (*ClientsProxyAPIHandler)(nil)

// GetRouter implements the ports.ClientsProxyApi interface.
func (h *ClientsProxyAPIHandler) GetRouter() http.Handler {
	return h.Router
}

// NewClientsProxyAPIAdapter initializes the ClientsProxyAPIHandler and sets up routes.
func NewClientsProxyAPIAdapter(allowedOrigins []string) *ClientsProxyAPIHandler {
	h := &ClientsProxyAPIHandler{
		Router:        mux.NewRouter(),
		adapterPrefix: "CLIENTS-PROXY-API",
	}

	h.SetupRoutes()

	// Configure CORS middleware
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	h.Router.Use(handlers.CORS(
		corsAllowedOrigins,
		corsAllowedMethods,
		corsAllowedHeaders,
	))

	return h
}

// SetupRoutes sets up all the routes for the Clients Proxy API.
func (h *ClientsProxyAPIHandler) SetupRoutes() {
	h.Router.PathPrefix("/clients/proxy").HandlerFunc(h.proxyHandler).Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
}

func (h *ClientsProxyAPIHandler) proxyHandler(w http.ResponseWriter, r *http.Request) {
	logger.DebugWithPrefix(h.adapterPrefix, "proxyHandler", "Received request", r.Method, r.URL)

	// Get the target URL from the query parameters.
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		logger.ErrorWithPrefix(h.adapterPrefix, "proxyHandler", "Missing 'url' query parameter")
		http.Error(w, `{"error": "Missing 'url' query parameter"}`, http.StatusBadRequest)
		return
	}

	// Parse the target URL.
	parsedURL, err := url.Parse(targetURL)
	if err != nil || !parsedURL.IsAbs() {
		logger.ErrorWithPrefix(h.adapterPrefix, "proxyHandler", "Invalid 'url' query parameter", targetURL)
		http.Error(w, `{"error": "Invalid 'url' query parameter"}`, http.StatusBadRequest)
		return
	}

	// Create the proxy request.
	proxyReq, err := http.NewRequest(r.Method, parsedURL.String(), r.Body)
	if err != nil {
		logger.ErrorWithPrefix(h.adapterPrefix, "proxyHandler", "Failed to create proxy request", err)
		http.Error(w, `{"error": "Failed to create proxy request"}`, http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request, excluding the Origin header.
	for key, values := range r.Header {
		if key == "Origin" {
			continue // Skip the Origin header.
		}
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	// Execute the proxy request.
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		logger.ErrorWithPrefix(h.adapterPrefix, "proxyHandler", "Failed to execute proxy request", err)
		http.Error(w, `{"error": "Failed to execute proxy request"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Add CORS headers to the response.
	w.Header().Set("Access-Control-Allow-Origin", "*")         // Allow requests from any origin.
	w.Header().Set("Access-Control-Allow-Credentials", "true") // Allow credentials if needed.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Copy other response headers, excluding CORS-related headers.
	for key, values := range resp.Header {
		if key == "Access-Control-Allow-Origin" || key == "Origin" {
			continue
		}
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Copy the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorWithPrefix(h.adapterPrefix, "proxyHandler", "Failed to read response body", err)
		http.Error(w, `{"error": "Failed to read response body"}`, http.StatusInternalServerError)
		return
	}

	_, _ = io.Copy(w, bytes.NewReader(body))
}
