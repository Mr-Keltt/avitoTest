package main

import (
	"avitoTest/api"
	"avitoTest/data/context"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/shared"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Loading configuration from environment variables
	conf := shared.LoadConfig()

	// Initializing the logger
	shared.InitLogger(conf)

	// Connecting to the database
	db, err := context.ConnectDB(conf.PostgresConn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	sqlDB, err := db.DB() // We get *sql. DB to close connection
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// We create a repository and organization service
	orgRepo := organization_repository.NewOrganizationRepository(db)
	orgService := organization_service.NewOrganizationService(orgRepo)

	// Creating a router
	router := mux.NewRouter()

	// We initialize routes, passing the necessary services
	api.InitRoutes(router, orgService)

	// Starting the HTTP server
	serverAddress := conf.ServerAddress
	log.Printf("Starting server on %s...", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
