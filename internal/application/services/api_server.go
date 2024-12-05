// internal/application/services/api_server_service.go
package services

import (
	"context"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"net/http"
	"strconv"
	"sync"
)

type APIServerService struct {
	server        *http.Server
	servicePrefix string
}

func NewAPIServerService(apiAdapter ports.API, port uint64) *APIServerService {
	return &APIServerService{
		server: &http.Server{
			Addr:    ":" + strconv.FormatUint(port, 10),
			Handler: apiAdapter.GetRouter(),
		},
		servicePrefix: "API",
	}
}

func (s *APIServerService) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.InfoWithPrefix(s.servicePrefix, "server started on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			logger.FatalWithPrefix(s.servicePrefix, "server ListenAndServe: %v", err)
		}
	}()
}

func (s *APIServerService) Shutdown(ctx context.Context) {
	if err := s.server.Shutdown(ctx); err != nil {
		logger.WarnWithPrefix(s.servicePrefix, "server Shutdown: %v", err)
	}
}
