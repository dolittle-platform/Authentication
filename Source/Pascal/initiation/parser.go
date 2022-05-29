package initiation

import (
	"dolittle.io/pascal/redirects"
	"go.uber.org/zap"
	"net/http"
)

type Parser interface {
	ParseFrom(r *http.Request) (*Request, error)
}

func NewParser(configuration redirects.Configuration, logger *zap.Logger) Parser {
	return &parser{
		configuration: configuration,
		logger:        logger,
	}
}

type parser struct {
	configuration redirects.Configuration
	logger        *zap.Logger
}

func (p *parser) ParseFrom(r *http.Request) (*Request, error) {
	returnTo, err := redirects.GetReturnToURL(p.configuration, p.configuration.DefaultLoginReturnTo(), r, p.logger)
	if err != nil {
		return nil, err
	}

	return &Request{
		ReturnTo: returnTo,
	}, nil
}
