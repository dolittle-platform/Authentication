package cookies

import (
	"dolittle.io/pascal/openid/issuer"
	"net/http"
)

type Reader interface {
	ReadTokenCookie(r *http.Request) (*issuer.Token, error)
}

func NewReader(configuration Configuration) Reader {
	return &reader{
		configuration: configuration,
	}
}

type reader struct {
	configuration Configuration
}

func (r *reader) ReadTokenCookie(request *http.Request) (*issuer.Token, error) {
	cookie, err := request.Cookie(r.configuration.Name())
	if err != nil {
		return nil, err
	}

	return &issuer.Token{
		Value:   cookie.Value,
		Expires: cookie.Expires,
	}, nil
}
