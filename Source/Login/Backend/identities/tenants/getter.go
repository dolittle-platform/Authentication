package tenants

type Getter interface {
	GetTenantFromID(tenantID TenantID) (*Tenant, error)
}

func NewGetter(configuration Configuration) Getter {
	return &getter{
		configuration,
	}
}

type getter struct {
	configuration Configuration
}

func (g *getter) GetTenantFromID(tenantID TenantID) (*Tenant, error) {
	tenant := &Tenant{
		ID:      tenantID,
		Display: string(tenantID),
	}

	if name, ok := g.configuration.TenantNames()[tenantID]; ok {
		tenant.Display = name
	}

	return tenant, nil
}
