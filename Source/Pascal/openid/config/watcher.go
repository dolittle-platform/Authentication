package config

import (
	"time"

	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/issuer"
	"go.uber.org/zap"
)

type Watcher interface {
	GetIssuer() (issuer.Issuer, error)
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
	provider      issuer.Issuer
	changed       chan struct{}
	logger        *zap.Logger
}

func (w *watcher) GetIssuer() (issuer.Issuer, error) {
	if w.provider == nil {
		return nil, ErrOpenIDIssuerNotReady
	}
	return w.provider, nil
}

func (w *watcher) configLoop() {
	for {
		provider, err := issuer.NewIssuer(
			w.configuration.Issuer(),
			w.configuration.ClientID(),
			w.configuration.ClientSecret(),
			w.configuration.Scopes(),
			w.configuration.RedirectURL())

		if err != nil {
			w.logger.Warn("OpenID Connect issuer error, not ready to serve requests", zap.Error(err))

			select {
			case <-time.After(10 * time.Second):
			case <-w.changed:
			}
		} else {
			w.logger.Info("OpenID Connect issuer ready")
			w.provider = provider

			<-w.changed
			w.provider = nil
		}
	}
}

func (w *watcher) handleConfigurationChanged() error {
	w.changed <- struct{}{}
	return nil
}
