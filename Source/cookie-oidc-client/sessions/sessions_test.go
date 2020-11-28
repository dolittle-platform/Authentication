package sessions_test

import (
	"testing"

	"dolittle.io/cookie-oidc-client/sessions/nonces"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestSessions(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "sessions")
}

type configuration struct {
	cookieName string
}

func (c *configuration) CookieName() string {
	return c.cookieName
}

func (*configuration) Nonce() nonces.Configuration {
	return nil
}
