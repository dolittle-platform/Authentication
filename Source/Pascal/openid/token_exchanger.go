package openid

import (
	"context"

	"dolittle.io/pascal/configuration/changes"
	"golang.org/x/oauth2"
)

type AuthenticationCode string

type TokenExchanger interface {
	Exchange(code AuthenticationCode) (*oauth2.Token, error)
}

func NewTokenExchanger(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (TokenExchanger, error) {
	watcher, err := newOauthConfigWatcher(configuration, notifier, "openid-exchanger")
	if err != nil {
		return nil, err
	}
	return &exchanger{watcher}, nil
}

type exchanger struct {
	*oauthConfigWatcher
}

func (e *exchanger) Exchange(code AuthenticationCode) (*oauth2.Token, error) {
	return e.oauthConfig.Exchange(context.Background(), string(code))
}
