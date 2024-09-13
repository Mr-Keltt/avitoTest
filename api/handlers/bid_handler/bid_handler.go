package bid_handler

import (
	"avitoTest/api/handlers/bid_handler/bid_handler_models"
	"avitoTest/services/bid_service"
	"avitoTest/services/bid_service/bid_models"
	"avitoTest/shared"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BidHandler struct {
	service bid_service.BidService
}

func NewBidHandler(service bid_service.BidService) *BidHandler {
	return &BidHandler{service: service}
}

// CreateBid creates a new bid
func (h *BidHandler) CreateBid(w http.ResponseWriter, r *http.Request) {
	var req bid_handler_models.CreateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the TenderID to ensure it's being received correctly
	shared.Logger.Infof("Received TenderID: %d", req.TenderID)

	bidCreateModel := bid_models.BidCreateModel{
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

	resp := bid_handler_models.BidResponse{
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

// GetBidByID retrieves a bid by its ID
func (h *BidHandler) GetBidByID(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	bid, err := h.service.GetBidByID(r.Context(), bidID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := bid_handler_models.BidResponse{
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

	var resp []bid_handler_models.BidResponse
	for _, bid := range bids {
		resp = append(resp, bid_handler_models.BidResponse{
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

// GetBidsByUserID returns bids created by a specific user
func (h *BidHandler) GetBidsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userId"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	bids, err := h.service.GetBidsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []bid_handler_models.BidResponse
	for _, bid := range bids {
		resp = append(resp, bid_handler_models.BidResponse{
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

// GetBidsByUsername returns bids for a specific username
func (h *BidHandler) GetBidsByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	bids, err := h.service.GetBidsByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []bid_handler_models.BidResponse
	for _, bid := range bids {
		resp = append(resp, bid_handler_models.BidResponse{
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

// UpdateBid updates an existing bid.
func (h *BidHandler) UpdateBid(w http.ResponseWriter, r *http.Request) {
	// Extract bid ID from the URL path
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into UpdateBidRequest
	var req bid_handler_models.UpdateBidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the update model
	bidUpdateModel := bid_models.BidUpdateModel{
		ID:          bidID,
		Name:        req.Name,
		Description: req.Description,
	}

	// Call the service to update the bid
	bid, err := h.service.UpdateBid(r.Context(), bidUpdateModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	resp := bid_handler_models.BidResponse{
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

	// Return the updated bid
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// ApproveBid approves a bid
func (h *BidHandler) ApproveBid(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	approverIDStr := mux.Vars(r)["approverId"]
	approverID, err := strconv.Atoi(approverIDStr)
	if err != nil {
		http.Error(w, "Invalid approver ID", http.StatusBadRequest)
		return
	}

	err = h.service.ApproveBid(r.Context(), bidID, approverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RejectBid rejects a bid
func (h *BidHandler) RejectBid(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	rejecterIDStr := mux.Vars(r)["rejecterId"]
	rejecterID, err := strconv.Atoi(rejecterIDStr)
	if err != nil {
		http.Error(w, "Invalid rejecter ID", http.StatusBadRequest)
		return
	}

	err = h.service.RejectBid(r.Context(), bidID, rejecterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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

	resp := bid_handler_models.BidResponse{
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

// DeleteBid deletes a bid
func (h *BidHandler) DeleteBid(w http.ResponseWriter, r *http.Request) {
	bidIDStr := mux.Vars(r)["bidId"]
	bidID, err := strconv.Atoi(bidIDStr)
	if err != nil {
		http.Error(w, "Invalid bid ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBid(r.Context(), bidID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
