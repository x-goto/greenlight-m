package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"system": "available",
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
