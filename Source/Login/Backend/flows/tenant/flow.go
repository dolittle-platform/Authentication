package tenant

import (
	"dolittle.io/login/identities/users"
)

type FlowID string

type Flow struct {
	ID   FlowID
	User *users.User
}
