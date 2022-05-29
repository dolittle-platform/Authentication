package logout

import (
	"dolittle.io/login/clients/hydra"
	"net/http"
)

type Getter interface {
	GetLogoutFlowFrom(r *http.Request) (*Flow, error)
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

func (g *getter) GetLogoutFlowFrom(r *http.Request) (*Flow, error) {
	id := r.URL.Query().Get(g.configuration.FlowIDQueryParameter())
	if id == "" {
		return nil, ErrLogoutChallengeNotFound
	}

	flow, err := g.hydra.GetLogoutFlow(r.Context(), id)
	if err != nil {
		return nil, err
	}

	return g.parser.ParseLogoutFlowFrom(flow)
}
