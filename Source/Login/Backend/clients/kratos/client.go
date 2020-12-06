package kratos

import (
	"context"

	"dolittle.io/login/configuration/changes"
	openapi "github.com/go-openapi/runtime/client"
	ory "github.com/ory/kratos-client-go/client"
	"github.com/ory/kratos-client-go/client/public"
	"github.com/ory/kratos-client-go/models"
)

type Client interface {
	GetCurrentUser(ctx context.Context, cookie string) (*models.Session, error)
	GetLoginFlow(ctx context.Context, flowID string) (*models.LoginFlow, error)
}

func NewClient(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (Client, error) {
	client := &client{
		configuration: configuration,
		client:        getORYClient(configuration),
	}
	if err := notifier.RegisterCallback("kratos-client", client.handleConfigurationChanged); err != nil {
		return nil, err
	}
	return client, nil
}

type client struct {
	configuration Configuration
	client        *ory.OryKratos
}

func (c *client) GetCurrentUser(ctx context.Context, cookie string) (*models.Session, error) {
	params := public.NewWhoamiParams().WithCookie(&cookie).WithContext(ctx)
	response, err := c.client.Public.Whoami(params, openapi.PassThroughAuth)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) GetLoginFlow(ctx context.Context, flowID string) (*models.LoginFlow, error) {
	params := public.NewGetSelfServiceLoginFlowParams().WithID(flowID).WithContext(ctx)
	response, err := c.client.Public.GetSelfServiceLoginFlow(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) handleConfigurationChanged() error {
	c.client = getORYClient(c.configuration)
	return nil
}

func getORYClient(configuration Configuration) *ory.OryKratos {
	url := configuration.Endpoint()
	config := ory.DefaultTransportConfig().WithSchemes([]string{url.Scheme}).WithHost(url.Host).WithBasePath(url.Path)
	return ory.NewHTTPClientWithConfig(nil, config)
}
