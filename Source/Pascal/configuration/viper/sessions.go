package viper

import (
	"net/http"
	"time"

	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/sessions"
	"dolittle.io/pascal/sessions/nonces"
	"github.com/spf13/viper"
)

const (
	sessionsNonceLengthKey = "sessions.nonce_length"
	sessionsKeysKey        = "sessions.keys"
	sessionsLifetimeKey    = "sessions.lifetime"
	sessionsCookiesKey     = "sessions.cookies"

	defaultSessionsNonceLength          = 18
	defaultSessionsLifetime             = 5 * time.Minute
	defeaultSessionsCookiesName         = ".dolittle.pascal.session"
	defeaultSessionsCookiesSameSiteMode = http.SameSiteLaxMode
	defeaultSessionsCookiesPath         = "/"
)

type sessionsConfiguration struct {
	nonce   *nonceConfiguration
	cookies *cookiesConfiguration
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

func (c *sessionsConfiguration) Lifetime() time.Duration {
	if viper.IsSet(sessionsLifetimeKey) {
		return viper.GetDuration(sessionsLifetimeKey)
	}
	return defaultSessionsLifetime
}

func (c *sessionsConfiguration) Nonce() nonces.Configuration {
	return c.nonce
}

func (c *sessionsConfiguration) Cookies() cookies.Configuration {
	return c.cookies
}

type nonceConfiguration struct{}

func (c *nonceConfiguration) Size() int {
	if length := viper.GetInt(sessionsNonceLengthKey); length > 0 {
		return length
	}
	return defaultSessionsNonceLength
}
