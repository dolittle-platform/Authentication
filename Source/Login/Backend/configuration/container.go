package configuration

import (
	"dolittle.io/login/clients/hydra"
	"dolittle.io/login/clients/kratos"
	"dolittle.io/login/configuration/changes"
	consentFlow "dolittle.io/login/flows/consent"
	loginFlow "dolittle.io/login/flows/login"
	logoutFlow "dolittle.io/login/flows/logout"
	tenantFlow "dolittle.io/login/flows/tenant"
	"dolittle.io/login/identities/current"
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/providers"
	"dolittle.io/login/server"
	"dolittle.io/login/server/public"
	"dolittle.io/login/server/public/consent"
	"dolittle.io/login/server/public/login"
	"dolittle.io/login/server/public/logout"
	"dolittle.io/login/server/public/tenant"
	"go.uber.org/zap"
)

type Container struct {
	Notifier changes.ConfigurationChangeNotifier

	HydraClient  hydra.Client
	KratosClient kratos.Client

	TenantGetter tenants.Getter

	CurrentUserParser current.Parser
	CurrentUserGetter current.Getter

	ProvidersGetter providers.Getter

	LoginFlowGetter loginFlow.Getter
	LoginFlowParser loginFlow.Parser

	TenantFlowParser   tenantFlow.Parser
	TenantFlowGetter   tenantFlow.Getter
	TenantFlowSelecter tenantFlow.Selecter

	ConsentFlowParser   consentFlow.Parser
	ConsentFlowGetter   consentFlow.Getter
	ConsentFlowAccepter consentFlow.Accepter

	LogoutFlowParser    logoutFlow.Parser
	LogoutFlowGetter    logoutFlow.Getter
	LogoutFlowAccepter  logoutFlow.Accepter
	LogoutFlowInitiator logoutFlow.Initiator

	FrontendHandler public.FrontendHandler

	LoginGetHandler login.GetHandler

	TenantInitiateHandler tenant.InitiateHandler
	TenantGetHandler      tenant.GetHandler
	TenantSelectHandler   tenant.SelectHandler

	ConsentInitiateHandler consent.InitiateHandler

	LogoutInitiateHandler logout.InitiateHandler

	Server server.Server
}

func NewContainer(configuration Configuration) (*Container, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	container := &Container{}

	container.Notifier = changes.NewConfigurationChangeNotifier(
		logger)

	hydraClient, err := hydra.NewClient(
		configuration.Clients().Hydra(),
		container.Notifier)
	if err != nil {
		return nil, err
	}
	container.HydraClient = hydraClient

	kratosClient, err := kratos.NewClient(
		configuration.Clients().Kratos(),
		container.Notifier)
	if err != nil {
		return nil, err
	}
	container.KratosClient = kratosClient

	container.TenantGetter = tenants.NewGetter(
		configuration.Identities())

	container.CurrentUserParser = current.NewParser(
		container.TenantGetter)

	container.CurrentUserGetter = current.NewGetter(
		configuration.Identities(),
		container.KratosClient,
		container.CurrentUserParser)

	container.ProvidersGetter = providers.NewGetter(
		configuration.Providers(),
	)

	container.LoginFlowParser = loginFlow.NewParser(
		configuration.Flows().Login(),
		container.ProvidersGetter)

	container.LoginFlowGetter = loginFlow.NewGetter(
		configuration.Flows().Login(),
		container.KratosClient,
		container.LoginFlowParser)

	container.TenantFlowParser = tenantFlow.NewParser(
		configuration.Flows().Tenant())

	container.TenantFlowGetter = tenantFlow.NewGetter(
		configuration.Flows().Tenant(),
		container.HydraClient,
		container.CurrentUserGetter,
		container.TenantFlowParser)

	container.TenantFlowSelecter = tenantFlow.NewSelecter(
		configuration.Flows().Tenant(),
		container.HydraClient)

	container.ConsentFlowParser = consentFlow.NewParser()

	container.ConsentFlowGetter = consentFlow.NewGetter(
		configuration.Flows().Consent(),
		container.HydraClient,
		container.ConsentFlowParser)

	container.ConsentFlowAccepter = consentFlow.NewAccepter(
		container.HydraClient)

	container.LogoutFlowParser = logoutFlow.NewParser()

	container.LogoutFlowGetter = logoutFlow.NewGetter(
		configuration.Flows().Logout(),
		container.HydraClient,
		container.LogoutFlowParser)

	container.LogoutFlowAccepter = logoutFlow.NewAccepter(
		container.HydraClient)

	container.LogoutFlowInitiator = logoutFlow.NewInitiator(
		configuration.Flows().Logout(),
		container.KratosClient)

	frontendHandler, err := public.NewFrontendHandler(
		configuration.Server())
	if err != nil {
		return nil, err
	}
	container.FrontendHandler = frontendHandler

	container.LoginGetHandler = login.NewGetHandler(
		container.LoginFlowGetter)

	container.TenantInitiateHandler = tenant.NewInitiateHandler(
		container.TenantFlowGetter,
		container.TenantFlowSelecter)

	container.TenantGetHandler = tenant.NewGetHandler(
		container.TenantFlowGetter)

	container.TenantSelectHandler = tenant.NewSelectHandler(
		container.TenantFlowGetter,
		container.TenantFlowSelecter)

	container.ConsentInitiateHandler = consent.NewInitiateHandler(
		container.ConsentFlowGetter,
		container.ConsentFlowAccepter)

	container.LogoutInitiateHandler = logout.NewInitiateHandler(
		container.LogoutFlowGetter,
		container.LogoutFlowAccepter,
		container.LogoutFlowInitiator)

	container.Server = server.NewServer(
		configuration.Server(),
		container.FrontendHandler,
		container.LoginGetHandler,
		container.TenantInitiateHandler,
		container.TenantGetHandler,
		container.TenantSelectHandler,
		container.ConsentInitiateHandler,
		container.LogoutInitiateHandler,
		logger)

	return container, nil
}
