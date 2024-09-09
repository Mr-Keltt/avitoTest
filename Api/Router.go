package api

import (
	"avitoTest/api/handlers/organization_handler"
	"avitoTest/services/organization_service"

	"github.com/go-chi/chi/v5"
)

func NewRouter(orgService organization_service.OrganizationService) *chi.Mux {
	r := chi.NewRouter()

	orgHandler := organization_handler.NewOrganizationHandler(orgService)

	r.Route("/api/organizations", func(r chi.Router) {
		r.Post("/", orgHandler.CreateOrganization)
		r.Get("/", orgHandler.GetOrganizations)
		r.Get("/{id}", orgHandler.GetOrganizationByID)
		r.Put("/{id}", orgHandler.UpdateOrganization)
		r.Delete("/{id}", orgHandler.DeleteOrganization)
	})

	return r
}
