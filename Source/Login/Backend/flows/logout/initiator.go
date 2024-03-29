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

func NewInitiator(configuration Configuration, kratos kratos.Client) Initiator {
	return &initiator{
		configuration: configuration,
		kratos:        kratos,
	}
}

type initiator struct {
	configuration Configuration
	kratos        kratos.Client
}

func (i *initiator) InitiateLogout(ctx context.Context, cookies []*http.Cookie) (*url.URL, error) {
	response, err := i.kratos.GetLogoutURL(ctx, cookies)
	if err == kratos.ErrKratosUnauthorized {
		return i.configuration.LoggedOutRedirect(), nil
	}
	if err != nil {
		return nil, err
	}

	redirect, err := url.Parse(response.LogoutUrl)
	if err != nil {
		return nil, err
	}

	return redirect, nil
}
