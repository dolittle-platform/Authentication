package consent

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
)

type FlowID string

type Flow struct {
	ID             FlowID
	User           *users.User
	SelectedTenant tenants.TenantID
}
