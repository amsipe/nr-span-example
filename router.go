package main

import (
	"github.com/amsipe/nr-span-example/server"
	"github.com/go-chi/chi"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type routerConfig struct {
	Server   *server.Server
	NewRelic *newrelic.Application
}

func newRouter(c *routerConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get(newrelic.WrapHandleFunc(c.NewRelic, "/make-spans/v1", c.Server.HandleMakeSpans()))
	})

	return router
}
