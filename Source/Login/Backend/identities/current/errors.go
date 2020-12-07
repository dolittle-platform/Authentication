package current

import "errors"

var (
	ErrKratosTraitsWasNotStringMap       = errors.New("traits received from Kratos was not a map[string]interface{}")
	ErrKratosTraitsDoesNotContainTenants = errors.New("traits received from Kratos does not contain 'tenants'")
	ErrKratosTenantsWasNotStringSlice    = errors.New("tenants recevied from Kratos was not a []string")

	ErrNoUserLoggedIn = errors.New("user isn't logged in")
)
