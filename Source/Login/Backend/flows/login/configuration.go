package login

type Configuration interface {
	FlowIdQueryParameter() string
	CSRFTokenFieldName() string
	ProviderFieldName() string
}
