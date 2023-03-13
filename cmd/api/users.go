package main

import (
	"context"
	userdto "goto/greenlight-m/internal/data/user/userdtos"
	"net/http"
	"time"
)

func (app *application) userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var input userdto.CreateUserDTO

	err := app.readRequest(r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = app.repositories.Users.Create(ctx, &input); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = app.writeResponse(w, http.StatusOK, input, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) testUserUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var input userdto.UpdateUserDTO
	input.ID = id

	app.readRequest(r, &input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = app.repositories.Users.UpdateUser(ctx, &input); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = app.writeResponse(w, http.StatusOK, input, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) testDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = app.repositories.Users.DeleteByID(ctx, id); err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = app.writeResponse(w, http.StatusOK, "user was deleted", nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) testGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := app.repositories.Users.GetByID(ctx, id)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	if err = app.writeResponse(w, http.StatusOK, user, nil); err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
