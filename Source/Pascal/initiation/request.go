package initiation

import "dolittle.io/pascal/sessions"

type Request struct {
	Host     string
	ReturnTo sessions.ReturnToURL
}
