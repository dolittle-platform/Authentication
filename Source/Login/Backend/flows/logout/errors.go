package logout

import "errors"

var (
	ErrLogoutChallengeNotFound = errors.New("no logout flow challenge from Hydra found on the request")
)
