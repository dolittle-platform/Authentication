package cookies

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Writer interface {
	WriteTokenCookie(token *oauth2.Token, w http.ResponseWriter) error
}

func NewWriter() Writer {
	return &writer{}
}

type writer struct{}

func (w *writer) WriteTokenCookie(token *oauth2.Token, responseWriter http.ResponseWriter) error {
	cookie := http.Cookie{
		Name:    "cookie-name",
		Value:   token.AccessToken,
		Path:    "/",
		Expires: token.Expiry,
	}

	http.SetCookie(responseWriter, &cookie)

	return nil
}
