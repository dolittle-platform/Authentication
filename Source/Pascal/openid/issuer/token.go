package issuer

import "time"

type Token struct {
	Value   string
	Expires time.Time
}
