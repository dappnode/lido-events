package services

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"lido-events/internal/logger"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ProxyAPIServerService struct {
	server        *http.Server
	servicePrefix string
	proxyApiURL   string
	router        *mux.Router
}

// NewProxyAPIServerService initializes the Proxy API server
func NewProxyAPIServerService(port uint64, proxyApiURL string, allowedOrigins []string) *ProxyAPIServerService {
	router := mux.NewRouter()

	service := &ProxyAPIServerService{
		server: &http.Server{
			Addr: ":" + strconv.FormatUint(port, 10),
		},
		servicePrefix: "Proxy API",
		proxyApiURL:   proxyApiURL,
		router:        router,
	}

	service.setupRoutes()

	// Configure CORS middleware
	corsAllowedOrigins := handlers.AllowedOrigins(allowedOrigins)
	corsAllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	corsAllowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	router.Use(handlers.CORS(
		corsAllowedOrigins,
		corsAllowedMethods,
		corsAllowedHeaders,
	))

	service.server.Handler = router

	return service
}

// Start starts the Proxy API server
func (s *ProxyAPIServerService) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.InfoWithPrefix(s.servicePrefix, "server started on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			logger.FatalWithPrefix(s.servicePrefix, "server ListenAndServe: %v", err)
		}
	}()
}

// Shutdown gracefully shuts down the Proxy API server
func (s *ProxyAPIServerService) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.WarnWithPrefix(s.servicePrefix, "server Shutdown: %v", err)
	}
}

// setupRoutes sets up all the routes for the Proxy API
func (s *ProxyAPIServerService) setupRoutes() {
	s.router.PathPrefix("/v1/").HandlerFunc(s.proxyHandler).Methods("GET", "POST", "OPTIONS")

	// Add a generic OPTIONS handler to ensure preflight requests are handled
	s.router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // Respond to OPTIONS requests with 200 OK
	})
}

// proxyHandler handles all requests and proxies them to the target API
func (s *ProxyAPIServerService) proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Construct the target URL
	targetURL, err := url.Parse(s.proxyApiURL)
	if err != nil {
		http.Error(w, "Invalid proxy URL", http.StatusInternalServerError)
		return
	}

	// Append the incoming request's path and query parameters to the target URL
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
			continue // Skip the Origin header
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

	// Add CORS headers to the response
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) // Reflect the origin for allowed requests
	w.Header().Set("Access-Control-Allow-Credentials", "true")            // Allow credentials if needed
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Copy other response headers, excluding CORS-related headers
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
