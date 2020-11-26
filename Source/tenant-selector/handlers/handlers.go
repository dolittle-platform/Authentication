package handlers

import (
	"dolittle.io/tenant-selector/configuration"

	hydra "github.com/ory/hydra-client-go/client"
	kratos "github.com/ory/kratos-client-go/client"
)

type Base struct {
	Configuration configuration.Configuration
	HydraClient   *hydra.OryHydra
	KratosClient  *kratos.OryKratos
}
