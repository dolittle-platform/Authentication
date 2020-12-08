package current

import "errors"

var (
	ErrKratosTraitsWasNotStringMap       = errors.New("traits received from Kratos was not a string map")
	ErrKratosTraitsDoesNotContainTenants = errors.New("traits received from Kratos does not contain 'tenants'")
	ErrKratosTenantsWasNotArray          = errors.New("tenants recevied from Kratos was not a slice")
	ErrKratosTenantWasNotString          = errors.New("tenant recevied from Kratos was not a string")

	ErrNoUserLoggedIn = errors.New("user isn't logged in")
)
