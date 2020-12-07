package tenant

import (
	"errors"
	"net/http"

	"dolittle.io/login/clients/hydra"
	"dolittle.io/login/identities/current"
)

type Getter interface {
	GetTenantFlowFrom(r *http.Request) (*Flow, error)
}

func NewGetter(configuration Configuration, hydra hydra.Client, users current.Getter, parser Parser) Getter {
	return &getter{
		configuration: configuration,
		hydra:         hydra,
		users:         users,
		parser:        parser,
	}
}

type getter struct {
	configuration Configuration
	hydra         hydra.Client
	users         current.Getter
	parser        Parser
}

func (g *getter) GetTenantFlowFrom(r *http.Request) (*Flow, error) {
	id := r.URL.Query().Get(g.configuration.FlowIDQueryParameter())
	if id == "" {
		return nil, errors.New("no flow id set in request")
	}

	flow, err := g.hydra.GetLoginFlow(r.Context(), id)
	if err != nil {
		return nil, err
	}

	user, err := g.users.GetCurrentUser(r)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseTenantFlowFrom(flow, user)
}
