package consent

import (
	"github.com/ory/hydra-client-go/models"

	"dolittle.io/login/flows/context"
)

type Parser interface {
	ParseConsentFlowFrom(response *models.ConsentRequest) (*Flow, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

func (p *parser) ParseConsentFlowFrom(response *models.ConsentRequest) (*Flow, error) {
	flowContext, err := context.RetrieveFrom(response)
	if err != nil {
		return nil, err
	}

	return &Flow{
		ID:             FlowID(*response.Challenge),
		User:           flowContext.User,
		SelectedTenant: flowContext.SelectedTenant,
	}, nil
}
