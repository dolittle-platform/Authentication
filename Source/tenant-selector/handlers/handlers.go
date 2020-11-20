package handlers

import (
	"log"
	"net/http"

	"dolittle.io/tenant-selector/configuration"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
	"github.com/ory/kratos-client-go/client/public"

	hydra "github.com/ory/hydra-client-go/client"
	kratos "github.com/ory/kratos-client-go/client"

	runtimeClient "github.com/go-openapi/runtime/client"
)

type Base struct {
	Configuration configuration.Configuration
	HydraClient   *hydra.OryHydra
	KratosClient  *kratos.OryKratos
}

type SelectedHandler struct {
	*Base
}

func (h *SelectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookieVal := cookie.String()

	whoami, err := h.KratosClient.Public.Whoami(public.NewWhoamiParams().WithCookie(&cookieVal), runtimeClient.PassThroughAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	identity := whoami.Payload.Identity
	subject := string(identity.ID)
	traits := identity.Traits.(map[string]interface{})
	log.Println("IDENTITY", identity, subject, traits)

	// Verify that selected tenant is actually in list of tenants
	body := &models.AcceptLoginRequest{
		Subject: &subject,
		Context: struct {
			Email  string `json:"email"`
			Tenant string `json:"tenant"`
		}{
			Email:  traits["email"].(string),
			Tenant: r.PostFormValue("tenant"),
		},
		Remember: false,
	}
	params := admin.NewAcceptLoginRequestParams().WithLoginChallenge(r.PostFormValue("login_challenge")).WithBody(body)
	acceptLogin, err := h.HydraClient.Admin.AcceptLoginRequest(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, *acceptLogin.Payload.RedirectTo, http.StatusFound)
}

type ConsentHandler struct {
	*Base
}

/* Get challenge string
get consent request from hydra (contains subject, email + selected tenant in context)
accept consent request (and pass it claims to the tokens)
redirect browser to result
*/
func (h *ConsentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
