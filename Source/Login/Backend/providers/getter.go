package providers

type Getter interface {
	GetProviderByID(id string) (Provider, error)
}

func NewGetter(configuration Configuration) Getter {
	return &getter{
		configuration,
	}
}

type getter struct {
	configuration Configuration
}

func (g *getter) GetProviderByID(id string) (Provider, error) {
	provider, ok := g.configuration.Providers()[id]
	if !ok {
		return Provider{
			ID:       id,
			Display:  id + " Display Name",
			ImageURL: "",
		}, nil
	}

	imageURLString := ""
	if provider.ImageURL != nil {
		imageURLString = provider.ImageURL.String()
	}

	return Provider{
		ID:       id,
		Display:  provider.Name,
		ImageURL: imageURLString,
	}, nil
}
