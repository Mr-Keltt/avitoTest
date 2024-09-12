package handlers

package bid_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"your_project/services/bid_service"
	"your_project/api/handlers/bid_handler/models"
)

type BidHandler struct {
	service bid_service.BidService
}

func NewBidHandler(service bid_service.BidService) *BidHandler {
	return &BidHandler{service: service}
}

// CreateBid creates a new bid
func (h *BidHandler) CreateBid(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bidCreateModel := bid_service.BidCreateModel{
		Name:           req.Name,
		Description:    req.Description,
		TenderID:       req.TenderID,
		OrganizationID: req.OrganizationID,
		CreatorID:      req.CreatorID,
		Status:         req.Status,
	}

	bid, err := h.service.CreateBid(r.Context(), bidCreateModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.BidResponse{
		ID:             bid.ID,
		Name:           bid.Name,
		Description:    bid.Description,
		TenderID:       bid.TenderID,
		OrganizationID: bid.OrganizationID,
		CreatorID:      bid.CreatorID,
		Status:         bid.Status,
		CreatedAt:      bid.CreatedAt,
		Version:        bid.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetMyBids returns bids for the logged-in user
func (h *BidHandler) GetMyBids(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	bids, err := h.service.GetBidsByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []models.BidResponse
	for _, bid := range bids {
		resp = append(resp, models.BidResponse{
			ID:             bid.ID,
			Name:           bid.Name,
			Description:    bid.Description,
			TenderID:       bid.TenderID,
			OrganizationID: bid.OrganizationID,
			CreatorID:      bid.CreatorID,
			Status:         bid.Status,
			CreatedAt:      bid.CreatedAt,
			Version:        bid.Version,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// GetBidsByTenderID returns bids associated with a specific tender
func (h *BidHandler) GetBidsByTenderID(w http.ResponseWriter, r *http.Request) {
	tenderIDStr := mux.Vars(r)["tenderId"]
	tenderID, err := strconv.Atoi(tenderIDStr)
	if err != nil {
		http.Error(w, "Invalid tender ID", http.StatusBadRequest)
		return
	}

	bids, err := h.service.GetBidsByTenderID(r.Context(), tenderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []models.BidResponse
	for _, bid := range bids {
		resp = append(resp, models.BidResponse{
			ID:             bid.ID,
			Name:           bid.Name,
			Description:    bid.Description,
			TenderID:       bid.TenderID,
			OrganizationID: bid.OrganizationID,
			CreatorID:      bid.CreatorID,
			Status:         bid.Status,
			CreatedAt:      bid.CreatedAt,
			Version:        bid.Version,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// UpdateBid edits a bid
func (h *BidHandler) UpdateBid(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bidUpdateModel := bid_service.BidUpdateModel{
		ID:          bidID,
		Name:        req.Name,
		Description: req.Description,
	}

	bid, err := h.service.UpdateBid(r.Context(), bidUpdateModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.BidResponse{
		ID:             bid.ID,
		Name:           bid.Name,
		Description:    bid.Description,
		TenderID:       bid.TenderID,
		OrganizationID: bid.OrganizationID,
		CreatorID:      bid.CreatorID,
		Status:         bid.Status,
		CreatedAt:      bid.CreatedAt,
		Version:        bid.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// RollbackBidVersion rolls back a bid to a specific version
func (h *BidHandler) RollbackBidVersion(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	versionStr := mux.Vars(r)["version"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		http.Error(w, "Invalid version number", http.StatusBadRequest)
		return
	}

	bid, err := h.service.RollbackBidVersion(r.Context(), bidID, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.BidResponse{
		ID:             bid.ID,
		Name:           bid.Name,
		Description:    bid.Description,
		TenderID:       bid.TenderID,
		OrganizationID: bid.OrganizationID,
		CreatorID:      bid.CreatorID,
		Status:         bid.Status,
		CreatedAt:      bid.CreatedAt,
		Version:        bid.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
