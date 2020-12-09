package tenant

import "net/url"

type Configuration interface {
	FlowIDQueryParameter() string
	FlowIDFormParameter() string
	FlowTenantFormParameter() string

	SelectTenantFormSubmitAction() *url.URL
}
