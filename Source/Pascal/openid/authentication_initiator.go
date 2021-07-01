package openid

import (
	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/sessions/nonces"
	"go.uber.org/zap"
)

type AuthenticationRedirectURL string

type AuthenticationInitiator interface {
	GetAuthenticationRedirect(nonce nonces.Nonce) (AuthenticationRedirectURL, error)
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

func (i *initiator) GetAuthenticationRedirect(nonce nonces.Nonce) (AuthenticationRedirectURL, error) {
	issuer, err := i.watcher.GetIssuer()
	if err != nil {
		return "", err
	}
	redirect, err := issuer.GetAuthenticationRedirectURL(nonce)
	if err != nil {
		return "", err
	}
	return AuthenticationRedirectURL(redirect), nil
}
