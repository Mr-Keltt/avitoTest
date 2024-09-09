// main.go
package main

import (
	"avitoTest/api"
	"avitoTest/shared"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Loading the configuration
	conf := shared.LoadConfig()

	// Initializing the logger
	shared.InitLogger(conf)

	// Connecting to the database
	//db := context.ConnectDB()

	// We create a repository and organization service
	//orgRepo := organization_repository.NewOrganizationRepository(db)
	//orgService := organization_service.NewOrganizationService(orgRepo)

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
