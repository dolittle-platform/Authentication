package sessions

import (
	"errors"

	"dolittle.io/cookie-oidc-client/configuration/changes"
	"github.com/gorilla/securecookie"
	gorilla "github.com/gorilla/sessions"
)

func NewCookieStore(configuration Configuration, notifier changes.ConfigurationChangeNotifier) (gorilla.Store, error) {
	codecs, err := createCodecsFromEncryptionKeys(configuration)
	if err != nil {
		return nil, err
	}

	store := &gorilla.CookieStore{
		Codecs: codecs,
		Options: &gorilla.Options{
			Path:     "/",
			MaxAge:   60 * 5,
			HttpOnly: true,
		},
	}
	store.MaxAge(store.Options.MaxAge)

	notifier.RegisterCallback("cookie-store", func() error {
		codecs, err := createCodecsFromEncryptionKeys(configuration)
		if err != nil {
			return err
		}

		store.Codecs = codecs
		return nil
	})

	return store, nil
}

func createCodecsFromEncryptionKeys(configuration Configuration) ([]securecookie.Codec, error) {
	keys := configuration.EncryptionKeys()
	if len(keys) < 1 {
		return nil, errors.New("NO KEYS")
	}

	codecs := make([]securecookie.Codec, 0)
	for _, key := range configuration.EncryptionKeys() {
		codecs = append(codecs, securecookie.New(key.HashKey, key.BlockKey))
	}

	return codecs, nil
}
