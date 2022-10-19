package redirects

type MatchMode int

const (
	MatchModeStrict MatchMode = iota
	MatchModePrefix
)

type Configuration interface {
	ReturnToParameter() string
	DefaultLoginReturnTo() string
	DefaultLogoutReturnTo() string
	AllowedReturnTo() []string
	ReturnToMatchMode() MatchMode
}
