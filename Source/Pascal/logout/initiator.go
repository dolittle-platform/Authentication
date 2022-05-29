package logout

import (
	"dolittle.io/pascal/openid"
	"go.uber.org/zap"
	"net/url"
)

type Initiator interface {
	Initiate(request *Request) (*url.URL, error)
}

func NewInitiator(revoker openid.TokenRevoker, logger *zap.Logger) Initiator {
	return &initiator{
		revoker: revoker,
		logger:  logger,
	}
}

type initiator struct {
	revoker openid.TokenRevoker
	logger  *zap.Logger
}

func (i *initiator) Initiate(request *Request) (*url.URL, error) {
	if request.Token == nil {
		i.logger.Info("no token found in logout request, not revoking")
	} else {
		if err := i.revoker.Revoke(request.Token); err != nil {
			return nil, err
		}
	}

	return request.ReturnTo, nil
}
