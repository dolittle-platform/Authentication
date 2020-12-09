package clients

import (
	"dolittle.io/login/clients/hydra"
	"dolittle.io/login/clients/kratos"
)

type Configuration interface {
	Hydra() hydra.Configuration
	Kratos() kratos.Configuration
}
