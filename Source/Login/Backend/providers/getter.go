package providers

type Getter interface {
	GetProviderByID(id string) (Provider, error)
}
