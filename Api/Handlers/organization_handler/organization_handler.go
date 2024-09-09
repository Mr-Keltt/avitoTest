package organization_handler

import (
	"avitoTest/api/handlers/organization_handler/handler_models"
	"avitoTest/services/organization_service"
	"avitoTest/services/organization_service/service_models"
	"avitoTest/shared"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type OrganizationHandler struct {
	service  organization_service.OrganizationService
	validate *validator.Validate
}

func NewOrganizationHandler(service organization_service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *OrganizationHandler) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var req handler_models.CreateOrganizationRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	org, err := h.service.CreateOrganization(r.Context(), service_models.OrganizationCreateModel{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	})
	if err != nil {
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	render.JSON(w, r, handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	var req handler_models.UpdateOrganizationRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	org, err := h.service.UpdateOrganization(r.Context(), service_models.OrganizationUpdateModel{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	})
	if err != nil {
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	render.JSON(w, r, handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	organizations, err := h.service.GetOrganizations(r.Context())
	if err != nil {
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	var response []handler_models.OrganizationResponse
	for _, org := range organizations {
		response = append(response, handler_models.OrganizationResponse{
			ID:          org.ID,
			Name:        org.Name,
			Description: org.Description,
			Type:        org.Type,
			CreatedAt:   org.CreatedAt,
			UpdatedAt:   org.UpdatedAt,
		})
	}

	render.JSON(w, r, response)
}

func (h *OrganizationHandler) GetOrganizationByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	org, err := h.service.GetOrganizationByID(r.Context(), id)
	if err != nil {
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	render.JSON(w, r, handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	err = h.service.DeleteOrganization(r.Context(), id)
	if err != nil {
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
