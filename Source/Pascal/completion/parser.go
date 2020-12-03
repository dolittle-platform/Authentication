package completion

import (
	"net/http"

	"dolittle.io/pascal/openid"
	"dolittle.io/pascal/sessions/nonces"
	"go.uber.org/zap"
)

const (
	codeParameter  = "code"
	stateParameter = "state"
)

type Parser interface {
	ParseFrom(r *http.Request) (*Response, error)
}

func NewParser(logger *zap.Logger) Parser {
	return &parser{
		logger: logger,
	}
}

type parser struct {
	logger *zap.Logger
}

func (p *parser) ParseFrom(r *http.Request) (*Response, error) {
	code := r.URL.Query().Get(codeParameter)
	if code == "" {
		p.logger.Error("code was empty")
		return nil, ErrMissingCodeParameter
	}

	state := r.URL.Query().Get(stateParameter)

	return &Response{
		Code:  openid.AuthenticationCode(code),
		State: nonces.Nonce(state),
	}, nil
}
