package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sheetpilot/sheet-pilot-api/internal/scaleservice"
)

type SheetPilotController struct {
	router       *mux.Router
	scaleservice *scaleservice.ScaleService
}

func NewSheetPilotController(router *mux.Router, scaleservice *scaleservice.ScaleService) *SheetPilotController {
	return &SheetPilotController{
		router:       router,
		scaleservice: scaleservice,
	}
}

func (controller *SheetPilotController) SetUpRouter() {
	controller.router.
		Methods("GET").
		Path("/healthcheck").
		HandlerFunc(controller.healthCheck)

	controller.router.
		Methods("POST").
		Path("/scale/process").
		HandlerFunc(controller.scaleProcess)
}

func (controller *SheetPilotController) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("200: status ok"))

	return
}

func (app *SheetPilotController) scaleProcess(w http.ResponseWriter, r *http.Request) {
	// ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	// defer cancel()

	// response, err := app.scaleservice.SendScaleRequest(ctx)

	return
}
