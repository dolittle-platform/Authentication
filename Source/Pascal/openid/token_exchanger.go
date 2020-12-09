package openid

import (
	"context"

	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/config"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type AuthenticationCode string

type TokenExchanger interface {
	Exchange(code AuthenticationCode) (*oauth2.Token, error)
}

func NewTokenExchanger(configuration config.Configuration, notifier changes.ConfigurationChangeNotifier, logger *zap.Logger) (TokenExchanger, error) {
	watcher, err := config.NewWatcher(configuration, notifier, logger, "openid-exchanger")
	if err != nil {
		return nil, err
	}
	return &exchanger{
		watcher: watcher,
	}, nil
}

type exchanger struct {
	watcher config.Watcher
}

func (e *exchanger) Exchange(code AuthenticationCode) (*oauth2.Token, error) {
	config, err := e.watcher.GetConfig()
	if err != nil {
		return nil, err
	}
	return config.Exchange(context.Background(), string(code))
}
