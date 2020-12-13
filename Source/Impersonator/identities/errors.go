package identities

import "errors"

var (
	ErrNoUserID          = errors.New("no user id set in headers")
	ErrMultipleUserIDs   = errors.New("multiple user ids set in headers")
	ErrNoTenantID        = errors.New("no tenant id set in headers")
	ErrMultipleTenantIDs = errors.New("multiple tenant ids set in headers")
)
