package tenant

import (
	"net/url"

	"dolittle.io/login/identities/users"
)

type FlowID string

type Flow struct {
	ID               FlowID
	FormSubmitAction *url.URL
	FormSubmitMethod string

	User *users.User
}
