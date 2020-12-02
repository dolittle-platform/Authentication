package sessions

import (
	"time"

	"dolittle.io/cookie-oidc-client/cookies"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
)

type SessionEncryptionKey struct {
	HashKey  []byte
	BlockKey []byte
}

// Configuration for sessions
type Configuration interface {
	// Lifetime returns the lifetime of authentication sessions
	Lifetime() time.Duration

	// Cookie returns the configuration for the session cookies
	Cookies() cookies.Configuration

	// EncryptionKeys returns the keys to use for session cookie encryption, decrytup, signing and verification
	EncryptionKeys() []SessionEncryptionKey

	// Nonce returns the configuration for nonce generation
	Nonce() nonces.Configuration
}
