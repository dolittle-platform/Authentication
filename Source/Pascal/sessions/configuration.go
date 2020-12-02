package sessions

import (
	"time"

	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/sessions/nonces"
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
