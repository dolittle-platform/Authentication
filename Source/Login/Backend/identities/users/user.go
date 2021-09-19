package users

import "dolittle.io/login/identities/tenants"

type User struct {
	Subject string           `json:"subject"`
	Tenants []tenants.Tenant `json:"tenants"`
}

func (u *User) HasAccessToTenant(givenID tenants.TenantID) bool {
	for _, tenant := range u.Tenants {
		if givenID == tenant.ID {
			return true
		}
	}
	return false
}
