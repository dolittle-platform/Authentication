package cookies

import (
	"net/http"
)

type Configuration interface {
	Name() string
	Secure() bool
	SameSite() http.SameSite
	Path() string
}
