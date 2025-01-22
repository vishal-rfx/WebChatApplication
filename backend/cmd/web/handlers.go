package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func (m *WebsocketManager) serveWS(w http.ResponseWriter, r  *http.Request){
	m.serveWSHelper(w, r)
}