package identities

type Configuration interface {
	UserIDHeader() string
	TenantIDHeader() string
}
