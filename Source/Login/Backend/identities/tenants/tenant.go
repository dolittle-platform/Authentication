package tenants

type TenantID string

type Tenant struct {
	ID      TenantID `json:"id"`
	Display string   `json:"display"`
}
