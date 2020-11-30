package initiation

import (
	"dolittle.io/cookie-oidc-client/openid"
	"dolittle.io/cookie-oidc-client/sessions"
	"go.uber.org/zap"
)

type Initiatior interface {
	Initiate(request *Request) (*sessions.Session, openid.AuthenticationRedirectURL, error)
}

func NewInitiator(validator Validator, creator sessions.Creator, openid openid.AuthenticationInitiator, logger *zap.Logger) Initiatior {
	return &initiator{
		validator: validator,
		creator:   creator,
		initiator: openid,
		logger:    logger,
	}
}

type initiator struct {
	validator Validator
	creator   sessions.Creator
	initiator openid.AuthenticationInitiator
	logger    *zap.Logger
}

func (i *initiator) Initiate(request *Request) (*sessions.Session, openid.AuthenticationRedirectURL, error) {
	if isValid, err := i.validator.Validate(request); !isValid {
		return nil, "", err
	}

	session, err := i.creator.NewSession(request.ReturnTo)
	if err != nil {
		return nil, "", err
	}

	redirect := i.initiator.GetAuthenticationRedirect(session.Nonce)
	return session, redirect, nil
}
