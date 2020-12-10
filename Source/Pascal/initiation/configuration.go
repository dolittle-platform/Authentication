package initiation

import "net/url"

type MatchMode int

const (
	MatchModeStrict MatchMode = iota
	MatchModePrefix
)

type Configuration interface {
	ReturnToParameter() string
	DefaultReturnTo() *url.URL
	AllowedReturnTo() []*url.URL
	ReturnToMatchMode() MatchMode
}
