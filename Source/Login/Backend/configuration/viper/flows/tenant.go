package flows

import "github.com/spf13/viper"

const (
	tenantFlowIDQueryParameterKey = "flows.tenant.flow_id_query_parameter"

	defaultTenantFlowIDQueryParameter = "login_challenge"
)

type Tenant struct{}

func (t *Tenant) FlowIDQueryParameter() string {
	if value := viper.GetString(tenantFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultTenantFlowIDQueryParameter
}
