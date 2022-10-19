package initiation

import (
	"dolittle.io/pascal/redirects"
	"net/url"
	"strings"

	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
)

type Validator interface {
	Validate(r *Request) (bool, error)
}

func NewValidator(configuration redirects.Configuration, logger *zap.Logger) Validator {
	return &validator{
		configuration: configuration,
		logger:        logger,
	}
}

type validator struct {
	configuration redirects.Configuration
	logger        *zap.Logger
}

func (v *validator) Validate(r *Request) (bool, error) {
	if !v.returnToURLIsAllowed(r.ReturnTo, r.Host) {
		returnTo := url.URL(*r.ReturnTo)
		v.logger.Warn("the requested return to URL is not allowed", zap.String("requested", returnTo.String()))
		return false, redirects.ErrRequestedReturnToIsNotAllowed
	}

	return true, nil
}

func (v *validator) returnToURLIsAllowed(requested sessions.ReturnToURL, allowedHost string) bool {
	for _, allowedPath := range v.configuration.AllowedReturnTo() {
		if urlEqualsSchemeHostPath(requested, allowedHost, allowedPath, v.configuration.ReturnToMatchMode()) {
			return true
		}
	}
	return false
}

func urlEqualsSchemeHostPath(requested *url.URL, allowedHost, allowedPath string, mode redirects.MatchMode) bool {
	requestedHost := (&url.URL{Scheme: requested.Scheme, Host: requested.Host}).String()
	if requestedHost != allowedHost {
		return false
	}

	switch mode {
	case redirects.MatchModePrefix:
		return strings.HasPrefix(requested.Path, allowedPath)
	case redirects.MatchModeStrict:
		fallthrough
	default:
		return requested.Path == allowedPath
	}
}
