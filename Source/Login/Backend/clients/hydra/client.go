package hydra

import (
	"context"

	"dolittle.io/login/configuration/changes"
	ory "github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
)

type Client interface {
	GetLoginFlow(ctx context.Context, flowID string) (*models.LoginRequest, error)
	GetConsentFlow(ctx context.Context, flowID string) (*models.ConsentRequest, error)
}

func NewClient(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (Client, error) {
	client := &client{
		configuration: configuration,
		client:        getORYClient(configuration),
	}
	if err := notifier.RegisterCallback("hydra-client", client.handleConfigurationChanged); err != nil {
		return nil, err
	}
	return client, nil
}

type client struct {
	configuration Configuration
	client        *ory.OryHydra
}

func (c *client) GetLoginFlow(ctx context.Context, flowID string) (*models.LoginRequest, error) {
	params := admin.NewGetLoginRequestParams().WithLoginChallenge(flowID).WithContext(ctx)
	response, err := c.client.Admin.GetLoginRequest(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) GetConsentFlow(ctx context.Context, flowID string) (*models.ConsentRequest, error) {
	params := admin.NewGetConsentRequestParams().WithConsentChallenge(flowID).WithContext(ctx)
	response, err := c.client.Admin.GetConsentRequest(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) handleConfigurationChanged() error {
	c.client = getORYClient(c.configuration)
	return nil
}

func getORYClient(configuration Configuration) *ory.OryHydra {
	url := configuration.Endpoint()
	config := ory.DefaultTransportConfig().WithSchemes([]string{url.Scheme}).WithHost(url.Host).WithBasePath(url.Path)
	return ory.NewHTTPClientWithConfig(nil, config)
}
