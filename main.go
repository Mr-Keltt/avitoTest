package main

import (
	"avitoTest/api"
	"avitoTest/data/context"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/shared"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Loading configuration from environment variables
	conf := shared.LoadConfig()
	shared.Logger.Infof("Configuration loaded: %+v", conf)

	// Initializing the logger
	shared.InitLogger(conf)
	shared.Logger.Infof("Logger initialized")

	// Connecting to the database
	shared.Logger.Info("Connecting to the database")
	db, err := context.ConnectDB(conf.PostgresConn)
	if err != nil {
		shared.Logger.Fatalf("Failed to connect to the database: %v", err)
	}
	shared.Logger.Info("Database connected successfully")

	sqlDB, err := db.DB() // We get *sql. DB to close connection
	if err != nil {
		shared.Logger.Fatalf("Failed to get sql.DB: %v", err)
	}
	defer func() {
		shared.Logger.Info("Closing database connection")
		if err := sqlDB.Close(); err != nil {
			shared.Logger.Errorf("Error while closing the database: %v", err)
		}
	}()

	// We create a repository and organization service
	shared.Logger.Info("Initializing repositories and services")
	orgRepo := organization_repository.NewOrganizationRepository(db)
	orgService := organization_service.NewOrganizationService(orgRepo)

	// Creating a router
	router := mux.NewRouter()

	// We initialize routes, passing the necessary services
	shared.Logger.Info("Initializing routes")
	api.InitRoutes(router, orgService)

	// Starting the HTTP server
	serverAddress := conf.ServerAddress
	shared.Logger.Infof("Starting server on %s...", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		shared.Logger.Fatalf("Failed to start server: %v", err)
	}
}
