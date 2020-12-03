package initiation

import (
	"net/url"

	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
)

type Validator interface {
	Validate(r *Request) (bool, error)
}

func NewValidator(configuration Configuration, logger *zap.Logger) Validator {
	return &validator{
		configuration: configuration,
		logger:        logger,
	}
}

type validator struct {
	configuration Configuration
	logger        *zap.Logger
}

func (v *validator) Validate(r *Request) (bool, error) {
	if !v.returnToURLIsAllowed(r.ReturnTo) {
		return false, ErrRequestedReturnToIsNotAllowed
	}

	return true, nil
}

func (v *validator) returnToURLIsAllowed(requested sessions.ReturnToURL) bool {
	for _, allowed := range v.configuration.AllowedReturnTo() {
		if urlEqualsSchemeHostPath(requested, allowed) {
			return true
		}
	}
	return false
}

func urlEqualsSchemeHostPath(left, right *url.URL) bool {
	return left.Scheme == right.Scheme && left.Host == right.Host && left.Path == right.Path
}
