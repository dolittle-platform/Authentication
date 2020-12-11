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

// Use tenant mapper
func (g *getter) GetTenantFromID(tenantID TenantID) (*Tenant, error) {
	tenantNamesMap := g.configuration.TenantNamesMap()
	displayName := string(tenantID)
	if value, ok := (*tenantNamesMap)[tenantID]; ok {
		displayName = value
	}

	return &Tenant{
		ID:      tenantID,
		Display: displayName,
	}, nil
}
