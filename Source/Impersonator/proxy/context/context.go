package context

import (
	"time"

	"dolittle.io/impersonator/identities"
)

type Context struct {
	Identity       *identities.Identity
	RequestStarted time.Time
}
