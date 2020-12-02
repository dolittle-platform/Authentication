package viper

import (
	"dolittle.io/cookie-oidc-client/sessions"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
	"github.com/spf13/viper"
)

const (
	sessionsNonceLengthKey = "sessions.nonce_length"
	sessionsCookieNameKey  = "sessions.cookie.name"
	sessionsKeysKey        = "sessions.keys"

	defaultSessionsNonceLength = 18
	defeaultSessionsCookieName = ".cookie-oidc-client.session"
)

type sessionsConfiguration struct {
	nonce *nonceConfiguration
}

func (c *sessionsConfiguration) CookieName() string {
	if name := viper.GetString(sessionsCookieNameKey); name != "" {
		return name
	}
	return defeaultSessionsCookieName
}

type configurationSessionKey struct {
	Hash  string
	Block string
}

func (c *sessionsConfiguration) EncryptionKeys() []sessions.SessionEncryptionKey {
	keys := make([]sessions.SessionEncryptionKey, 0)
	config := make([]configurationSessionKey, 0)
	if err := viper.UnmarshalKey(sessionsKeysKey, &config); err == nil {
		for _, key := range config {
			if key.Block == "" || key.Hash == "" {
				return nil
			}
			keys = append(keys, sessions.SessionEncryptionKey{
				BlockKey: []byte(key.Block),
				HashKey:  []byte(key.Hash),
			})
		}
	}
	return keys
}

func (c *sessionsConfiguration) Nonce() nonces.Configuration {
	return c.nonce
}

type nonceConfiguration struct{}

func (c *nonceConfiguration) Size() int {
	if length := viper.GetInt(sessionsNonceLengthKey); length > 0 {
		return length
	}
	return defaultSessionsNonceLength
}
