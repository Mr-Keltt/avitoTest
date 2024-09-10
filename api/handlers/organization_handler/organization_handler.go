package organization_handler

import (
	"avitoTest/api/handlers/organization_handler/organization_handler_models"
	"avitoTest/api/handlers/user_handler/user_handler_models"
	"avitoTest/services/organization_service"
	"avitoTest/services/organization_service/organization_models"
	"avitoTest/shared"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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
	shared.Logger.Infof("CreateOrganization: Handling request from %s", r.RemoteAddr)
	var req organization_handler_models.CreateOrganizationRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		shared.Logger.Errorf("CreateOrganization: Failed to decode request: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	if err := h.validate.Struct(req); err != nil {
		shared.Logger.Errorf("CreateOrganization: Validation failed: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	shared.Logger.Infof("CreateOrganization: Request - Name: %s, Description: %s, Type: %s", req.Name, req.Description, req.Type)

	org, err := h.service.CreateOrganization(r.Context(), organization_models.OrganizationCreateModel{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	})
	if err != nil {
		shared.Logger.Errorf("CreateOrganization: Failed to create organization: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	shared.Logger.Infof("CreateOrganization: Organization created successfully: ID=%d", org.ID)
	render.JSON(w, r, organization_handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	shared.Logger.Infof("UpdateOrganization: Handling update for organization ID: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		shared.Logger.Errorf("UpdateOrganization: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	var req organization_handler_models.UpdateOrganizationRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		shared.Logger.Errorf("UpdateOrganization: Failed to decode request body: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	shared.Logger.Infof("UpdateOrganization: Request - Name: %s, Description: %s, Type: %s", req.Name, req.Description, req.Type)

	if err := h.validate.Struct(req); err != nil {
		shared.Logger.Errorf("UpdateOrganization: Validation failed: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	org, err := h.service.UpdateOrganization(r.Context(), organization_models.OrganizationUpdateModel{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
	})
	if err != nil {
		shared.Logger.Errorf("UpdateOrganization: Failed to update organization: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	shared.Logger.Infof("UpdateOrganization: Organization updated successfully: ID=%d", org.ID)
	render.JSON(w, r, organization_handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	shared.Logger.Infof("GetOrganizations: Handling request to fetch all organizations")

	organizations, err := h.service.GetOrganizations(r.Context())
	if err != nil {
		shared.Logger.Errorf("GetOrganizations: Failed to fetch organizations: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	var response []organization_handler_models.OrganizationResponse
	for _, org := range organizations {
		shared.Logger.Infof("GetOrganizations: Found organization - ID=%d, Name=%s", org.ID, org.Name)
		response = append(response, organization_handler_models.OrganizationResponse{
			ID:          org.ID,
			Name:        org.Name,
			Description: org.Description,
			Type:        org.Type,
			CreatedAt:   org.CreatedAt,
			UpdatedAt:   org.UpdatedAt,
		})
	}

	shared.Logger.Infof("GetOrganizations: Successfully fetched %d organizations", len(response))
	render.JSON(w, r, response)
}

func (h *OrganizationHandler) GetOrganizationByID(w http.ResponseWriter, r *http.Request) {
	shared.Logger.Infof("GetOrganizationByID: Received request from %s", r.RemoteAddr)
	vars := mux.Vars(r)
	idParam := vars["id"]

	// Log the incoming ID
	shared.Logger.Infof("GetOrganizationByID: Fetching organization with ID: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		shared.Logger.Errorf("GetOrganizationByID: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	// Log before calling the service
	shared.Logger.Infof("GetOrganizationByID: Calling service to get organization with ID: %d", id)
	org, err := h.service.GetOrganizationByID(r.Context(), id)
	if err != nil {
		shared.Logger.Errorf("GetOrganizationByID: Failed to get organization by ID: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	// Log successful retrieval
	shared.Logger.Infof("GetOrganizationByID: Successfully retrieved organization with ID: %d", org.ID)

	// Respond with the organization details
	render.JSON(w, r, organization_handler_models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		Type:        org.Type,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	})
}

func (h *OrganizationHandler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	shared.Logger.Infof("DeleteOrganization: Received request from %s", r.RemoteAddr)
	vars := mux.Vars(r)
	idParam := vars["id"]

	// Log the incoming ID
	shared.Logger.Infof("DeleteOrganization: Deleting organization with ID: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		shared.Logger.Errorf("DeleteOrganization: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	// Log before calling the service
	shared.Logger.Infof("DeleteOrganization: Calling service to delete organization with ID: %d", id)
	err = h.service.DeleteOrganization(r.Context(), id)
	if err != nil {
		shared.Logger.Errorf("DeleteOrganization: Failed to delete organization: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	// Log successful deletion
	shared.Logger.Infof("DeleteOrganization: Successfully deleted organization with ID: %d", id)

	// Return no content response
	w.WriteHeader(http.StatusNoContent)
}

// AddResponsible - Adds a user as a responsible person for an organization.
func (h *OrganizationHandler) AddResponsible(w http.ResponseWriter, r *http.Request) {
	orgIDParam := mux.Vars(r)["org_id"]
	userIDParam := mux.Vars(r)["user_id"]

	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		shared.Logger.Errorf("AddResponsible: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		shared.Logger.Errorf("AddResponsible: Invalid user ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	if err := h.service.AddResponsible(r.Context(), orgID, userID); err != nil {
		shared.Logger.Errorf("AddResponsible: Failed to add responsible user: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	shared.Logger.Infof("AddResponsible: Successfully added user ID=%d as responsible for organization ID=%d", userID, orgID)
	w.WriteHeader(http.StatusNoContent)
}

// DeleteResponsible - Removes a user as a responsible person for an organization.
func (h *OrganizationHandler) DeleteResponsible(w http.ResponseWriter, r *http.Request) {
	orgIDParam := mux.Vars(r)["org_id"]
	userIDParam := mux.Vars(r)["user_id"]

	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		shared.Logger.Errorf("DeleteResponsible: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		shared.Logger.Errorf("DeleteResponsible: Invalid user ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	if err := h.service.DeleteResponsible(r.Context(), orgID, userID); err != nil {
		shared.Logger.Errorf("DeleteResponsible: Failed to delete responsible user: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	shared.Logger.Infof("DeleteResponsible: Successfully removed user ID=%d from being responsible for organization ID=%d", userID, orgID)
	w.WriteHeader(http.StatusNoContent)
}

// GetResponsibles - Retrieves all users responsible for a specific organization.
func (h *OrganizationHandler) GetResponsibles(w http.ResponseWriter, r *http.Request) {
	orgIDParam := mux.Vars(r)["org_id"]

	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		shared.Logger.Errorf("GetResponsibles: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	users, err := h.service.GetResponsibles(r.Context(), orgID)
	if err != nil {
		shared.Logger.Errorf("GetResponsibles: Failed to get responsible users: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	var response []user_handler_models.UserResponse
	for _, user := range users {
		response = append(response, user_handler_models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	shared.Logger.Infof("GetResponsibles: Successfully retrieved %d responsible users for organization ID=%d", len(response), orgID)
	render.JSON(w, r, response)
}

// GetResponsibleByID - Retrieves a specific responsible user by organization ID and user ID.
func (h *OrganizationHandler) GetResponsibleByID(w http.ResponseWriter, r *http.Request) {
	orgIDParam := mux.Vars(r)["org_id"]
	userIDParam := mux.Vars(r)["user_id"]

	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		shared.Logger.Errorf("GetResponsibleByID: Invalid organization ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		shared.Logger.Errorf("GetResponsibleByID: Invalid user ID: %v", err)
		render.Render(w, r, shared.ErrInvalidRequest(err))
		return
	}

	user, err := h.service.GetResponsibleByID(r.Context(), orgID, userID)
	if err != nil {
		shared.Logger.Errorf("GetResponsibleByID: Failed to get responsible user: %v", err)
		render.Render(w, r, shared.ErrInternal(err))
		return
	}

	shared.Logger.Infof("GetResponsibleByID: Successfully retrieved user ID=%d as responsible for organization ID=%d", userID, orgID)
	render.JSON(w, r, user_handler_models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
}
