package consent

import (
	"context"
	"net/url"

	"dolittle.io/login/clients/hydra"
	"github.com/ory/hydra-client-go/models"
)

type Accepter interface {
	AcceptConsentFlow(ctx context.Context, flow *Flow) (*url.URL, error)
}

func NewAccepter(hydra hydra.Client) Accepter {
	return &accepter{
		hydra: hydra,
	}
}

type accepter struct {
	hydra hydra.Client
}

func (a *accepter) AcceptConsentFlow(ctx context.Context, flow *Flow) (*url.URL, error) {
	tokenData := a.createTokenDataFrom(flow)

	response, err := a.hydra.AcceptConsentRequest(ctx, string(flow.ID), &models.AcceptConsentRequest{
		Remember: false,
		Session: &models.ConsentRequestSession{
			AccessToken: tokenData,
			IDToken:     tokenData,
		},
	})
	if err != nil {
		return nil, err
	}

	redirect, err := url.Parse(*response.RedirectTo)
	if err != nil {
		return nil, err
	}

	return redirect, nil
}

type tokenData struct {
	Subject string
	Tenant  string
}

func (a *accepter) createTokenDataFrom(flow *Flow) tokenData {
	return tokenData{
		Subject: flow.User.Subject,
		Tenant:  string(flow.SelectedTenant),
	}
}
