package cookies

import (
	"net/http"
)

type Crumbler interface {
	DestroyTokenCookie(w http.ResponseWriter) error
}

func NewCrumbler(configuration Configuration) Crumbler {
	return &crumbler{
		configuration: configuration,
	}
}

type crumbler struct {
	configuration Configuration
}

func (c *crumbler) DestroyTokenCookie(responseWriter http.ResponseWriter) error {
	cookie := http.Cookie{
		Name:     c.configuration.Name(),
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   c.configuration.Secure(),
		SameSite: c.configuration.SameSite(),
		Path:     c.configuration.Path(),
	}

	http.SetCookie(responseWriter, &cookie)

	return nil
}
