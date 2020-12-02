package openid

import (
	"dolittle.io/cookie-oidc-client/configuration/changes"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
)

type AuthenticationRedirectURL string

type AuthenticationInitiator interface {
	GetAuthenticationRedirect(nonce nonces.Nonce) AuthenticationRedirectURL
}

func NewAuthenticationInitiator(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (AuthenticationInitiator, error) {
	watcher, err := newOauthConfigWatcher(configuration, notifier, "openid-initiator")
	if err != nil {
		return nil, err
	}
	return &initiator{watcher}, nil
}

type initiator struct {
	*oauthConfigWatcher
}

func (i *initiator) GetAuthenticationRedirect(nonce nonces.Nonce) AuthenticationRedirectURL {
	return AuthenticationRedirectURL(i.oauthConfig.AuthCodeURL(string(nonce)))
}
