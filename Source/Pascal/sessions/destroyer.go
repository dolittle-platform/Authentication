package sessions

import (
	"net/http"

	gorilla "github.com/gorilla/sessions"
	"go.uber.org/zap"
)

type Destroyer interface {
	Destroy(r *http.Request, w http.ResponseWriter) error
}

func NewDestroyer(configuration Configuration, store gorilla.Store, logger *zap.Logger) Destroyer {
	return &destroyer{
		configuration: configuration,
		store:         store,
		logger:        logger,
	}
}

type destroyer struct {
	configuration Configuration
	store         gorilla.Store
	logger        *zap.Logger
}

func (d *destroyer) Destroy(r *http.Request, w http.ResponseWriter) error {
	cookie, err := d.store.Get(r, d.configuration.Cookies().Name())
	if err != nil {
		return nil
	}

	cookie.Options.MaxAge = -1
	err = d.store.Save(r, w, cookie)
	return err
}
