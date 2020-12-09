package providers

type Getter interface {
	GetProviderByID(id string) (Provider, error)
}

func NewGetter() Getter {
	return &getter{}
}

type getter struct{}

func (g *getter) GetProviderByID(id string) (Provider, error) {
	return Provider{
		ID:      id,
		Display: id + " Display Name",
	}, nil
}
