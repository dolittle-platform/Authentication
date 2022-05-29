package providers

type ProviderID = string

type Provider struct {
	ID       ProviderID `json:"id"`
	Display  string     `json:"display"`
	ImageURL string     `json:"imageURL"`
}
