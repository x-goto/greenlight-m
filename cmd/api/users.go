package main

import "net/http"

func (app *application) userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	if err = app.writeJSON(w, http.StatusOK, input, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
