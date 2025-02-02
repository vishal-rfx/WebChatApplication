package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/signup", app.signup)
	mux.HandleFunc("POST /auth/login", app.login)

	// Creates a middleware chain containing our standard middleware which will be used for every request
	// our application receives.
	standard := alice.New(app.enableCors, app.logRequest)

	return standard.Then(mux)

}