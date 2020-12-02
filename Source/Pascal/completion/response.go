package completion

import (
	"dolittle.io/pascal/openid"
	"dolittle.io/pascal/sessions/nonces"
)

type Response struct {
	Code  openid.AuthenticationCode
	State nonces.Nonce
}
