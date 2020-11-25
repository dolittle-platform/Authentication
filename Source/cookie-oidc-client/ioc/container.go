package ioc

import (
	"log"

	"dolittle.io/cookie-oidc-client/utils"
)

type Container interface {
	Logger() log.Logger
	utils.Container
}

var _ Container = &container{}

func NewServiceProvider() Container {
	return &container{
		utils.NewContainer(),
		log.Logger{},
	}
}

type container struct {
	utils.Container
	logger log.Logger
}

func (c *container) Logger() log.Logger {
	return c.logger
}
