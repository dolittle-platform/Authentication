package handlers

import (
	"log"
	"net/http"

	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
	"github.com/ory/kratos-client-go/client/public"

	runtimeClient "github.com/go-openapi/runtime/client"
)

type selectedHandler struct {
	*Base
}

func NewSelectedHandler(base *Base) http.Handler {
	return &selectedHandler{base}
}

func (h *selectedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
