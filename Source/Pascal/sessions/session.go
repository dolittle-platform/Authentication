package sessions

import (
	"net/url"

	"dolittle.io/pascal/sessions/nonces"
)

// ReturnToURL represents a URL to return to after the OIDC flow is completed
type ReturnToURL *url.URL

// Session represents an OIDC flow session
type Session struct {
	// Nonce is the unique nonce tied to this session
	Nonce nonces.Nonce

	// ReturnTo defines where to redirect the browser after the OIDC flow is completed
	ReturnTo ReturnToURL
}
