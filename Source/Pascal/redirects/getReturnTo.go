package redirects

import (
	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

func GetReturnToURL(configuration Configuration, defaultValue sessions.ReturnToURL, r *http.Request, logger *zap.Logger) (sessions.ReturnToURL, error) {
	returnToFromQueryString := r.URL.Query().Get(configuration.ReturnToParameter())
	if returnToFromQueryString != "" {
		returnToFromQuery, err := url.Parse(returnToFromQueryString)
		if err != nil {
			logger.Error("return to from query is not a valid URL", zap.Error(err))
			return nil, ErrRequestedReturnToWasNotValidURL
		}

		return returnToFromQuery, nil
	}

	return defaultValue, nil
}
