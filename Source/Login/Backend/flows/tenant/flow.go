package tenant

import "dolittle.io/login/identities/tenants"

type FlowID string

type Flow struct {
	ID FlowID

	AvailableTenants []tenants.Tenant
}
