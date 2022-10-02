package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type SheetPilotController struct {
	router *mux.Router
}

func NewSheetPilotController(router *mux.Router) *SheetPilotController {
	return &SheetPilotController{
		router: router,
	}
}

func (controller *SheetPilotController) SetUpRouter() {
	controller.router.
		Methods("GET").
		Path("/healthcheck").
		HandlerFunc(controller.healthCheck)
}

func (controller *SheetPilotController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("200: status ok")) 

	return
}
