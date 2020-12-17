package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"dolittle.io/impersonator/configuration/changes"
	"dolittle.io/impersonator/proxy"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type Server interface {
	Run() error
}

func NewServer(configuration Configuration, notifier changes.ConfigurationChangeNotifier, handler proxy.Handler, logger *zap.Logger) Server {
	return &server{
		configuration: configuration,
		notifier:      notifier,
		handler:       handler,
		logger:        logger,
	}
}

type server struct {
	configuration Configuration
	notifier      changes.ConfigurationChangeNotifier
	handler       proxy.Handler
	logger        *zap.Logger
}

func (s *server) Run() error {
	wg := &sync.WaitGroup{}

	if err := s.startServerLoop("server-proxy", wg, s.createProxyServer); err != nil {
		return err
	}
	if err := s.startServerLoop("server-metrics", wg, s.createMetricServer); err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func (s *server) createProxyServer() *http.Server {
	router := http.NewServeMux()

	pathPrefix := s.configuration.ProxyPathPrefix()
	stripPrefix := pathPrefix
	if strings.HasSuffix(stripPrefix, "/") {
		stripPrefix = stripPrefix[:len(stripPrefix)-1]
	}
	handler := s.handler
	if len(stripPrefix) > 0 {
		handler = http.StripPrefix(stripPrefix, handler)
	}
	router.Handle(pathPrefix, handler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.ProxyPort()),
		Handler: router,
	}
}

func (s *server) createMetricServer() *http.Server {
	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.MetricsPort()),
		Handler: router,
	}
}

func (s *server) startServerLoop(component changes.ComponentName, wg *sync.WaitGroup, createServer func() *http.Server) error {
	server := &http.Server{}
	shutdownComplete := make(chan struct{})

	if err := s.notifier.RegisterCallback(component, func() error {
		server.Shutdown(context.Background())
		shutdownComplete <- struct{}{}
		return nil
	}); err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		for {
			server = createServer()
			s.logger.Info("Starting server", zap.String("address", server.Addr), zap.String("component", string(component)))
			if err := server.ListenAndServe(); err != nil {
				if err == http.ErrServerClosed {
					<-shutdownComplete
				} else {
					s.logger.Error("http server failed", zap.String("component", string(component)), zap.Error(err))
					<-time.After(2 * time.Second)
				}
			}
		}
		wg.Done()
	}()
	return nil
}
