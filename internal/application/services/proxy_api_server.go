// internal/application/services/proxy_api_server_service.go
package services

import (
	"context"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"net/http"
	"strconv"
	"sync"
)

type ProxyAPIServerService struct {
	server        *http.Server
	servicePrefix string
}

func NewProxyAPIServerService(proxyApiAdapter ports.ProxyAPI, port uint64) *ProxyAPIServerService {
	return &ProxyAPIServerService{
		server: &http.Server{
			Addr:    ":" + strconv.FormatUint(port, 10),
			Handler: proxyApiAdapter.GetRouter(),
		},
		servicePrefix: "Proxy API",
	}
}

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

func (s *ProxyAPIServerService) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.WarnWithPrefix(s.servicePrefix, "server Shutdown: %v", err)
	}
}
