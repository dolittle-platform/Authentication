package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"

	oidc "github.com/coreos/go-oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"

	hydra "github.com/ory/hydra-client-go/client"
	admin "github.com/ory/hydra-client-go/client/admin"
	models "github.com/ory/hydra-client-go/models"

	kratos "github.com/ory/kratos-client-go/client"
	public "github.com/ory/kratos-client-go/client/public"

	runtimeClient "github.com/go-openapi/runtime/client"
)

type handler struct {
	oidcConfig   oidc.Config
	oauthConfig  oauth2.Config
	sessionStore sessions.Store
	hydraClient  *hydra.OryHydra
	kratosClient *kratos.OryKratos
}

func generateNonce() (string, error) {
	buf := make([]byte, 18)
	_, err := io.ReadAtLeast(rand.Reader, buf, 18)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(buf[0:18]), nil
}

type selectedHandler struct {
	*handler
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

	whoami, err := h.kratosClient.Public.Whoami(public.NewWhoamiParams().WithCookie(&cookieVal), runtimeClient.PassThroughAuth)
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
	acceptLogin, err := h.hydraClient.Admin.AcceptLoginRequest(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, *acceptLogin.Payload.RedirectTo, http.StatusFound)
}

type consentHandler struct {
	*handler
}

/* Get challenge string
get consent request from hydra (contains subject, email + selected tenant in context)
accept consent request (and pass it claims to the tokens)
redirect browser to result
*/
func (h *consentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	consentChallenge := r.URL.Query().Get("consent_challenge")

	consentRequest, err := h.hydraClient.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(consentChallenge))
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
	acceptResponse, err := h.hydraClient.Admin.AcceptConsentRequest(acceptConsentRequestParams)
	if err != nil {
		log.Println("error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store token in cookies
	http.Redirect(w, r, *acceptResponse.Payload.RedirectTo, http.StatusFound)
}

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	}

	hydraAdminURL, _ := url.Parse("http://localhost:4445")
	hydraAdmin := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{
		Schemes:  []string{hydraAdminURL.Scheme},
		Host:     hydraAdminURL.Host,
		BasePath: hydraAdminURL.Path,
	})

	kratosPublicURL, _ := url.Parse("http://localhost:8080/.ory/kratos/public/")
	kratosPublic := kratos.NewHTTPClientWithConfig(nil, &kratos.TransportConfig{
		Schemes:  []string{kratosPublicURL.Scheme},
		Host:     kratosPublicURL.Host,
		BasePath: kratosPublicURL.Path,
	})

	h := &handler{
		oidcConfig: oidc.Config{
			ClientID: "do",
		},
		oauthConfig: oauth2.Config{
			ClientID:     "do",
			ClientSecret: "little",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://localhost:8080/.auth/callback/",
			Scopes:       []string{oidc.ScopeOpenID},
		},
		sessionStore: sessions.NewCookieStore([]byte("super-secret-value")),
		hydraClient:  hydraAdmin,
		kratosClient: kratosPublic,
	}

	http.Handle("/selected-tenant/", &selectedHandler{h})
	http.Handle("/consent/", &consentHandler{h})

	http.Handle("/", http.FileServer(http.Dir("spa")))

	log.Println("Listening on http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
