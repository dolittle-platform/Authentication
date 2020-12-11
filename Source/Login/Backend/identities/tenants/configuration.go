package tenants

type Configuration interface {
	TenantNamesMap() map[TenantID]string
}
