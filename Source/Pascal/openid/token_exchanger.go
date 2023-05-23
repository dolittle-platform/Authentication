package openid

import (
	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/openid/issuer"
	"go.uber.org/zap"
)

type AuthenticationCode string

type TokenExchanger interface {
	Exchange(host string, code AuthenticationCode) (*issuer.Token, error)
}

func NewTokenExchanger(configuration config.Configuration, notifier changes.ConfigurationChangeNotifier, logger *zap.Logger) (TokenExchanger, error) {
	watcher, err := config.NewWatcher(configuration, notifier, logger, "openid-exchanger")
	if err != nil {
		return nil, err
	}
	return &exchanger{
		configuration: configuration,
		watcher:       watcher,
		logger:        logger,
	}, nil
}

type exchanger struct {
	configuration config.Configuration
	watcher       config.Watcher
	logger        *zap.Logger
}

func (e *exchanger) Exchange(host string, code AuthenticationCode) (*issuer.Token, error) {
	issuer, err := e.watcher.GetIssuer()
	if err != nil {
		return nil, err
	}

	switch e.configuration.TokenType() {
	case config.IDToken:
		e.logger.Debug("Exchanging code for id token")
		return issuer.ExchangeCodeForIDToken(string(code))
	case config.AccessToken:
		e.logger.Debug("Exchanging code for access token")
		return issuer.ExchangeCodeForAccessToken(host, string(code))
	default:
		e.logger.Warn("Invalid token type configured, falling back to access token", zap.String("token_type", string(e.configuration.TokenType())))
		return issuer.ExchangeCodeForAccessToken(host, string(code))
	}
}
