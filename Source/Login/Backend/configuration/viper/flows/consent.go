package flows

import (
	"github.com/spf13/viper"
)

const (
	consentFlowIDQueryParameterKey = "flows.consent.flow_id_query_parameter"

	defaultConsentFlowIDQueryParameter = "consent_challenge"
)

type Consent struct{}

func (c *Consent) FlowIDQueryParameter() string {
	if value := viper.GetString(consentFlowIDQueryParameterKey); value != "" {
		return value
	}
	return defaultConsentFlowIDQueryParameter
}
