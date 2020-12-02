package sessions

import "dolittle.io/cookie-oidc-client/sessions/nonces"

type SessionEncryptionKey struct {
	HashKey  []byte
	BlockKey []byte
}

// Configuration for sessions
type Configuration interface {
	// CookieName returns the name of the cookie to store sessions in
	CookieName() string

	// EncryptionKeys returns the keys to use for session cookie encryption, decrytup, signing and verification
	EncryptionKeys() []SessionEncryptionKey

	// Nonce returns the configuration for nonce generation
	Nonce() nonces.Configuration
}
