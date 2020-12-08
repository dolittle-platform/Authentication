package tenant

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"dolittle.io/login/clients/hydra"
	flowContext "dolittle.io/login/flows/context"
	"dolittle.io/login/identities/tenants"
	"github.com/ory/hydra-client-go/models"
)

type Selecter interface {
	SelectTenant(ctx context.Context, flow *Flow, tenant tenants.TenantID) (*url.URL, error)
	SelectTenantFrom(ctx context.Context, flow *Flow, r *http.Request) (*url.URL, error)
}

func NewSelecter(configuration Configuration, hydra hydra.Client) Selecter {
	return &selecter{
		configuration: configuration,
		hydra:         hydra,
	}
}

type selecter struct {
	configuration Configuration
	hydra         hydra.Client
}

func (s *selecter) SelectTenant(ctx context.Context, flow *Flow, tenant tenants.TenantID) (*url.URL, error) {
	if !flow.User.HasAccessToTenant(tenant) {
		return nil, errors.New("user doesn't have access to that tenant")
	}

	body := &models.AcceptLoginRequest{
		Subject:  &flow.User.Subject,
		Remember: false,
	}

	flowContext.StoreIn(body, &flowContext.Context{
		User:           flow.User,
		SelectedTenant: tenant,
	})

	response, err := s.hydra.AcceptLoginRequest(ctx, string(flow.ID), body)
	if err != nil {
		return nil, err
	}

	redirect, err := url.Parse(*response.RedirectTo)
	if err != nil {
		return nil, err
	}
	return redirect, nil
}

func (s *selecter) SelectTenantFrom(ctx context.Context, flow *Flow, r *http.Request) (*url.URL, error) {
	// TODO: Make parsing of form nice in one place
	tenantID := r.Form.Get(s.configuration.FlowTenantFormParameter())
	if tenantID == "" {
		return nil, errors.New("no tenant provided in POST form")
	}
	return s.SelectTenant(ctx, flow, tenants.TenantID(tenantID))
}
