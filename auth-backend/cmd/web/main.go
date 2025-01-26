package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vishal-rfx/auth-backend/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	logger *slog.Logger
	user *models.UserModel
	PORT string
	MONGO_URI string
	SECRET string
}

var PORT string
var MONGO_URI string
var SECRET string

func loadEnvironmentVariables(){
	err := godotenv.Load("./cmd/web/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	PORT = os.Getenv("PORT") 
	if PORT == "" {
		PORT = ":8001"
	}

	MONGO_URI = os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("Mongo URI not found in environment variables")
	}

	SECRET = os.Getenv("SECRET")
	if SECRET == "" {
		log.Fatal("Secret is not defined")
	}

}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func main(){
	loadEnvironmentVariables()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	client, err := connectToMongoDB(MONGO_URI)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)	
	}
	logger.Debug("Connected to MongoDB")

	// Disconnect all the connections in the client 
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := &application{
		logger: logger,
		user: &models.UserModel{Client: client},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/signup", app.signup)
	mux.HandleFunc("POST /auth/signin", app.signin)

	app.logger.Info("Starting a server on %s", "addr", PORT)
	err = http.ListenAndServe(PORT, mux)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}