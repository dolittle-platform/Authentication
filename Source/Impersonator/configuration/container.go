package configuration

import (
	"dolittle.io/impersonator/audit"
	"dolittle.io/impersonator/configuration/changes"
	"dolittle.io/impersonator/identities"
	"dolittle.io/impersonator/proxy"
	"dolittle.io/impersonator/proxy/client"
	"dolittle.io/impersonator/proxy/context"
	"dolittle.io/impersonator/server"
	"go.uber.org/zap"
)

type Container struct {
	Notifier changes.ConfigurationChangeNotifier

	AuditLogger audit.Logger

	IdentitiesReader identities.Reader
	IdentitiesWriter identities.Writer

	ContextCreator context.Creator

	ProxyClient   client.Client
	ProxyModifier proxy.Modifier
	ProxyHandler  proxy.Handler

	Server server.Server
}

func NewContainer(config Configuration) (*Container, error) {
	logger, _ := zap.NewDevelopment()
	container := &Container{}

	container.Notifier = changes.NewConfigurationChangeNotifier(logger)
	config.OnChange(container.Notifier.TriggerChanged)

	container.AuditLogger = audit.NewLogger(
		logger)

	container.IdentitiesReader = identities.NewReader(
		config.Identities())
	container.IdentitiesWriter = identities.NewWriter(
		config.Identities())

	container.ContextCreator = context.NewCreator(
		container.IdentitiesReader)

	proxyClient, err := client.NewClient(
		config.Proxy())
	if err != nil {
		return nil, err
	}
	container.ProxyClient = proxyClient

	container.ProxyModifier = proxy.NewModifier(
		config.Proxy(),
		container.AuditLogger,
		container.IdentitiesWriter,
		logger)
	container.ProxyHandler = proxy.NewHandler(
		container.ContextCreator,
		container.ProxyModifier,
		container.ProxyClient,
		logger)

	container.Server = server.NewServer(
		config.Server(),
		container.Notifier,
		container.ProxyHandler,
		logger)

	return container, nil
}
