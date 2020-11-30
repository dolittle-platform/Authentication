package completion

import (
	"dolittle.io/cookie-oidc-client/openid"
	"dolittle.io/cookie-oidc-client/sessions/nonces"
)

type Response struct {
	Code  openid.AuthenticationCode
	State nonces.Nonce
}
