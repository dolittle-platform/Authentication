package flows

import (
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

const (
	tenantFlowIDQueryParameterKey    = "flows.tenant.flow_id_query_parameter"
	tenantFlowIDFormParameterKey     = "flows.tenant.flow_id_form_parameter"
	tenantFlowTenantFormParameterKey = "flows.tenant.flow_tenant_form_parameter"
	serveBaseURLKey                  = "serve.base_url"

	defaultTenantFlowIDQueryParameter    = "login_challenge"
	defaultTenantFlowIDFormParameter     = "login_challenge"
	defaultTenantFlowTenantFormParameter = "tenant"
)

var (
	defaultServeBaseURL = &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/",
	}
)

type Tenant struct{}

func (t *Tenant) FlowIDQueryParameter() string {
	if value := viper.GetString(tenantFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultTenantFlowIDQueryParameter
}

func (t *Tenant) FlowIDFormParameter() string {
	if value := viper.GetString(tenantFlowIDFormParameterKey); value != "" {
		return value
	}
	return defaultTenantFlowIDFormParameter
}

func (t *Tenant) FlowTenantFormParameter() string {
	if value := viper.GetString(tenantFlowTenantFormParameterKey); value != "" {
		return value
	}
	return defaultTenantFlowTenantFormParameter
}

func (t *Tenant) SelectTenantFormSubmitAction() *url.URL {
	baseURL := t.getBaseURL()
	// TODO: Proper clone of URL
	submitURL := &url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
	}
	if strings.HasSuffix(baseURL.Path, "/") {
		submitURL.Path = baseURL.Path + "self-service/tenant/select"
	} else {
		submitURL.Path = baseURL.Path + "/self-service/tenant/select"
	}
	return submitURL
}

func (t *Tenant) getBaseURL() *url.URL {
	baseURLValue := viper.GetString(serveBaseURLKey)
	if baseURLValue == "" {
		return defaultServeBaseURL
	}
	baseURL, err := url.Parse(baseURLValue)
	if err != nil {
		return defaultServeBaseURL
	}
	return baseURL
}
