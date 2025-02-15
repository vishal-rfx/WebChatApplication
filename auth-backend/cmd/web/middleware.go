package main

import "net/http"

func (app *application) enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight request (Before making a Cross-Origin request (eg. POST with JSON payloads), browsers send an OPTIONS request)
		// to check if the server allows it.
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.URL.RequestURI()
		)

		app.logger.Info("received request", "ip", ip, "protocol", proto, "method", method, "uri", uri)
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("jwt")
		if tokenString == "" {
			app.logger.Debug("Token does not exist", "token", tokenString)
			return
		}

		ok, err := app.verifyJwtToken(tokenString)
		if err != nil {
			app.logger.Debug("Token verification failed", "token", tokenString);
			return
		}
		if !ok {
			app.logger.Debug("Invalid token", "token", tokenString)
			return
		}

		next.ServeHTTP(w, r)
	})
}
