package context

import (
	"net/http"
	"time"

	"dolittle.io/impersonator/identities"
)

type Creator interface {
	CreateFor(r *http.Request) (*Context, error)
}

func NewCreator(reader identities.Reader) Creator {
	return &creator{
		reader: reader,
	}
}

type creator struct {
	reader identities.Reader
}

func (c *creator) CreateFor(r *http.Request) (*Context, error) {
	identity, err := c.reader.ReadIdentityFrom(r)
	if err != nil {
		return nil, err
	}

	return &Context{
		Identity:       identity,
		RequestStarted: time.Now(),
	}, nil
}
