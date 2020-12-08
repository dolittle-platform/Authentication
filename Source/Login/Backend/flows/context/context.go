package context

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
	"github.com/mitchellh/mapstructure"
	"github.com/ory/hydra-client-go/models"
)

type Context struct {
	SelectedTenant tenants.TenantID
	User           *users.User
}

func StoreIn(request *models.AcceptLoginRequest, context *Context) {
	request.Context = context
}

func RetrieveFrom(response *models.ConsentRequest) (*Context, error) {
	context := &Context{}
	return context, mapstructure.Decode(response.Context, context)
}
