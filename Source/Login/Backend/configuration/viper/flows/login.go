package flows

import "github.com/spf13/viper"

const (
	loginFlowIDQueryParameterKey            = "flows.login.flow_id_query_parameter"
	loginFlowCSRFTokenFieldNameParameterKey = "flows.login.csrf_token_parameter"
	loginFlowProviderFieldNameKey           = "flows.login.provider_parameter"

	defaultLoginFlowIDQueryParameter            = "id"
	defaultLoginFlowCSRFTokenFieldNameParameter = "csrf_token"
	defaultLoginFlowProviderFieldName           = "provider"
)

type Login struct{}

func (l *Login) FlowIDQueryParameter() string {
	if value := viper.GetString(loginFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultLoginFlowIDQueryParameter
}

func (l *Login) CSRFTokenFieldName() string {
	if value := viper.GetString(loginFlowCSRFTokenFieldNameParameterKey); value != "" {
		return value
	}
	return defaultLoginFlowCSRFTokenFieldNameParameter
}

func (l *Login) ProviderFieldName() string {
	if value := viper.GetString(loginFlowProviderFieldNameKey); value != "" {
		return value
	}
	return defaultLoginFlowProviderFieldName
}
