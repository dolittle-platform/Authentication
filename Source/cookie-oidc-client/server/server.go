package server

import (
	"fmt"
	"net/http"

	"dolittle.io/cookie-oidc-client/server/handling"
	"dolittle.io/cookie-oidc-client/server/public"
	"go.uber.org/zap"
)

type Server interface {
	Run() error
}

func NewServer(configuration Configuration, initiate public.InitiateHandler, complete public.CompleteHandler, logger *zap.Logger) Server {
	handlers := handling.NewRouter(configuration, logger)

	handlers.Handle("/initiate", initiate)
	handlers.Handle("/complete", complete)

	return &server{
		configuration: configuration,
		logger:        logger,
		handlers:      handlers,
	}
}

type server struct {
	configuration Configuration
	logger        *zap.Logger
	handlers      handling.Router
}

func (s *server) Run() error {
	s.logger.Info("Server starting", zap.Int("port", s.configuration.Port()))
	return http.ListenAndServe(fmt.Sprintf(":%d", s.configuration.Port()), s.handlers)
}
