package config

import (
	"context"
	"time"

	"dolittle.io/pascal/configuration/changes"
	"github.com/coreos/go-oidc"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Watcher interface {
	GetConfig() (*oauth2.Config, error)
}

func NewWatcher(configuration Configuration, notifier changes.ConfigurationChangeNotifier, logger *zap.Logger, component changes.ComponentName) (Watcher, error) {
	w := &watcher{
		configuration: configuration,
		changed:       make(chan struct{}),
		logger:        logger,
	}
	if err := notifier.RegisterCallback(component, w.handleConfigurationChanged); err != nil {
		return nil, err
	}
	go w.configLoop()
	return w, nil
}

type watcher struct {
	configuration Configuration
	config        *oauth2.Config
	changed       chan struct{}
	logger        *zap.Logger
}

func (w *watcher) GetConfig() (*oauth2.Config, error) {
	if w.config == nil {
		return nil, ErrOpenIDIssuerNotReady
	}
	return w.config, nil
}

func (w *watcher) configLoop() {
	for {
		provider, err := oidc.NewProvider(context.Background(), w.configuration.Issuer().String())
		if err != nil {
			w.logger.Warn("OpenID Connect issuer error, not ready to serve requests", zap.Error(err))

			select {
			case <-time.After(10 * time.Second):
			case <-w.changed:
			}
		} else {
			w.logger.Info("OpenID Connect issuer ready")
			w.config = &oauth2.Config{
				ClientID:     w.configuration.ClientID(),
				ClientSecret: w.configuration.ClientSecret(),
				Endpoint:     provider.Endpoint(),
				Scopes:       w.getScopesWithOpenID(),
				RedirectURL:  w.configuration.RedirectURL().String(),
			}

			<-w.changed
			w.config = nil
		}
	}
}

func (w *watcher) handleConfigurationChanged() error {
	w.changed <- struct{}{}
	return nil
}

func (w *watcher) getScopesWithOpenID() []string {
	scopes := w.configuration.Scopes()
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			return scopes
		}
	}
	return append(scopes, oidc.ScopeOpenID)
}
