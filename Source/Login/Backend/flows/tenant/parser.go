package tenant

import (
	"dolittle.io/login/identities/users"
	"github.com/ory/hydra-client-go/models"
)

type Parser interface {
	ParseTenantFlowFrom(response *models.LoginRequest, user *users.User) (*Flow, error)
}
