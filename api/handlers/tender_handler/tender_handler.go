package tender_handler

import (
	"avitoTest/api/handlers/tender_handler/tender_handler_models"
	"avitoTest/services/tender_service"
	"avitoTest/services/tender_service/tender_models"
	"avitoTest/services/user_service"
	"avitoTest/shared/constants"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TenderHandler struct {
	tender_service tender_service.TenderService
	user_service   user_service.UserService
}

func NewTenderHandler(tender_service tender_service.TenderService, user_service user_service.UserService) *TenderHandler {
	return &TenderHandler{
		tender_service: tender_service,
		user_service:   user_service,
	}
}

// CreateTender handles the creation of a new tender
func (h *TenderHandler) CreateTender(w http.ResponseWriter, r *http.Request) {
	var req tender_handler_models.CreateTenderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get CreatorID based on CreatorUsername
	user, err := h.user_service.GetUserByUsername(r.Context(), req.CreatorUsername)
	if err != nil {
		http.Error(w, "Invalid creator username", http.StatusBadRequest)
		return
	}

	// Create model for creating the tender
	tenderCreateModel := tender_models.TenderCreateModel{
		Name:           req.Name,
		Description:    req.Description,
		ServiceType:    req.ServiceType,
		Status:         constants.TenderStatus(req.Status),
		OrganizationID: req.OrganizationID,
		CreatorID:      user.ID,
	}

	// Create tender via service
	tender, err := h.tender_service.CreateTender(r.Context(), tenderCreateModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Form response with all relevant fields
	resp := tender_handler_models.TenderResponse{
		ID:             tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType, // Correctly set ServiceType
		Status:         string(tender.Status),
		OrganizationID: tender.OrganizationID,
		CreatedAt:      tender.CreatedAt,
		Version:        tender.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// GetTenders handles fetching all tenders with optional filtering by service type
func (h *TenderHandler) GetTenders(w http.ResponseWriter, r *http.Request) {
	// Getting the filter value from the query parameters
	serviceTypeFilter := r.URL.Query().Get("serviceType")

	// Calling a service with a filter
	tenders, err := h.tender_service.GetAllTenders(r.Context(), serviceTypeFilter)
	if err != nil {
		if err.Error() == "invalid service type" {
			http.Error(w, "Invalid service type provided", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []tender_handler_models.TenderResponse
	for _, tender := range tenders {
		resp = append(resp, tender_handler_models.TenderResponse{
			ID:             tender.ID,
			Name:           tender.Name,
			Description:    tender.Description,
			ServiceType:    tender.ServiceType,
			Status:         string(tender.Status),
			OrganizationID: tender.OrganizationID,
			CreatedAt:      tender.CreatedAt,
			Version:        tender.Version,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// GetTenderByID handles fetching a single tender by its ID
func (h *TenderHandler) GetTenderByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["tenderId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	tender, err := h.tender_service.GetTenderByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := tender_handler_models.TenderResponse{
		ID:             tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		Status:         string(tender.Status),
		OrganizationID: tender.OrganizationID,
		CreatedAt:      tender.CreatedAt,
		Version:        tender.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// UpdateTender handles updating an existing tender
func (h *TenderHandler) UpdateTender(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["tenderId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	var req tender_handler_models.UpdateTenderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tenderUpdateModel := tender_models.TenderUpdateModel{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	tender, err := h.tender_service.UpdateTender(r.Context(), tenderUpdateModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := tender_handler_models.TenderResponse{
		ID:             tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		Status:         string(tender.Status),
		OrganizationID: tender.OrganizationID,
		CreatedAt:      tender.CreatedAt,
		Version:        tender.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// PublishTender handles publishing a tender
func (h *TenderHandler) PublishTender(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["tenderId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	if err := h.tender_service.PublishTender(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CloseTender handles closing a tender
func (h *TenderHandler) CloseTender(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["tenderId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	if err := h.tender_service.CloseTender(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RollbackTenderVersion handles rolling back a tender to a previous version
func (h *TenderHandler) RollbackTenderVersion(w http.ResponseWriter, r *http.Request) {
	tenderIDStr := mux.Vars(r)["tenderId"]
	tenderID, err := strconv.Atoi(tenderIDStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	versionStr := mux.Vars(r)["version"]
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		http.Error(w, "Invalid version number", http.StatusBadRequest)
		return
	}

	tender, err := h.tender_service.RollbackTenderVersion(r.Context(), tenderID, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := tender_handler_models.TenderResponse{
		ID:             tender.ID,
		Name:           tender.Name,
		Description:    tender.Description,
		ServiceType:    tender.ServiceType,
		Status:         string(tender.Status),
		OrganizationID: tender.OrganizationID,
		CreatedAt:      tender.CreatedAt,
		Version:        tender.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
