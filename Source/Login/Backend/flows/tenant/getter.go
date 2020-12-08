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
	var (
		id  string
		err error
	)

	switch r.Method {
	case http.MethodGet:
		id, err = g.getFlowIDFromGetRequest(r)
	case http.MethodPost:
		id, err = g.getFlowIDFromPostRequest(r)
	default:
		return nil, http.ErrNotSupported
	}
	if err != nil {
		return nil, err
	}

	loginFlowRequest, err := g.hydra.GetLoginFlow(r.Context(), id)
	if err != nil {
		return nil, err
	}

	user, err := g.users.GetCurrentUser(r)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseTenantFlowFrom(loginFlowRequest, user)
}

func (g *getter) getFlowIDFromGetRequest(r *http.Request) (string, error) {
	id := r.URL.Query().Get(g.configuration.FlowIDQueryParameter())
	if id == "" {
		return "", errors.New("no flow id set in GET request")
	}
	return id, nil
}

func (g *getter) getFlowIDFromPostRequest(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	id := r.Form.Get(g.configuration.FlowIDFormParameter())
	if id == "" {
		return "", errors.New("no flow id set in POST request")
	}
	return id, nil
}
