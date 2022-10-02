package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/sheetpilot/sheet-pilot-api/internal/controller"
	"github.com/sheetpilot/sheet-pilot-api/internal/scaleservice"
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
func (r *Router) SetupRouter(scaleservice *scaleservice.ScaleService) http.Handler {
	sheetpilotcontroller := controller.NewSheetPilotController(r.router, scaleservice)

	sheetpilotcontroller.SetUpRouter()
	logrus.Debug("Registering routers")

	headers := handlers.AllowedHeaders(r.configs.CorsHeaders)
	methods := handlers.AllowedHeaders(r.configs.CorsMethods)
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(headers, methods, origins)(r.router)
}

func (r *Router) SetupSvcConn(ctx context.Context, svcAddr string) (*grpc.ClientConn, error) {
	svc, err := grpc.DialContext(ctx, svcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("grpc.DialContext(): %w", err)
	}

	return svc, nil
}
