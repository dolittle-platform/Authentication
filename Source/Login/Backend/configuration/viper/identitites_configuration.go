package viper

import (
	"dolittle.io/login/identities/tenants"
	"github.com/spf13/viper"
)

const (
	identitiesCurrentUserCookieNameKey = "identities.cookie_name"
	identitiesTenantNameMapKey         = "identities.tenants"

	defaultIdentitiesCurrentUserCookieNameValue = "ory_kratos_session"
)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) Cookie() string {
	if value := viper.GetString(identitiesCurrentUserCookieNameKey); value != "" {
		return value
	}
	return defaultIdentitiesCurrentUserCookieNameValue
}

func (c *identitiesConfiguration) TenantNamesMap() map[tenants.TenantID]string {
	return createTenantNameMap(viper.GetStringMapString(identitiesTenantNameMapKey))
}

func createTenantNameMap(in map[string]string) map[tenants.TenantID]string {
	if in == nil || len(in) == 0 {
		return nil
	}
	out := make(map[tenants.TenantID]string, len(in))
	for tenantIDString, tenantName := range in {
		out[tenants.TenantID(tenantIDString)] = tenantName
	}

	return out
}
