package config

import (
	"context"
	"net/url"
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
		issuer, query := w.splitIssuerUrlAndQuery(w.configuration.Issuer())

		ctx := context.Background()
		if len(query) > 0 {
			client := getHttpClientFor(issuer, query)
			ctx = context.WithValue(ctx, oauth2.HTTPClient, client)
		}

		provider, err := oidc.NewProvider(ctx, issuer.String())
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

func (w *watcher) splitIssuerUrlAndQuery(issuer *url.URL) (*url.URL, url.Values) {
	if len(issuer.Query()) > 0 {
		return &url.URL{
			Scheme:  issuer.Scheme,
			Opaque:  issuer.Opaque,
			User:    issuer.User,
			Host:    issuer.Host,
			Path:    issuer.Path,
			RawPath: issuer.RawPath,
		}, issuer.Query()
	} else {
		return issuer, nil
	}
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
