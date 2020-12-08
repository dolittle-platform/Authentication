package consent

import (
	"errors"
	"net/http"

	"dolittle.io/login/clients/hydra"
)

type Getter interface {
	GetConsentFlowFrom(r *http.Request) (*Flow, error)
}

func NewGetter(configuration Configuration, hydra hydra.Client, parser Parser) Getter {
	return &getter{
		configuration: configuration,
		hydra:         hydra,
		parser:        parser,
	}
}

type getter struct {
	configuration Configuration
	hydra         hydra.Client
	parser        Parser
}

func (g *getter) GetConsentFlowFrom(r *http.Request) (*Flow, error) {
	id := r.URL.Query().Get(g.configuration.FlowIDQueryParameter())
	if id == "" {
		return nil, errors.New("no flow id set in request")
	}

	flow, err := g.hydra.GetConsentFlow(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseConsentFlowFrom(flow)

}
