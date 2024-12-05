package ports

import "net/http"

// ProxyAPI is the interface for the Proxy API adapter.
type ProxyAPI interface {
	GetRouter() http.Handler
}
