package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.LogError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{
		"error": message,
	}

	if err := app.writeJSON(w, status, env, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "server is unable to process your request due to problems"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) resourceNotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "resource couldn't be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
