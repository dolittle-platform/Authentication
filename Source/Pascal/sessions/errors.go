package sessions

import "errors"

var (
	// ErrSessionIsMissingNonce reports that a stored session is missing the nonce value
	ErrSessionIsMissingNonce = errors.New("session is missing nonce")

	// ErrSessionIsMissingReturnTo reports that a stored session is missing the return to value
	ErrSessionIsMissingReturnTo = errors.New("session is missing return to")

	// ErrNonceWasNotAString reports that a stored session nonce value was not a string
	ErrNonceWasNotAString = errors.New("session nonce value was not a string")

	// ErrReturnToWasNotAURL reports that a stored session return to value was not a URL
	ErrReturnToWasNotAURL = errors.New("session return to was not a valid URL")
)
