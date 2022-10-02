package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Config provides the API server configuration
type Config struct {
	CorsHeaders []string
	CorsMethods []string

	// Hosts is a list of addresses for the API.
	Hosts []string
}

type Router struct {
	configs Config
}

func New(configs Config) *Router {
	return &Router{
		configs: configs,
	}
}

// SetupRouter initializes the main router the API server uses.
func (r *Router) SetupRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	logrus.Debug("Registering routers")

	return router
}
