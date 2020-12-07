package current

import (
	"net/http"

	"dolittle.io/login/clients/kratos"
	"dolittle.io/login/identities/users"
	"github.com/ory/kratos-client-go/client/public"
)

type Getter interface {
	GetCurrentUser(r *http.Request) (*users.User, error)
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

func (g *getter) GetCurrentUser(r *http.Request) (*users.User, error) {
	cookie, err := r.Cookie(g.configuration.Cookie())
	if err == http.ErrNoCookie {
		return nil, ErrNoUserLoggedIn
	}
	if err != nil {
		return nil, err
	}
	session, err := g.kratos.GetCurrentUser(r.Context(), cookie.Value)
	if _, unauthorized := err.(*public.WhoamiUnauthorized); unauthorized {
		return nil, ErrNoUserLoggedIn
	}
	if err != nil {
		return nil, err
	}

	return g.parser.ParseUserFrom(session)
}
