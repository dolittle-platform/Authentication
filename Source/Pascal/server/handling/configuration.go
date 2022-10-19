package handling

type Configuration interface {
	AllowedHosts() []string
	ErrorRedirect() string
}
