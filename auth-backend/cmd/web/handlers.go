package main

import (
	"encoding/json"
	"net/http"

)

type SignupData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	StatusCode int `json:"StatusCode"`
	Message string `json:"message"`
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	var user SignupData
	// Decode the json request and store it in user variable
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// They must be validated in the front-end or in the backend
	username, password := user.Username, user.Password

	ok, err := app.user.Exists(username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ok {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	err = app.user.Insert(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{StatusCode: http.StatusOK, Message: "Signup Successful"}
	jsonStr, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonStr)

}