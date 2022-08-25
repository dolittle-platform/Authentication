package server

import (
	"dolittle.io/login/server/httputils"
	"dolittle.io/login/server/public/logout"
	"fmt"
	"net/http"

	"dolittle.io/login/server/handling"
	"dolittle.io/login/server/public"
	"dolittle.io/login/server/public/consent"
	"dolittle.io/login/server/public/login"
	"dolittle.io/login/server/public/tenant"
	"go.uber.org/zap"
)

type Server interface {
	Run() error
}

func NewServer(
	configuration Configuration,
	frontend public.FrontendHandler,
	loginGet login.GetHandler,
	tenantInitiate tenant.InitiateHandler,
	tenantGet tenant.GetHandler,
	tenantSelect tenant.SelectHandler,
	consentInitiate consent.InitiateHandler,
	logoutInitiate logout.InitiateHandler,
	logger *zap.Logger) Server {
	return &server{
		configuration:   configuration,
		frontend:        frontend,
		loginGet:        loginGet,
		tenantInitiate:  tenantInitiate,
		tenantGet:       tenantGet,
		tenantSelect:    tenantSelect,
		consentInitiate: consentInitiate,
		logoutInitiate:  logoutInitiate,
		logger:          logger,
	}
}

type server struct {
	configuration Configuration

	frontend public.FrontendHandler

	loginGet login.GetHandler

	tenantInitiate tenant.InitiateHandler
	tenantGet      tenant.GetHandler
	tenantSelect   tenant.SelectHandler

	consentInitiate consent.InitiateHandler

	logoutInitiate logout.InitiateHandler

	logger *zap.Logger
}

func (s *server) Run() error {
	router := handling.NewRouter(s.configuration, s.logger)

	router.Handle("/.auth/", s.frontend)

	router.Handle("/.auth/self-service/login/flows", s.loginGet)

	router.Handle("/.auth/self-service/tenant/browser", s.tenantInitiate)
	router.Handle("/.auth/self-service/tenant/flows", s.tenantGet)
	router.Handle("/.auth/self-service/tenant/select", s.tenantSelect)

	router.Handle("/.auth/self-service/consent/browser", s.consentInitiate)

	router.Handle("/.auth/self-service/logout/browser", s.logoutInitiate)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.Port()),
		Handler: httputils.WithCacheControl("no-store", router),
	}

	s.logger.Info("Starting server", zap.Int("port", s.configuration.Port()))

	return server.ListenAndServe()
}
