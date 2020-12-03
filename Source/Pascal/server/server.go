package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/server/handling"
	"dolittle.io/pascal/server/public"
	"go.uber.org/zap"
)

type Server interface {
	Run() error
}

func NewServer(configuration Configuration, notifier changes.ConfigurationChangeNotifier, initiate public.InitiateHandler, complete public.CompleteHandler, logger *zap.Logger) Server {
	return &server{
		configuration:    configuration,
		notifier:         notifier,
		initiateHandler:  initiate,
		completeHandler:  complete,
		logger:           logger,
		shutdownComplete: make(chan struct{}),
	}
}

type server struct {
	configuration    Configuration
	notifier         changes.ConfigurationChangeNotifier
	initiateHandler  public.InitiateHandler
	completeHandler  public.CompleteHandler
	logger           *zap.Logger
	httpServer       *http.Server
	shutdownComplete chan struct{}
}

func (s *server) Run() error {
	if err := s.notifier.RegisterCallback("server", s.handleConfigurationChanged); err != nil {
		return err
	}
	return s.loop()
}

func (s *server) loop() error {
	for {
		if err := s.run(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("http server failed", zap.Error(err))

			select {
			case <-s.shutdownComplete:
			case <-time.After(2 * time.Second):
			}
		}
	}
}

func (s *server) run() error {
	s.httpServer = &http.Server{}

	s.logger.Info("Exposing initate endpoint on", zap.String("path", s.configuration.InitiatePath()))
	s.logger.Info("Exposing complete endpoint on", zap.String("path", s.configuration.CompletePath()))

	router := handling.NewRouter(s.configuration, s.logger)
	router.Handle(s.configuration.InitiatePath(), s.initiateHandler)
	router.Handle(s.configuration.CompletePath(), s.completeHandler)

	s.logger.Info("Starting server", zap.Int("port", s.configuration.Port()))
	s.httpServer.Addr = fmt.Sprintf(":%d", s.configuration.Port())
	s.httpServer.Handler = router

	return s.httpServer.ListenAndServe()
}

func (s *server) handleConfigurationChanged() error {
	err := s.httpServer.Shutdown(context.Background())
	select {
	case s.shutdownComplete <- struct{}{}:
	default:
	}
	return err
}
