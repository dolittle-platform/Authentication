package current

import (
	"dolittle.io/login/identities/users"
	"github.com/ory/kratos-client-go/models"
)

type Parser interface {
	ParseUserFrom(session *models.Session) (*users.User, error)
}
