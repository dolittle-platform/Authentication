package context

import (
	"dolittle.io/login/identities/tenants"
	"dolittle.io/login/identities/users"
	"github.com/mitchellh/mapstructure"
	"github.com/ory/hydra-client-go/models"
)

type Context struct {
	User           *users.User      `json:"user"`
	SelectedTenant tenants.TenantID `json:"selectedTenant"`
}

func StoreIn(request *models.AcceptLoginRequest, context *Context) {
	request.Context = context
}

func RetrieveFrom(response *models.ConsentRequest) (*Context, error) {
	context := &Context{}
	return context, mapstructure.Decode(response.Context, context)
}
