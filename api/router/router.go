package router

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/sheetpilot/sheet-pilot-api/internal/controller"
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
	router  *mux.Router
}

func New(configs Config) *Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	return &Router{
		configs: configs,
		router:  router,
	}
}

// SetupRouter initializes the main router the API server uses.
func (r *Router) SetupRouter() http.Handler {
	sheetpilotcontroller := controller.NewSheetPilotController(r.router)

	sheetpilotcontroller.SetUpRouter()
	logrus.Debug("Registering routers")

	headers := handlers.AllowedHeaders(r.configs.CorsHeaders)
	methods := handlers.AllowedHeaders(r.configs.CorsMethods)
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r.router)
}
