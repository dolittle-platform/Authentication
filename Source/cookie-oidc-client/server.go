package main

import (
	"fmt"
	"log"
	"net/http"

	"dolittle.io/cookie-oidc-client/configuration"
	"dolittle.io/cookie-oidc-client/utils"
)

type Server struct {
	config *configuration.Configuration
}

func NewServer(config *configuration.Configuration) *Server {
	return &Server{config}
}

func (self *Server) Start() {
	log.Println(fmt.Sprintf("Listening on port %d", self.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", self.config.Port), nil)
	if utils.TryLogIfErrorFatal(err) {
		return
	}
}
