package redirects

import "errors"

var (
	ErrRequestedReturnToWasNotValidURL = errors.New("requested return to was not a valid URL")
	ErrRequestedReturnToIsNotAllowed   = errors.New("requested return to URL is not allowed")
)
