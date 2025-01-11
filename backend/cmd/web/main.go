package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvironmentVariables(){
	err := godotenv.Load("./cmd/web/.env")
	if err != nil {
		log.Fatal("Error in loading .env file", err.Error())
	}
}


func main(){
	loadEnvironmentVariables()

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("Port is not present in the Environment Variables")
		port = ":4000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ws", serveWS)

	log.Printf("Starting a server on %s", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}