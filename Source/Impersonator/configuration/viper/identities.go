package viper

import "github.com/spf13/viper"

const (
	identitiesUserIDHeaderKey   = "identities.headers.user_id"
	identitiesTenantIDHeaderKey = "identities.headers.tenant_id"

	defaultIdentitiesUserIDHeader   = "User-ID"
	defaultIdentitiesTenantIDHeader = "Tenant-ID"
)

type identitiesConfiguration struct{}

func (c *identitiesConfiguration) UserIDHeader() string {
	if header := viper.GetString(identitiesUserIDHeaderKey); header != "" {
		return header
	}
	return defaultIdentitiesUserIDHeader
}

func (c *identitiesConfiguration) TenantIDHeader() string {
	if header := viper.GetString(identitiesTenantIDHeaderKey); header != "" {
		return header
	}
	return defaultIdentitiesTenantIDHeader
}
