package logout

import "github.com/ory/hydra-client-go/models"

type Parser interface {
	ParseLogoutFlowFrom(response *models.LogoutRequest) (*Flow, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

func (p *parser) ParseLogoutFlowFrom(response *models.LogoutRequest) (*Flow, error) {
	return &Flow{
		ID:      FlowID(response.Challenge),
		Subject: response.Subject,
	}, nil
}
