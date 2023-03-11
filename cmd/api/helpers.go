package main

import "net/http"

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
