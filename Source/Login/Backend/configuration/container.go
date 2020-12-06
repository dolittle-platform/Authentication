package configuration

import (
	"dolittle.io/login/server"
	"go.uber.org/zap"
)

type Container struct {
	Server server.Server
}

func NewContainer(configuration Configuration) (*Container, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	container := &Container{}

	container.Server = server.NewServer(configuration.Server(), logger)

	return container, nil
}
