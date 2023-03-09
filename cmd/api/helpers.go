package main

import (
	"encoding/json"
	"goto/greenlight-m/internal/data"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, val := range headers {
		w.Header()[key] = val
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) isRoleInSet(role data.Role, allowedRoles []data.Role) bool {
	for _, val := range allowedRoles {
		if val == role {
			return true
		}
	}

	return false
}
