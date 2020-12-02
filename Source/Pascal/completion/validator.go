package completion

import (
	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
)

type Validator interface {
	Validate(response *Response, session *sessions.Session) (bool, error)
}

func NewValidator(logger *zap.Logger) Validator {
	return &validator{
		logger: logger,
	}
}

type validator struct {
	logger *zap.Logger
}

func (v *validator) Validate(response *Response, session *sessions.Session) (bool, error) {
	if response.State != session.Nonce {
		v.logger.Error("state nonce does not match session nonce")
		return false, ErrStateDoesNotMatchSession
	}

	return true, nil
}
