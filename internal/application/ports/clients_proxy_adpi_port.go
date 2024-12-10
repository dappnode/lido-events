package ports

import "net/http"

// ClientsProxyApi is the interface for the ClientsProxy API adapter.
type ClientsProxyApi interface {
	GetRouter() http.Handler
}
