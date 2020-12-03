package sessions

import (
	"dolittle.io/pascal/sessions/nonces"
	"go.uber.org/zap"
)

// Creator creates new sessions
type Creator interface {
	// NewSession returns a new Session
	NewSession(returnTo ReturnToURL) (*Session, error)
}

// NewCreator returns a new Creator
func NewCreator(generator nonces.Generator, logger *zap.Logger) Creator {
	return &creator{
		generator: generator,
		logger:    logger,
	}
}

type creator struct {
	generator nonces.Generator
	logger    *zap.Logger
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
