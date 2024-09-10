// File: main.go

package main

import (
	"avitoTest/api"
	"avitoTest/data/context"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/shared"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {
	conf := loadConfiguration()
	initLogger(conf)
	db := connectToDatabase(conf)
	defer closeDatabaseConnection(db)

	orgService := initializeServices(db)
	router := setupRouter(orgService)

	startServer(conf.ServerAddress, router)
}

// loadConfiguration loads the application configuration from environment variables.
func loadConfiguration() *shared.Config {
	conf := shared.LoadConfig()
	shared.Logger.Infof("Configuration loaded: %+v", conf)
	return conf
}

// initLogger initializes the logger with the loaded configuration.
func initLogger(conf *shared.Config) {
	shared.InitLogger(conf)
	shared.Logger.Infof("Logger initialized")
}

// connectToDatabase connects to the database using the configuration and returns the DB connection.
func connectToDatabase(conf *shared.Config) *gorm.DB {
	shared.Logger.Info("Connecting to the database")
	db, err := context.ConnectDB(conf.PostgresConn)
	if err != nil {
		shared.Logger.Fatalf("Failed to connect to the database: %v", err)
	}
	shared.Logger.Info("Database connected successfully")
	return db
}

// closeDatabaseConnection gracefully closes the database connection.
func closeDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB() // Get *sql.DB to close connection
	if err != nil {
		shared.Logger.Fatalf("Failed to get sql.DB: %v", err)
	}
	shared.Logger.Info("Closing database connection")
	if err := sqlDB.Close(); err != nil {
		shared.Logger.Errorf("Error while closing the database: %v", err)
	}
}

// initializeServices initializes the necessary repositories and services.
func initializeServices(db *gorm.DB) organization_service.OrganizationService {
	shared.Logger.Info("Initializing repositories and services")
	orgRepo := organization_repository.NewOrganizationRepository(db)
	return organization_service.NewOrganizationService(orgRepo)
}

// setupRouter sets up the HTTP router with the necessary routes.
func setupRouter(orgService organization_service.OrganizationService) *mux.Router {
	shared.Logger.Info("Initializing routes")
	router := mux.NewRouter()
	api.InitRoutes(router, orgService)
	return router
}

// startServer starts the HTTP server with the specified address and router.
func startServer(address string, router *mux.Router) {
	shared.Logger.Infof("Starting server on %s...", address)
	if err := http.ListenAndServe(address, router); err != nil {
		shared.Logger.Fatalf("Failed to start server: %v", err)
	}
}