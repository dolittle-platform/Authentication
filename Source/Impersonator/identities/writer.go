package identities

import "net/http"

const (
	ImpersonateUserHeader  = "Impersonate-User"
	ImpersonateGroupHeader = "Impersonate-Group"
)

type Writer interface {
	WriteIdentityTo(identity *Identity, r *http.Request) error
}

func NewWriter(configuration Configuration) Writer {
	return &writer{
		configuration: configuration,
	}
}

type writer struct {
	configuration Configuration
}

func (w *writer) WriteIdentityTo(identity *Identity, r *http.Request) error {
	r.Header.Set(ImpersonateUserHeader, identity.UserID)
	r.Header.Set(ImpersonateGroupHeader, "tenant-"+identity.TenantID)
	return nil
}
