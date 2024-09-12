package api

import (
	"avitoTest/api/handlers/organization_handler"
	"avitoTest/api/handlers/ping_handler"
	"avitoTest/api/handlers/tender_handler"
	"avitoTest/api/handlers/user_handler"
	"avitoTest/services/organization_service"
	"avitoTest/services/tender_service"
	"avitoTest/services/user_service"

	"github.com/gorilla/mux"
)

// InitRoutes initializes all API routes.
func InitRoutes(router *mux.Router,
	orgService organization_service.OrganizationService,
	userService user_service.UserService,
	tenderService tender_service.TenderService) {
	// Initialize individual route groups
	initPingRoutes(router)
	initOrganizationRoutes(router, orgService)
	initUserRoutes(router, userService)
	initTenderRoutes(router, tenderService, userService)
}

// initPingRoutes sets up routes for server availability checks.
func initPingRoutes(router *mux.Router) {
	router.HandleFunc("/api/ping", ping_handler.PingHandler).Methods("GET")
}

// initOrganizationRoutes sets up routes for organization-related operations.
func initOrganizationRoutes(router *mux.Router, orgService organization_service.OrganizationService) {
	orgHandler := organization_handler.NewOrganizationHandler(orgService)

	router.HandleFunc("/api/organizations/new", orgHandler.CreateOrganization).Methods("POST")
	router.HandleFunc("/api/organizations/", orgHandler.GetOrganizations).Methods("GET")
	router.HandleFunc("/api/organizations/{org_id}", orgHandler.GetOrganizationByID).Methods("GET")
	router.HandleFunc("/api/organizations/{org_id}/edit", orgHandler.UpdateOrganization).Methods("PATCH")
	router.HandleFunc("/api/organizations/{org_id}/delete", orgHandler.DeleteOrganization).Methods("DELETE")

	// Routes for managing responsibilities
	router.HandleFunc("/api/organizations/{org_id}/responsibles", orgHandler.GetResponsibles).Methods("GET")
	router.HandleFunc("/api/organizations/{org_id}/responsibles/{user_id}", orgHandler.GetResponsibleByID).Methods("GET")
	router.HandleFunc("/api/organizations/{org_id}/responsibles/{user_id}/new", orgHandler.AddResponsible).Methods("POST")
	router.HandleFunc("/api/organizations/{org_id}/responsibles/{user_id}/delete", orgHandler.DeleteResponsible).Methods("DELETE")
}

// initUserRoutes sets up routes for user-related operations.
func initUserRoutes(router *mux.Router, userService user_service.UserService) {
	userHandler := user_handler.NewUserHandler(userService)

	router.HandleFunc("/api/users/new", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{user_id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{user_id}/edit", userHandler.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/users/{user_id}/delete", userHandler.DeleteUser).Methods("DELETE")
}

// initTenderRoutes sets up routes for tender-related operations.
func initTenderRoutes(router *mux.Router, tenderService tender_service.TenderService, userService user_service.UserService) {
	tenderHandler := tender_handler.NewTenderHandler(tenderService, userService)

	router.HandleFunc("/api/tenders/new", tenderHandler.CreateTender).Methods("POST")
	router.HandleFunc("/api/tenders/", tenderHandler.GetTenders).Methods("GET")
	router.HandleFunc("/api/tenders/{tenderId}", tenderHandler.GetTenderByID).Methods("GET")
	router.HandleFunc("/api/tenders/my/{username}", tenderHandler.GetTendersByUsername).Methods("GET")
	router.HandleFunc("/api/tenders/{tenderId}/edit", tenderHandler.UpdateTender).Methods("PATCH")
	router.HandleFunc("/api/tenders/{tenderId}/publish", tenderHandler.PublishTender).Methods("POST")
	router.HandleFunc("/api/tenders/{tenderId}/close", tenderHandler.CloseTender).Methods("POST")
	router.HandleFunc("/api/tenders/{tenderId}/rollback/{version}", tenderHandler.RollbackTenderVersion).Methods("PUT")
	router.HandleFunc("/api/tenders/{tenderId}/delete", tenderHandler.DeleteTender).Methods("DELETE")
}
