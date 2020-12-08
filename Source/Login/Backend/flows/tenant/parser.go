package tenant

import (
	"dolittle.io/login/identities/users"
	"github.com/ory/hydra-client-go/models"
)

type Parser interface {
	ParseTenantFlowFrom(response *models.LoginRequest, user *users.User) (*Flow, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

func (p *parser) ParseTenantFlowFrom(response *models.LoginRequest, user *users.User) (*Flow, error) {
	return &Flow{
		ID:   FlowID(*response.Challenge),
		User: user,
	}, nil
}
