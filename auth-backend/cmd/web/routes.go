package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// Creates a middleware chain containing our standard middleware which will be used for every request
	// our application receives.
	standard := alice.New(app.enableCors, app.logRequest)
	mux.Handle("POST /auth/signup", standard.ThenFunc(app.signup))
	mux.Handle("POST /auth/login", standard.ThenFunc(app.login))
	
	// Creates a middleware chain containing our standard middleware and the authenticate middleware
	protected := standard.Append(app.authenticate)
	// All routes which needs authentication will be handled by this middleware chain
	mux.Handle("GET /users", protected.ThenFunc(app.getUsers))

	return standard.Then(mux)

}