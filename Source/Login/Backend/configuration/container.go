package configuration

import (
	"dolittle.io/login/server"
	"dolittle.io/login/server/public"
	"dolittle.io/login/server/public/consent"
	"dolittle.io/login/server/public/login"
	"dolittle.io/login/server/public/tenant"
	"go.uber.org/zap"
)

type Container struct {
	FrontendHandler public.FrontendHandler

	LoginGetHandler login.GetHandler

	TenantInitiateHandler tenant.InitiateHandler
	TenantGetHandler      tenant.GetHandler
	TenantSelectHandler   tenant.SelectHandler

	ConsentInitiateHandler consent.InitiateHandler

	Server server.Server
}

func NewContainer(configuration Configuration) (*Container, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	container := &Container{}

	container.FrontendHandler = public.NewFrontendHandler(
		configuration.Server())

	container.Server = server.NewServer(
		configuration.Server(),
		container.FrontendHandler,
		container.LoginGetHandler,
		container.TenantInitiateHandler,
		container.TenantGetHandler,
		container.TenantSelectHandler,
		container.ConsentInitiateHandler,
		logger)

	return container, nil
}
