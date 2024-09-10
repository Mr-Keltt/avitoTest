package api

import (
	"avitoTest/api/handlers/organization_handler"
	"avitoTest/api/handlers/ping_handler"
	"avitoTest/services/organization_service"

	"github.com/gorilla/mux"
)

// InitRoutes - initializes API routes
func InitRoutes(router *mux.Router,
	orgService organization_service.OrganizationService) {
	// Initializing an organization service
	orgHandler := organization_handler.NewOrganizationHandler(orgService)

	// Route to check server availability
	router.HandleFunc("/api/ping", ping_handler.PingHandler).Methods("GET")

	// Routes for working with organizations
	router.HandleFunc("/api/organizations/", orgHandler.CreateOrganization).Methods("POST")
	router.HandleFunc("/api/organizations/", orgHandler.GetOrganizations).Methods("GET")
	router.HandleFunc("/api/organizations/{id}", orgHandler.GetOrganizationByID).Methods("GET")
	router.HandleFunc("/api/organizations/{id}", orgHandler.UpdateOrganization).Methods("PATCH")
	router.HandleFunc("/api/organizations/{id}", orgHandler.DeleteOrganization).Methods("DELETE")
}
