package cookies

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Writer interface {
	WriteTokenCookie(token *oauth2.Token, w http.ResponseWriter) error
}

func NewWriter(configuration Configuration) Writer {
	return &writer{
		configuration: configuration,
	}
}

type writer struct {
	configuration Configuration
}

func (w *writer) WriteTokenCookie(token *oauth2.Token, responseWriter http.ResponseWriter) error {
	cookie := http.Cookie{
		Name:     w.configuration.Name(),
		Value:    token.AccessToken,
		Expires:  token.Expiry,
		HttpOnly: true,
		Secure:   w.configuration.Secure(),
		SameSite: w.configuration.SameSite(),
		Path:     w.configuration.Path(),
	}

	http.SetCookie(responseWriter, &cookie)

	return nil
}
