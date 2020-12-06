package current

import (
	"net/http"

	"dolittle.io/login/clients/kratos"
	"dolittle.io/login/identities/users"
)

type Getter interface {
	GetCurrentUser(r *http.Request) (*users.User, error)
}

func NewGetter() Getter {
	return &getter{}
}

type getter struct {
	configuration Configuration
	kratos        kratos.Client
	parser        Parser
}

func (g *getter) GetCurrentUser(r *http.Request) (*users.User, error) {
	cookie, err := r.Cookie(g.configuration.Cookie())
	if err != nil {
		return nil, err
	}

	session, err := g.kratos.GetCurrentUser(r.Context(), cookie.Value)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseUserFrom(session)
}
