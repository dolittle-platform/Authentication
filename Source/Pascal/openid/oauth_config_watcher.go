package openid

import (
	"context"

	"dolittle.io/pascal/configuration/changes"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type oauthConfigWatcher struct {
	configuration Configuration
	oauthConfig   *oauth2.Config
}

func newOauthConfigWatcher(configuration Configuration, notifier changes.ConfigurationChangeNotifier, component changes.ComponentName) (*oauthConfigWatcher, error) {
	watcher := &oauthConfigWatcher{
		configuration: configuration,
	}
	oauthConfig, err := watcher.getConfig()
	if err != nil {
		return nil, err
	}
	watcher.oauthConfig = oauthConfig
	notifier.RegisterCallback(component, watcher.handleConfigurationChanged)
	return watcher, nil
}

func (w *oauthConfigWatcher) handleConfigurationChanged() error {
	ouathConfig, err := w.getConfig()
	if err != nil {
		return err
	}
	w.oauthConfig = ouathConfig
	return nil
}

func (w *oauthConfigWatcher) getConfig() (*oauth2.Config, error) {
	provider, err := oidc.NewProvider(context.Background(), w.configuration.Issuer().String())
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     w.configuration.ClientID(),
		ClientSecret: w.configuration.ClientSecret(),
		Endpoint:     provider.Endpoint(),
		Scopes:       w.getScopesWithOpenID(),
		RedirectURL:  w.configuration.RedirectURL().String(),
	}, nil
}

func (w *oauthConfigWatcher) getScopesWithOpenID() []string {
	scopes := w.configuration.Scopes()
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			return scopes
		}
	}
	return append(scopes, oidc.ScopeOpenID)
}
