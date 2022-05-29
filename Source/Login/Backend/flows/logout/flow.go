package logout

type FlowID string

type Flow struct {
	ID      FlowID `json:"id"`
	Subject string `json:"subject"`
}
