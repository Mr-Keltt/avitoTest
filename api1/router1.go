package api

import (
	"avitoTest/api/handlers/organization_handler"
	"avitoTest/api/handlers/ping_handler"
	"avitoTest/services/organization_service"

	"github.com/gorilla/mux"
)

// InitRoutes - initializes API routes
func InitRoutes(router *mux.Router, orgService organization_service.OrganizationService) {
	// Инициализация сервиса организации
	orgHandler := organization_handler.NewOrganizationHandler(orgService)

	// Маршрут для проверки доступности сервера
	router.HandleFunc("/api/ping", ping_handler.PingHandler).Methods("GET")

	// Маршруты для работы с организациями
	router.HandleFunc("/api/organizations/", orgHandler.CreateOrganization).Methods("POST")
	router.HandleFunc("/api/organizations/", orgHandler.GetOrganizations).Methods("GET")
	router.HandleFunc("/api/organizations/{id}", orgHandler.GetOrganizationByID).Methods("GET")
	router.HandleFunc("/api/organizations/{id}", orgHandler.UpdateOrganization).Methods("PATCH")
	router.HandleFunc("/api/organizations/{id}", orgHandler.DeleteOrganization).Methods("DELETE")
}
