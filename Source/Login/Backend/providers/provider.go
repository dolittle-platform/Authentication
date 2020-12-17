package providers

type ProviderID = string

type Provider struct {
	ID       ProviderID
	Display  string
	ImageURL string
}
