package completion

import "errors"

var (
	ErrMissingCodeParameter                = errors.New("callback request was missing code parameter")
	ErrStateDoesNotMatchSession            = errors.New("callback state does not match session")
	ErrSessionDoesNotMatchProviderCallback = errors.New("session does not match provider callback")
)
