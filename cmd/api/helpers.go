package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int, error) {
	id, err := strconv.Atoi(httprouter.ParamsFromContext(r.Context()).ByName("id"))
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}
	return id, nil
}

func (app *application) readRequest(r *http.Request, dst any) error {
	return app.codec.Decode(r.Body, dst)
}

func (app *application) writeResponse(w http.ResponseWriter, status int, data any, headers http.Header) error {
	responseData, err := app.codec.Encode(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", app.codec.ContentType())
	w.WriteHeader(status)

	w.Write(responseData)
	return nil
}
