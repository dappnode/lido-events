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

type ClientsProxyAPIServerService struct {
	server        *http.Server
	servicePrefix string
}

func NewClientsProxyAPIServerService(clientsProxyApiAdapter ports.ClientsProxyApi, port uint64) *ClientsProxyAPIServerService {
	return &ClientsProxyAPIServerService{
		server: &http.Server{
			Addr:    ":" + strconv.FormatUint(port, 10),
			Handler: clientsProxyApiAdapter.GetRouter(),
		},
		servicePrefix: "Proxy API",
	}
}

func (s *ClientsProxyAPIServerService) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.InfoWithPrefix(s.servicePrefix, "server started on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			logger.FatalWithPrefix(s.servicePrefix, "server ListenAndServe: %v", err)
		}
	}()
}

func (s *ClientsProxyAPIServerService) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.WarnWithPrefix(s.servicePrefix, "server Shutdown: %v", err)
	}
}
