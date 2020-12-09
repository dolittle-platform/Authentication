package viper

import (
	"dolittle.io/login/clients/hydra"
	"dolittle.io/login/clients/kratos"
	"dolittle.io/login/configuration/viper/clients"
)

type clientsConfiguration struct {
	hydra  *clients.Hydra
	kratos *clients.Kratos
}

func (c *clientsConfiguration) Hydra() hydra.Configuration {
	return c.hydra
}

func (c *clientsConfiguration) Kratos() kratos.Configuration {
	return c.kratos
}
