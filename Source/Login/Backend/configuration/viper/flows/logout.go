package flows

import "github.com/spf13/viper"

const (
	logoutFlowIDQueryParameterKey = "flows.logout.flow_id_query_parameter"

	defaultLogoutFlowIDQueryParameter = "logout_challenge"
)

type Logout struct{}

func (l *Logout) FlowIDQueryParameter() string {
	if value := viper.GetString(logoutFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultLogoutFlowIDQueryParameter
}
