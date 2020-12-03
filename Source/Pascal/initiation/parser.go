package initiation

import (
	"net/http"
	"net/url"

	"dolittle.io/pascal/sessions"
	"go.uber.org/zap"
)

type Parser interface {
	ParseFrom(r *http.Request) (*Request, error)
}

func NewParser(configuration Configuration, logger *zap.Logger) Parser {
	return &parser{
		configuration: configuration,
		logger:        logger,
	}
}

type parser struct {
	configuration Configuration
	logger        *zap.Logger
}

func (p *parser) ParseFrom(r *http.Request) (*Request, error) {
	returnTo, err := p.getReturnToURL(r)
	if err != nil {
		return nil, err
	}

	return &Request{
		ReturnTo: returnTo,
	}, nil
}

func (p *parser) getReturnToURL(r *http.Request) (sessions.ReturnToURL, error) {
	returnToFromQueryString := r.URL.Query().Get(p.configuration.ReturnToParameter())
	if returnToFromQueryString != "" {
		returnToFromQuery, err := url.Parse(returnToFromQueryString)
		if err != nil {
			p.logger.Error("return to from query is not a valid URL", zap.Error(err))
			return nil, ErrRequestedReturnToWasNotValidURL
		}

		return returnToFromQuery, nil
	}

	return p.configuration.DefaultReturnTo(), nil
}
