package tenant

import (
	"dolittle.io/login/flows/forms"
	"dolittle.io/login/identities/users"
	"github.com/ory/hydra-client-go/models"
)

type Parser interface {
	ParseTenantFlowFrom(response *models.LoginRequest, user *users.User) (*Flow, error)
}

func NewParser(configuration Configuration) Parser {
	return &parser{
		configuration: configuration,
	}
}

type parser struct {
	configuration Configuration
}

func (p *parser) ParseTenantFlowFrom(response *models.LoginRequest, user *users.User) (*Flow, error) {
	return &Flow{
		ID: FlowID(*response.Challenge),
		Form: forms.Form{
			SubmitMethod: "POST",
			SubmitAction: p.configuration.SelectTenantFormSubmitAction().String(),
		},
		User: user,
	}, nil
}
