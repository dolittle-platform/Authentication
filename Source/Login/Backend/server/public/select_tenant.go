package public

// import (
// 	"context"
// 	"errors"
// 	"net/http"

// 	"dolittle.io/login/logins"
// 	"dolittle.io/login/server/handling"
// 	"dolittle.io/login/tenants"
// 	users "dolittle.io/login/users/current"
// )

// type SelectTenantHandler handling.Handler

// func NewSelectTenantHandler() SelectTenantHandler {
// 	return &selectTenant{}
// }

// type selectTenant struct {
// 	userReader     users.Reader
// 	tenantReader   tenants.Reader
// 	loginCompleter logins.Completer
// 	loginGetter    logins.Getter
// }

// // POST http://localhost:8080/select-tenant

// func (t *selectTenant) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
// 	// Get the login request
// 	loginRequest, err := t.loginGetter.GetRequest(r, ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// Get the signed in user
// 	user, err := t.userReader.GetCurrentUser(r, ctx)
// 	if err != nil {
// 		return err
// 	}

// 	// Get the selected tenant (from POST body)
// 	tenant, err := t.tenantReader.GetSelectedTenant(r, ctx)
// 	if err != nil {
// 		return err
// 	}

// 	if !user.HasAccessToTenant(tenant) {
// 		return errors.New("user doesn't have access to tenant")
// 	}

// 	// Complete the (hydra) login request with the user information + the selected tenant
// 	// which returns a redirect URL, respond with that
// 	redirect, err := t.loginCompleter.Complete(loginRequest, user, tenant)
// 	if err != nil {
// 		return err
// 	}

// 	http.Redirect(w, r, redirect.String(), http.StatusFound)
// 	return nil
// }
