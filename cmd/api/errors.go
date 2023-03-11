package main

import (
	"goto/greenlight-m/pkg/utils"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.LogError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := utils.Envelope{
		"error": message,
	}

	err := app.writeResponse(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "server is unable to process your request due to problems"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) validationFailedResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) resourceNotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "resource couldn't be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
