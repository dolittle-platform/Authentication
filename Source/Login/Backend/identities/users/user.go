package users

import "dolittle.io/login/identities/tenants"

type User struct {
	Subject string
	Tenants []tenants.Tenant
}

func (u *User) HasAccessToTenant(givenID tenants.TenantID) bool {
	for _, tenant := range u.Tenants {
		if givenID == tenant.ID {
			return true
		}
	}
	return false
}
