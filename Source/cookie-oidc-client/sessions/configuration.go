package sessions

import "dolittle.io/cookie-oidc-client/sessions/nonces"

// Configuration for sessions
type Configuration interface {
	// CookieName returns the name of the cookie to store sessions in
	CookieName() string

	// Nonce returns the configuration for nonce generation
	Nonce() nonces.Configuration
}
