package tenants

type TenantID string

type Tenant struct {
	ID      TenantID
	Display string
}
