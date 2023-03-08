package main

import (
	data "goto/greenlight-m/internal/data/user"
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

	return router
}
