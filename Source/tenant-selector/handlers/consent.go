package handlers

import (
	"log"
	"net/http"

	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
)

type consentHandler struct {
	*Base
}

func NewConsentHandler(base *Base) http.Handler {
	return &consentHandler{base}
}

/* Get challenge string
get consent request from hydra (contains subject, email + selected tenant in context)
accept consent request (and pass it claims to the tokens)
redirect browser to result
*/
func (h *consentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	consentChallenge := r.URL.Query().Get("consent_challenge")

	consentRequest, err := h.HydraClient.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(consentChallenge))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Consent request", consentRequest, consentRequest.Payload.Subject, consentRequest.Payload.Context)

	tokenData := struct {
		Subject string `json:"sub"`
		Email   string `json:"email"`
		Tenant  string `json:"tenant"`
	}{
		Subject: consentRequest.Payload.Subject,
		Email:   consentRequest.Payload.Context.(map[string]interface{})["email"].(string),
		Tenant:  consentRequest.Payload.Context.(map[string]interface{})["tenant"].(string),
	}

	body := &models.AcceptConsentRequest{
		Remember: false,
		Session: &models.ConsentRequestSession{
			AccessToken: tokenData,
			IDToken:     tokenData,
		},
		GrantScope: []string{"openid"},
	}
	// lest just accept the whole thing as its for ourselves
	acceptConsentRequestParams := admin.NewAcceptConsentRequestParams().WithConsentChallenge(consentChallenge).WithBody(body)
	acceptResponse, err := h.HydraClient.Admin.AcceptConsentRequest(acceptConsentRequestParams)
	if err != nil {
		log.Println("error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store token in cookies
	http.Redirect(w, r, *acceptResponse.Payload.RedirectTo, http.StatusFound)
}
