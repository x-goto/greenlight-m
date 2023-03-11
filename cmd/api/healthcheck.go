package main

import (
	"goto/greenlight-m/pkg/utils"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := utils.Envelope{
		"system": "available",
	}

	err := app.writeResponse(w, http.StatusOK, env, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
