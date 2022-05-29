package logout

import (
	"context"
	"dolittle.io/login/clients/kratos"
	"net/http"
	"net/url"
)

type Initiator interface {
	InitiateLogout(ctx context.Context, cookies []*http.Cookie) (*url.URL, error)
}

func NewInitiator(kratos kratos.Client) Initiator {
	return &initiator{
		kratos: kratos,
	}
}

type initiator struct {
	kratos kratos.Client
}

func (i *initiator) InitiateLogout(ctx context.Context, cookies []*http.Cookie) (*url.URL, error) {
	response, err := i.kratos.GetLogoutURL(ctx, cookies)
	if err != nil {
		return nil, err
	}

	redirect, err := url.Parse(response.LogoutUrl)
	if err != nil {
		return nil, err
	}

	return redirect, nil
}
