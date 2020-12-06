package login

import (
	"errors"
	"net/http"

	"dolittle.io/login/clients/kratos"
)

type Getter interface {
	GetLoginFlowFrom(r *http.Request) (*Flow, error)
}

func NewGetter() Getter {
	return &getter{}
}

type getter struct {
	configuration Configuration
	kratos        kratos.Client
	parser        Parser
}

func (g *getter) GetLoginFlowFrom(r *http.Request) (*Flow, error) {
	id := r.URL.Query().Get(g.configuration.FlowIdQueryParameter())
	if id == "" {
		return nil, errors.New("no flow id set in request")
	}

	flow, err := g.kratos.GetLoginFlow(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseLoginFlowFrom(flow)
}
