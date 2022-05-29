package forms

type Form struct {
	SubmitMethod string `json:"submitMethod"`
	SubmitAction string `json:"submitAction"`
	CSRFToken    string `json:"csrfToken"`
}
