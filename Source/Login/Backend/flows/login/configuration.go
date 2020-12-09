package login

type Configuration interface {
	FlowIDQueryParameter() string
	CSRFTokenFieldName() string
	ProviderFieldName() string
}
