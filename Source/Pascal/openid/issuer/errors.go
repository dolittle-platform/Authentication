package issuer

import "errors"

var (
	ErrIDTokenMissing = errors.New("ID token not received in token from issuer")
)
