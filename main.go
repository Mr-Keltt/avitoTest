// main.go
package main

import (
	"avitoTest/api"
	"avitoTest/data/context"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/shared"
	"log"
	"net/http"
)

func main() {
	// Loading the configuration
	conf := shared.LoadConfig()

	// Initializing the logger
	shared.InitLogger(conf)

	// Connecting to the database
	db := context.ConnectDB()

	// We create a repository and organization service
	orgRepo := organization_repository.NewOrganizationRepository(db)
	orgService := organization_service.NewOrganizationService(orgRepo)

	// Creating a router
	r := api.NewRouter(orgService)

	// Starting the server
	log.Printf("Server started on %s", conf.ServerAddress)
	if err := http.ListenAndServe(conf.ServerAddress, r); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
