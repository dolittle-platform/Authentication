package flows

import "github.com/spf13/viper"

const (
	tenantFlowIDQueryParameterKey    = "flows.tenant.flow_id_query_parameter"
	tenantFlowIDFormParameterKey     = "flows.tenant.flow_id_form_parameter"
	tenantFlowTenantFormParameterKey = "flows.tenant.flow_tenant_form_parameter"

	defaultTenantFlowIDQueryParameter    = "login_challenge"
	defaultTenantFlowIDFormParameter     = "login_challenge"
	defaultTenantFlowTenantFormParameter = "tenant"
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
