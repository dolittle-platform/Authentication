package tenant

type Configuration interface {
	FlowIDQueryParameter() string
	FlowIDFormParameter() string
	FlowTenantFormParameter() string
}
