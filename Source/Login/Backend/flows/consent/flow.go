package consent

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
)

type FlowID string

type Flow struct {
	ID             FlowID           `json:"id"`
	User           *users.User      `json:"user"`
	SelectedTenant tenants.TenantID `json:"selectedTenant"`
}
