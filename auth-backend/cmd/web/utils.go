package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Create a JWT Token with userId and secretKey. It uses HS256 signing method.
func (app *application) createJwtToken(userId, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid" : userId,
			"exp" : time.Now().Add(time.Hour * 4).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// serverError helper writes a log entry at Error level (including the request method and request URI as attributes),
// then sends a generic 500 Internal server error response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error){
	var (
		method = r.Method
		uri = r.RequestURI
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) clientError(w http.ResponseWriter, errMsg string, code int){
	http.Error(w, errMsg, code)
}