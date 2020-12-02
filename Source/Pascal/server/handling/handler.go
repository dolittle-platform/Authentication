package handling

import (
	"context"
	"net/http"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error
}
