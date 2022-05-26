package openid

import (
	"dolittle.io/pascal/configuration/changes"
	"dolittle.io/pascal/openid/config"
	"dolittle.io/pascal/openid/issuer"
	"go.uber.org/zap"
)

type TokenRevoker interface {
	Revoke(token *issuer.Token) error
}

func NewTokenRevoker(configuration config.Configuration, notifier changes.ConfigurationChangeNotifier, logger *zap.Logger) (TokenRevoker, error) {
	watcher, err := config.NewWatcher(configuration, notifier, logger, "openid-revoker")
	if err != nil {
		return nil, err
	}
	return &revoker{
		configuration: configuration,
		watcher:       watcher,
		logger:        logger,
	}, nil
}

type revoker struct {
	configuration config.Configuration
	watcher       config.Watcher
	logger        *zap.Logger
}

func (r revoker) Revoke(token *issuer.Token) error {
	issuer, err := r.watcher.GetIssuer()
	if err != nil {
		return err
	}

	if !issuer.RevocationIsSupported() {
		r.logger.Info("Revocation is not supported by the issuer, not revoking the token")
		return nil
	}

	err = issuer.RevokeToken(token)
	if err != nil {
		r.logger.Error("Failed to revoke token from issuer", zap.Error(err))
	} else {
		r.logger.Debug("Successfully revoked token from issuer")
	}

	return err
}
