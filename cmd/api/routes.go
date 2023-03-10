package main

import (
	"goto/greenlight-m/internal/data/users"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	moderatorAndHigher = []users.Role{users.RoleModerator, users.RoleAdmin}
	onlyAdmin          = []users.Role{users.RoleAdmin}
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//users
	router.HandlerFunc(http.MethodPost, "/v1/registration", app.userRegistrationHandler)

	return app.recoverPanic(router)
}
