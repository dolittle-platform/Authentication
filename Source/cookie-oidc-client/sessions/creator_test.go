package sessions_test

import (
	"net/url"

	"dolittle.io/cookie-oidc-client/sessions"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Creator", func() {
	Context("given a generator, a logger and a return to URL", func() {
		var (
			generator *mockNonceGenerator
			logger    *zap.Logger
			returnTo  sessions.ReturnToURL
		)
		BeforeEach(func() {
			generator = &mockNonceGenerator{nonce: "lZehpZpPki"}
			logger = zap.NewNop()
			returnTo, _ = url.Parse("http://localhost:8080/return_to_url?with_params=yes")
		})

		When("creating a new session", func() {
			var (
				created *sessions.Session
				err     error
			)
			BeforeEach(func() {
				creator := sessions.NewCreator(generator, logger)
				created, err = creator.NewSession(returnTo)
			})
			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})
			It("should create a session with the generated nonce", func() {
				Expect(created.Nonce).To(Equal(generator.nonce))
			})
			It("should create a session with the return to URL", func() {
				Expect(created.ReturnTo).To(Equal(returnTo))
			})
		})
	})
})

type mockNonceGenerator struct {
	nonce nonces.Nonce
}

func (m *mockNonceGenerator) Generate() (nonces.Nonce, error) {
	return m.nonce, nil
}
