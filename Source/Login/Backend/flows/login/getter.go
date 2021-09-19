package login

import (
	"errors"
	"net/http"
	"strings"

	"dolittle.io/login/clients/kratos"
)

type Getter interface {
	GetLoginFlowFrom(r *http.Request) (*Flow, error)
}

func NewGetter(configuration Configuration, kratos kratos.Client, parser Parser) Getter {
	return &getter{
		configuration: configuration,
		kratos:        kratos,
		parser:        parser,
	}
}

type getter struct {
	configuration Configuration
	kratos        kratos.Client
	parser        Parser
}

func (g *getter) GetLoginFlowFrom(r *http.Request) (*Flow, error) {
	id := r.URL.Query().Get(g.configuration.FlowIDQueryParameter())
	if id == "" {
		return nil, errors.New("no flow id set in request")
	}

	csrfCookie, err := g.getCSRFCookie(r)
	if err != nil {
		return nil, err
	}

	flow, err := g.kratos.GetLoginFlow(r.Context(), id, csrfCookie)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseLoginFlowFrom(flow)
}

func (g *getter) getCSRFCookie(r *http.Request) (*http.Cookie, error) {
	cookieNamePrefix := g.configuration.CookiePrefix()
	for _, cookie := range r.Cookies() {
		cookieName := cookie.Name
		if strings.HasPrefix(cookieName, cookieNamePrefix) {
			return cookie, nil
		}
	}
	return nil, errors.New("CSRF token cookie was not found")
}
