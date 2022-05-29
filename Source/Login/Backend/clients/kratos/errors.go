package kratos

import "errors"

var (
	ErrKratosUnauthorized = errors.New("request returned 401 unauthorized")
)
