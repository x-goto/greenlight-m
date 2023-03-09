package main

import (
	"goto/greenlight-m/internal/data"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	moderatorAndHigher = []data.Role{data.RoleModerator, data.RoleAdmin}
	onlyAdmin          = []data.Role{data.RoleAdmin}
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//users
	router.HandlerFunc(http.MethodPost, "/v1/registration", app.userRegistrationHandler)

	return router
}
