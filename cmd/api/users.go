package main

import (
	"goto/greenlight-m/pkg/validator"
	"net/http"
)

func (app *application) userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	v.Check(len(input.Username) >= 8, "username", "length of the username must be greater than or equal to 8")
	v.Check(len(input.Username) <= 25, "username", "length of the username must be less than or equal to 25")

	if !v.Valid() {
		app.validationFailedResponse(w, r, v.Errors)
		return
	}

	if err = app.writeJSON(w, http.StatusOK, input, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
