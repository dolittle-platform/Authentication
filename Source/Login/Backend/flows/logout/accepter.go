package logout

import (
	"context"
	"dolittle.io/login/clients/hydra"
	"net/url"
)

type Accepter interface {
	AcceptLogoutFlow(ctx context.Context, flow *Flow) (*url.URL, error)
}

func NewAccepter(hydra hydra.Client) Accepter {
	return &accepter{
		hydra: hydra,
	}
}

type accepter struct {
	hydra hydra.Client
}

func (a *accepter) AcceptLogoutFlow(ctx context.Context, flow *Flow) (*url.URL, error) {
	response, err := a.hydra.AcceptLogoutRequest(ctx, string(flow.ID))
	if err != nil {
		return nil, err
	}

	redirect, err := url.Parse(*response.RedirectTo)
	if err != nil {
		return nil, err
	}

	return redirect, nil
}
