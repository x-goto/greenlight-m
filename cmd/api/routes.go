package main

import (
	"goto/greenlight-m/internal/data/user"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	moderatorAndHigher = []user.Role{user.RoleModerator, user.RoleAdmin}
	onlyAdmin          = []user.Role{user.RoleAdmin}
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	//users
	router.HandlerFunc(http.MethodPost, "/v1/registration", app.userRegistrationHandler)
	router.HandlerFunc(http.MethodPut, "/v1/user/update/:id", app.testUserUpdate)
	router.HandlerFunc(http.MethodDelete, "/v1/user/delete/:id", app.testDeleteUser)
	router.HandlerFunc(http.MethodGet, "/v1/user/get/:id", app.testGetUser)

	return app.recoverPanic(router)
}
