package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/vishal-rfx/auth-backend/internal/models"
)

type SignupData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user SignupData
	// Decode the json request and store it in user variable
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.clientError(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	// They must be validated in the front-end or in the backend
	username, password := user.Username, user.Password
	ok, err := app.user.Exists(username)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if ok {
		response := Response{Message: "Username already exists"}
		w.WriteHeader(201)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			app.serverError(w, r, err)
		}
		return
	}
	err = app.user.Insert(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := Response{Message: "Signup Successful"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var user SigninData
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.clientError(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	username, password := user.Username, user.Password
	userId, err := app.user.Authenticate(username, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) || errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		app.serverError(w, r, err)
		return
	}

	jwtToken, err := app.createJwtToken(userId, SECRET)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name: "jwt",
		Value: jwtToken,
		Expires: time.Now().Add(time.Hour*4),
		SameSite: http.SameSiteStrictMode,
		Secure: false,
	})

	w.WriteHeader(http.StatusOK)
	response := Response{ Message: "Sign In Successful"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}