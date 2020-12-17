package proxy

import "errors"

var (
	ErrNoRequestSetInResponse    = errors.New("no request set in response")
	ErrNoContextStoredForRequest = errors.New("no context stored for request")
)
