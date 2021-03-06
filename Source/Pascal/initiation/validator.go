package initiation

import (
	"net/url"
	"strings"

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
		returnTo := url.URL(*r.ReturnTo)
		v.logger.Warn("the requested return to URL is not allowed", zap.String("requested", returnTo.String()))
		return false, ErrRequestedReturnToIsNotAllowed
	}

	return true, nil
}

func (v *validator) returnToURLIsAllowed(requested sessions.ReturnToURL) bool {
	for _, allowed := range v.configuration.AllowedReturnTo() {
		if urlEqualsSchemeHostPath(requested, allowed, v.configuration.ReturnToMatchMode()) {
			return true
		}
	}
	return false
}

func urlEqualsSchemeHostPath(requested, allowed *url.URL, mode MatchMode) bool {
	if requested.Scheme != allowed.Scheme || requested.Host != allowed.Host {
		return false
	}

	switch mode {
	case MatchModePrefix:
		return strings.HasPrefix(requested.Path, allowed.Path)
	case MatchModeStrict:
		fallthrough
	default:
		return requested.Path == allowed.Path
	}
}
