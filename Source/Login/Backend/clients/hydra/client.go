package hydra

import (
	"context"

	"dolittle.io/login/configuration/changes"
	runtime "github.com/go-openapi/runtime/client"
	ory "github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
)

type Client interface {
	GetLoginFlow(ctx context.Context, flowID string) (*models.LoginRequest, error)
	GetConsentFlow(ctx context.Context, flowID string) (*models.ConsentRequest, error)
	GetLogoutFlow(ctx context.Context, flowID string) (*models.LogoutRequest, error)
	AcceptLoginRequest(ctx context.Context, flowID string, body *models.AcceptLoginRequest) (*models.CompletedRequest, error)
	AcceptConsentRequest(ctx context.Context, flowID string, body *models.AcceptConsentRequest) (*models.CompletedRequest, error)
	AcceptLogoutRequest(ctx context.Context, flowID string) (*models.CompletedRequest, error)
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

func (c *client) GetLogoutFlow(ctx context.Context, flowID string) (*models.LogoutRequest, error) {
	params := admin.NewGetLogoutRequestParams().WithLogoutChallenge(flowID).WithContext(ctx)
	response, err := c.client.Admin.GetLogoutRequest(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) AcceptLoginRequest(ctx context.Context, flowID string, body *models.AcceptLoginRequest) (*models.CompletedRequest, error) {
	params := admin.NewAcceptLoginRequestParams().WithLoginChallenge(flowID).WithBody(body).WithContext(ctx)
	response, err := c.client.Admin.AcceptLoginRequest(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) AcceptConsentRequest(ctx context.Context, flowID string, body *models.AcceptConsentRequest) (*models.CompletedRequest, error) {
	params := admin.NewAcceptConsentRequestParams().WithConsentChallenge(flowID).WithBody(body).WithContext(ctx)
	response, err := c.client.Admin.AcceptConsentRequest(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (c *client) AcceptLogoutRequest(ctx context.Context, flowID string) (*models.CompletedRequest, error) {
	params := admin.NewAcceptLogoutRequestParams().WithLogoutChallenge(flowID).WithContext(ctx)
	response, err := c.client.Admin.AcceptLogoutRequest(params)
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
	url := configuration.AdminEndpoint()

	http := newClientWithDefaultHeaders()
	http.Header.Add("X-Forwarded-Proto", "https")

	client := runtime.NewWithClient(url.Host, url.Path, []string{url.Scheme}, http.Client)

	return ory.New(client, nil)
}
