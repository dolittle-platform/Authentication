package sessions

import (
	"net/url"

	"dolittle.io/cookie-oidc-client/sessions/nonce"
	"go.uber.org/zap"
)

// ReturnToURL represents a URL to return to after the OIDC flow is completed
type ReturnToURL *url.URL

// Session represents an OIDC flow session
type Session struct {
	// Nonce is the unique nonce tied to this session
	Nonce nonce.Nonce

	// ReturnTo defines where to redirect the browser after the OIDC flow is completed
	ReturnTo ReturnToURL
}

// Creator creates new sessions
type Creator interface {
	// NewSession returns a new Session
	NewSession(returnTo ReturnToURL) (*Session, error)
}

// NewCreator returns a new Creator
func NewCreator(generator nonce.Generator, logger zap.Logger) Creator {
	return &creator{
		generator: generator,
		logger:    logger,
	}
}

type creator struct {
	generator nonce.Generator
	logger    zap.Logger
}

func (c *creator) NewSession(returnTo ReturnToURL) (*Session, error) {
	nonce, err := c.generator.Generate()
	if err != nil {
		c.logger.Error("could not create new session", zap.Error(err))
		return nil, err
	}

	return &Session{
		Nonce:    nonce,
		ReturnTo: returnTo,
	}, nil
}
