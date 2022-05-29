package logout

import (
	"dolittle.io/pascal/openid"
	"go.uber.org/zap"
)

type Initiator interface {
	Initiate(request *Request) (openid.AuthenticationRedirectURL, error)
}

func NewInitiator(revoker openid.TokenRevoker, openidInitiator openid.AuthenticationInitiator, logger *zap.Logger) Initiator {
	return &initiator{
		revoker:   revoker,
		initiator: openidInitiator,
		logger:    logger,
	}
}

type initiator struct {
	revoker   openid.TokenRevoker
	initiator openid.AuthenticationInitiator
	logger    *zap.Logger
}

func (i *initiator) Initiate(request *Request) (openid.AuthenticationRedirectURL, error) {
	if request.Token == nil {
		i.logger.Info("no token found in logout request, not revoking")
	} else {
		if err := i.revoker.Revoke(request.Token); err != nil {
			return "", err
		}
	}

	return i.initiator.GetLogoutRedirect(nil, request.ReturnTo)
}
