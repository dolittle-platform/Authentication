package kratos

import (
	"context"
	"net/http"

	"dolittle.io/login/configuration/changes"
	ory "github.com/ory/kratos-client-go"
)

type Client interface {
	GetCurrentUser(ctx context.Context, cookie *http.Cookie) (*ory.Session, error)
	GetLoginFlow(ctx context.Context, flowID string, cookie *http.Cookie) (*ory.SelfServiceLoginFlow, error)
}

func NewClient(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (Client, error) {
	apiClient := getORYClient(configuration)
	client := &client{
		configuration: configuration,
		client:        apiClient,
		api:           apiClient.V0alpha1Api,
	}
	if err := notifier.RegisterCallback("kratos-client", client.handleConfigurationChanged); err != nil {
		return nil, err
	}
	return client, nil
}

type client struct {
	configuration Configuration
	client        *ory.APIClient
	api           ory.V0alpha1Api
}

func (c *client) GetCurrentUser(ctx context.Context, cookie *http.Cookie) (*ory.Session, error) {
	cookieHeaderValue := cookie.String()
	session, response, err := c.api.ToSession(ctx).Cookie(cookieHeaderValue).Execute()
	if response.StatusCode == http.StatusUnauthorized {
		return nil, ErrKratosUnauthorized
	}
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (c *client) GetLoginFlow(ctx context.Context, flowID string, cookie *http.Cookie) (*ory.SelfServiceLoginFlow, error) {
	cookieHeaderValue := cookie.String()
	flow, _, err := c.api.GetSelfServiceLoginFlow(ctx).Id(flowID).Cookie(cookieHeaderValue).Execute()
	if err != nil {
		return nil, err
	}
	return flow, nil
}

func (c *client) handleConfigurationChanged() error {
	apiClient := getORYClient(c.configuration)
	c.client = apiClient
	c.api = apiClient.V0alpha1Api
	return nil
}

func getORYClient(configuration Configuration) *ory.APIClient {
	url := configuration.PublicEndpoint()

	config := ory.NewConfiguration()
	config.Scheme = url.Scheme
	config.Host = url.Host

	return ory.NewAPIClient(config)
}
