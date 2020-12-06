package users

import "dolittle.io/login/identities/tenants"

type User struct {
	Subject string
	Tenants []tenants.Tenant
}

func (u *User) HasAccessToTenant(givenTenant tenants.Tenant) bool {
	for _, tenant := range u.Tenants {
		if givenTenant == tenant {
			return true
		}
	}
	return false
}
