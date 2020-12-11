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

var defaultTenantNameMap = make(map[tenants.TenantID]string, 0)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) Cookie() string {
	if value := viper.GetString(identitiesCurrentUserCookieNameKey); value != "" {
		return value
	}
	return defaultIdentitiesCurrentUserCookieNameValue
}

func (t *identitiesConfiguration) TenantNamesMap() *map[tenants.TenantID]string {
	tenantNameMap := defaultTenantNameMap
	if value := viper.GetStringMapString(identitiesTenantNameMapKey); value != nil && len(value) > 0 {
		tenantNameMap = createTenantNameMap(value)
	}
	return &tenantNameMap
}

func createTenantNameMap(in map[string]string) map[tenants.TenantID]string {
	out := make(map[tenants.TenantID]string, len(in))
	for tenantIDString, tenantName := range in {
		out[tenants.TenantID(tenantIDString)] = tenantName
	}

	return out
}
