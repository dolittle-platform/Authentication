package logout

import (
	"dolittle.io/pascal/cookies"
	"dolittle.io/pascal/redirects"
	"go.uber.org/zap"
	"net/http"
)

type Parser interface {
	ParseFrom(r *http.Request) (*Request, error)
}

func NewParser(configuration redirects.Configuration, reader cookies.Reader, logger *zap.Logger) Parser {
	return &parser{
		configuration: configuration,
		reader:        reader,
		logger:        logger,
	}
}

type parser struct {
	configuration redirects.Configuration
	reader        cookies.Reader
	logger        *zap.Logger
}

func (p *parser) ParseFrom(r *http.Request) (*Request, error) {
	token, err := p.reader.ReadTokenCookie(r)
	if err != nil {
		p.logger.Info("no token found in logout request")
	}

	defaultReturnTo, err := redirects.GetAbsoluteUrlFor(r, p.configuration.DefaultLogoutReturnTo())
	if err != nil {
		return nil, err
	}

	returnTo, err := redirects.GetReturnToURL(p.configuration, defaultReturnTo, r, p.logger)
	if err != nil {
		return nil, err
	}

	return &Request{
		Token:    token,
		ReturnTo: returnTo,
	}, nil
}
