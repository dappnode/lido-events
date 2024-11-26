package proxyapi

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// APIHandler holds the necessary dependencies for API endpoints
type APIHandler struct {
	Router        *mux.Router
	proxyApiURL   string
	adapterPrefix string
}

// NewProxyAPIAdapter initializes the APIHandler and sets up routes
func NewProxyAPIAdapter(allowedOrigins []string, proxyApiURL string) *APIHandler {
	h := &APIHandler{
		mux.NewRouter(),
		proxyApiURL,
		"PROXY-API",
	}

	h.SetupRoutes()

	// Configure CORS middleware with restricted headers
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	h.Router.Use(func(next http.Handler) http.Handler {
		return handlers.CORS(corsAllowedMethods, corsAllowedHeaders)(next)
	})

	return h
}

// SetupRoutes sets up all the routes for the Proxy API
func (h *APIHandler) SetupRoutes() {
	h.Router.PathPrefix("/v1/").HandlerFunc(h.proxyHandler)
}

// proxyHandler handles all requests and proxies them to the target API
func (h *APIHandler) proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Construct the target URL
	targetURL, err := url.Parse(h.proxyApiURL)
	if err != nil {
		http.Error(w, "Invalid proxy URL", http.StatusInternalServerError)
		return
	}

	// Append the incoming request's path to the target URL
	targetURL.Path = r.URL.Path
	targetURL.RawQuery = r.URL.RawQuery

	// Create the proxy request
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create proxy request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request, excluding the Origin header
	for key, values := range r.Header {
		if key == "Origin" {
			continue
		}
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	// Execute the proxy request
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to execute proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy response headers, excluding CORS-related headers
	for key, values := range resp.Header {
		if key == "Access-Control-Allow-Origin" || key == "Origin" {
			continue
		}
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	_, _ = io.Copy(w, bytes.NewReader(body))
}
