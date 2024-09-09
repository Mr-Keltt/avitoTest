// main.go
package main

import (
	"avitoTest/api"
	"avitoTest/data/context"
	"avitoTest/shared"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Loading the configuration
	conf := shared.LoadConfig()

	// Initialize the logger
	shared.InitLogger(conf)

	// Connecting to the database
	context.ConnectDB()

	// Creating a new router
	router := mux.NewRouter()

	// Initializing routes
	api.InitRoutes(router)

	// Запуск HTTP-сервера
	serverAddress := conf.ServerAddress
	log.Printf("Starting server on %s...", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
