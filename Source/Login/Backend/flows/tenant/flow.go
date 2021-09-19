package tenant

import (
	"dolittle.io/login/flows/forms"
	"dolittle.io/login/identities/users"
)

type FlowID string

type Flow struct {
	ID   FlowID      `json:"id"`
	Form forms.Form  `json:"form"`
	User *users.User `json:"user"`
}
