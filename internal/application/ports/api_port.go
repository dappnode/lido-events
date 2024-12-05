package ports

import "net/http"

// API is the interface for the API adapter.
type API interface {
	GetRouter() http.Handler
}
