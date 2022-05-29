package redirects

import "net/url"

type MatchMode int

const (
	MatchModeStrict MatchMode = iota
	MatchModePrefix
)

type Configuration interface {
	ReturnToParameter() string
	DefaultLoginReturnTo() *url.URL
	DefaultLogoutReturnTo() *url.URL
	AllowedReturnTo() []*url.URL
	ReturnToMatchMode() MatchMode
}
