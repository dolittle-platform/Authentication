package tenant

import (
	"net/http"
	"net/url"

	"dolittle.io/login/identities/tenants"
)

type Selecter interface {
	SelectTenant(flow *Flow, tenant tenants.Tenant) (*url.URL, error)
	SelectTenantFrom(flow *Flow, r *http.Request) (*url.URL, error)
}

func NewSelecter() Selecter {
	return &selecter{}
}

type selecter struct{}

func (s *selecter) SelectTenant(flow *Flow, tenant tenants.Tenant) (*url.URL, error) {

}

func (s *selecter) SelectTenantFrom(flow *Flow, r *http.Request) (*url.URL, error) {

}
