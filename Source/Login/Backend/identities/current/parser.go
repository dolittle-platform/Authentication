package current

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
	"github.com/ory/kratos-client-go/models"
)

type Parser interface {
	ParseUserFrom(session *models.Session) (*users.User, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct{}

func (p *parser) ParseUserFrom(session *models.Session) (*users.User, error) {
	tenants, err := p.getTenantsFromTraits(session.Identity.Traits)
	if err != nil {
		return nil, err
	}

	return &users.User{
		Subject: string(session.Identity.ID),
		Tenants: tenants,
	}, nil
}

func (p *parser) getTenantsFromTraits(traits models.Traits) ([]tenants.Tenant, error) {
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
		tenantString, ok := tenantValue.(string)
		if !ok {
			return nil, ErrKratosTenantWasNotString
		}
		userTenants = append(userTenants, tenants.Tenant(tenantString))
	}

	return userTenants, nil
}
