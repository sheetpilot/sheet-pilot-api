package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sheetpilot/sheet-pilot-api/internal/model"
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
	scale, err := model.TransformScaleObject(r.Body)
	if err != nil {
		return // TODO: return http error
	}

	if err := scale.Validate(); err != nil {
		return // TODO: return http error
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute*3)
	defer cancel()

	_, err = app.scaleservice.SendScaleRequest(ctx, scale.UpdatedRow)
	if err != nil {
		return // TODO: return http error
	}

	// TODO: return http response using response.Data
	return
}
