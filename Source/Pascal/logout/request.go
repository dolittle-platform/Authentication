package logout

import (
	"dolittle.io/pascal/openid/issuer"
	"dolittle.io/pascal/sessions"
)

type Request struct {
	Token    *issuer.Token
	ReturnTo sessions.ReturnToURL
}
