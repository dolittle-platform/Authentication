package tenants

type Getter interface {
	GetTenantFromID(tenantID TenantID) (*Tenant, error)
}

func NewGetter() Getter {
	return &getter{}
}

type getter struct{}

func (g *getter) GetTenantFromID(tenantID TenantID) (*Tenant, error) {
	return &Tenant{
		ID:      tenantID,
		Display: string(tenantID),
	}, nil
}
