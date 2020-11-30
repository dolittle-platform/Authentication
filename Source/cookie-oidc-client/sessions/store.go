package sessions

import (
	gorilla "github.com/gorilla/sessions"
)

func NewCookieStore(configuration Configuration) gorilla.Store {
	store := &gorilla.CookieStore{
		Options: &gorilla.Options{
			Path:     "/",
			MaxAge:   60 * 5,
			HttpOnly: true,
		},
	}
	store.MaxAge(store.Options.MaxAge)
	return store
}
