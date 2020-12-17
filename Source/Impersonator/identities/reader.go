package identities

import "net/http"

type Reader interface {
	ReadIdentityFrom(r *http.Request) (*Identity, error)
}

func NewReader(configuration Configuration) Reader {
	return &reader{
		configuration: configuration,
	}
}

type reader struct {
	configuration Configuration
}

func (r *reader) ReadIdentityFrom(req *http.Request) (*Identity, error) {
	userIDs := req.Header.Values(r.configuration.UserIDHeader())
	if len(userIDs) < 1 {
		return nil, ErrNoUserID
	}
	if len(userIDs) > 1 {
		return nil, ErrMultipleUserIDs
	}

	tenantIDs := req.Header.Values(r.configuration.TenantIDHeader())
	if len(tenantIDs) < 1 {
		return nil, ErrNoTenantID
	}
	if len(tenantIDs) > 1 {
		return nil, ErrMultipleTenantIDs
	}

	return &Identity{
		UserID:   userIDs[0],
		TenantID: tenantIDs[0],
	}, nil
}
