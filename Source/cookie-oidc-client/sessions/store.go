package sessions

import (
	"errors"
	"time"

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
			MaxAge:   int(configuration.Lifetime() / time.Second),
			Secure:   configuration.Cookies().Secure(),
			SameSite: configuration.Cookies().SameSite(),
			Path:     configuration.Cookies().Path(),
			HttpOnly: true,
		},
	}
	store.MaxAge(store.Options.MaxAge)

	notifier.RegisterCallback("cookie-store", func() error {
		store.Options.MaxAge = int(configuration.Lifetime() / time.Second)
		store.Options.Path = configuration.Cookies().Path()
		store.Options.Secure = configuration.Cookies().Secure()
		store.Options.SameSite = configuration.Cookies().SameSite()

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
