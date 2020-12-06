package consent

import "github.com/ory/hydra-client-go/models"

type Parser interface {
	ParseConsentFlowFrom(response *models.ConsentRequest) (*Flow, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

func (p *parser) ParseConsentFlowFrom(response *models.ConsentRequest) (*Flow, error) {
	return &Flow{}, nil
}
