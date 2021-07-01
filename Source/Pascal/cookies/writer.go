package cookies

import (
	"net/http"

	"dolittle.io/pascal/openid/issuer"
)

type Writer interface {
	WriteTokenCookie(token *issuer.Token, w http.ResponseWriter) error
}

func NewWriter(configuration Configuration) Writer {
	return &writer{
		configuration: configuration,
	}
}

type writer struct {
	configuration Configuration
}

func (w *writer) WriteTokenCookie(token *issuer.Token, responseWriter http.ResponseWriter) error {
	cookie := http.Cookie{
		Name:     w.configuration.Name(),
		Value:    token.Value,
		Expires:  token.Expires,
		HttpOnly: true,
		Secure:   w.configuration.Secure(),
		SameSite: w.configuration.SameSite(),
		Path:     w.configuration.Path(),
	}

	http.SetCookie(responseWriter, &cookie)

	return nil
}
