package openid

import (
	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/openid/issuer"
	"dolittle.io/pascal/sessions"
	"dolittle.io/pascal/sessions/nonces"
	"go.uber.org/zap"
	"net/url"
)

type AuthenticationRedirectURL string

type AuthenticationInitiator interface {
	GetAuthenticationRedirect(host string, nonce nonces.Nonce) (AuthenticationRedirectURL, error)
	GetLogoutRedirect(token *issuer.Token, returnTo sessions.ReturnToURL) (AuthenticationRedirectURL, error)
}

func NewAuthenticationInitiator(configuration config.Configuration, notifier changes.ConfigurationChangeNotifier, logger *zap.Logger) (AuthenticationInitiator, error) {
	watcher, err := config.NewWatcher(configuration, notifier, logger, "openid-initiator")
	if err != nil {
		return nil, err
	}
	return &initiator{
		watcher: watcher,
	}, nil
}

type initiator struct {
	watcher config.Watcher
}

func (i *initiator) GetAuthenticationRedirect(host string, nonce nonces.Nonce) (AuthenticationRedirectURL, error) {
	issuer, err := i.watcher.GetIssuer()
	if err != nil {
		return "", err
	}
	redirect, err := issuer.GetAuthenticationRedirectURL(host, nonce)
	if err != nil {
		return "", err
	}
	return AuthenticationRedirectURL(redirect), nil
}

func (i *initiator) GetLogoutRedirect(token *issuer.Token, returnTo sessions.ReturnToURL) (AuthenticationRedirectURL, error) {
	issuer, err := i.watcher.GetIssuer()
	if err != nil {
		return "", err
	}

	redirect := (*url.URL).String(returnTo)
	idTokenHint := ""
	if token != nil {
		idTokenHint = token.Value
	}

	if issuer.LogoutIsSupported() {
		redirect, err = issuer.GetLogoutRedirectURL(idTokenHint, "", redirect)
		if err != nil {
			return "", err
		}
	}

	return AuthenticationRedirectURL(redirect), nil
}
