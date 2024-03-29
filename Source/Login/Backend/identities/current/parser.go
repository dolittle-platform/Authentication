package current

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
	ory "github.com/ory/kratos-client-go"
)

type Parser interface {
	ParseUserFrom(session *ory.Session) (*users.User, error)
}

func NewParser(tenants tenants.Getter) Parser {
	return &parser{
		tenants: tenants,
	}
}

type parser struct {
	tenants tenants.Getter
}

func (p *parser) ParseUserFrom(session *ory.Session) (*users.User, error) {
	tenants, err := p.getTenantsFromTraits(session.Identity.Traits)
	if err != nil {
		return nil, err
	}

	return &users.User{
		Subject: session.Identity.Id,
		Tenants: tenants,
	}, nil
}

func (p *parser) getTenantsFromTraits(traits interface{}) ([]tenants.Tenant, error) {
	traitsMap, ok := traits.(map[string]interface{})
	if !ok {
		return nil, ErrKratosTraitsWasNotStringMap
	}

	tenantsValue, ok := traitsMap["tenants"]
	if !ok {
		return nil, ErrKratosTraitsDoesNotContainTenants
	}

	tenantsSlice, ok := tenantsValue.([]interface{})
	if !ok {
		return nil, ErrKratosTenantsWasNotArray
	}

	userTenants := make([]tenants.Tenant, 0)
	for _, tenantValue := range tenantsSlice {
		tenantID, ok := tenantValue.(string)
		if !ok {
			return nil, ErrKratosTenantWasNotString
		}
		tenant, err := p.tenants.GetTenantFromID(tenants.TenantID(tenantID))
		if err != nil {
			return nil, err
		}
		userTenants = append(userTenants, *tenant)
	}

	return userTenants, nil
}
