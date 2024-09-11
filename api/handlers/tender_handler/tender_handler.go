package tender_handler

import (
	"avitoTest/api/handlers/tender_handler/tender_handler_models"
	"avitoTest/services/tender_service"
	"avitoTest/services/tender_service/tender_models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TenderHandler struct {
	service tender_service.TenderService
}

func NewTenderHandler(service tender_service.TenderService) *TenderHandler {
	return &TenderHandler{service: service}
}

// CreateTender handles the creation of a new tender
func (h *TenderHandler) CreateTender(w http.ResponseWriter, r *http.Request) {
	var req tender_handler_models.CreateTenderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tenderCreateModel := tender_models.TenderCreateModel{
		Name:           req.Name,
		Description:    req.Description,
		ServiceType:    req.ServiceType,
		OrganizationID: req.OrganizationID,
		CreatorID:      req.CreatorID, // Assume CreatorID is passed in the request body
	}

	tender, err := h.service.CreateTender(r.Context(), tenderCreateModel)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetTenders handles fetching all tenders with optional filtering
func (h *TenderHandler) GetTenders(w http.ResponseWriter, r *http.Request) {
	tenders, err := h.service.GetAllTenders(r.Context())
	if err != nil {
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

	tender, err := h.service.GetTenderByID(r.Context(), id)
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
		ServiceType: req.ServiceType,
	}

	tender, err := h.service.UpdateTender(r.Context(), tenderUpdateModel)
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

	if err := h.service.PublishTender(r.Context(), id); err != nil {
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

	if err := h.service.CloseTender(r.Context(), id); err != nil {
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

	tender, err := h.service.RollbackTenderVersion(r.Context(), tenderID, version)
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
