package tenants

type Configuration interface {
	TenantNames() map[TenantID]string
}
