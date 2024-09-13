package bid_handler_test

import (
	"avitoTest/api/handlers/bid_handler"
	"avitoTest/api/handlers/bid_handler/bid_handler_models"
	"avitoTest/services/bid_service"
	"avitoTest/services/bid_service/bid_models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestHandler() (*bid_handler.BidHandler, *bid_service.MockBidService) {
	mockService := new(bid_service.MockBidService)
	handler := bid_handler.NewBidHandler(mockService)
	return handler, mockService
}

// Test CreateBid endpoint
func TestCreateBid_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	reqBody := bid_handler_models.CreateBidRequest{
		Name:           "New Bid",
		Description:    "Description for new bid",
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	expectedBid := &bid_models.BidModel{
		ID:             1,
		Name:           "New Bid",
		Description:    "Description for new bid",
		Status:         "CREATED",
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
		Version:        1,
	}

	mockService.On("CreateBid", mock.Anything, mock.AnythingOfType("bid_models.BidCreateModel")).Return(expectedBid, nil)

	req := httptest.NewRequest("POST", "/bids", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateBid(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var resp bid_handler_models.BidResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedBid.Name, resp.Name)
	assert.Equal(t, expectedBid.ID, resp.ID)
	mockService.AssertExpectations(t)
}

// Test UpdateBid endpoint
func TestUpdateBid_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	reqBody := bid_handler_models.UpdateBidRequest{
		Name:        "Updated Bid",
		Description: "Updated Description",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	expectedBid := &bid_models.BidModel{
		ID:          1,
		Name:        "Updated Bid",
		Description: "Updated Description",
		Status:      "UPDATED",
		Version:     2,
	}

	mockService.On("UpdateBid", mock.Anything, mock.AnythingOfType("bid_models.BidUpdateModel")).Return(expectedBid, nil)

	req := httptest.NewRequest("PUT", "/bids/1", bytes.NewReader(reqBodyBytes))
	req = mux.SetURLVars(req, map[string]string{"bidId": "1"})
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.UpdateBid(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp bid_handler_models.BidResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedBid.Name, resp.Name)
	assert.Equal(t, expectedBid.ID, resp.ID)
	mockService.AssertExpectations(t)
}

// Test GetBidByID endpoint
func TestGetBidByID_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	expectedBid := &bid_models.BidModel{
		ID:          1,
		Name:        "Test Bid",
		Description: "Description",
		Status:      "CREATED",
		Version:     1,
	}

	mockService.On("GetBidByID", mock.Anything, 1).Return(expectedBid, nil)

	req := httptest.NewRequest("GET", "/bids/1", nil)
	req = mux.SetURLVars(req, map[string]string{"bidId": "1"})
	rr := httptest.NewRecorder()

	handler.GetBidByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp bid_handler_models.BidResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedBid.Name, resp.Name)
	assert.Equal(t, expectedBid.ID, resp.ID)
	mockService.AssertExpectations(t)
}

// Test DeleteBid endpoint
func TestDeleteBid_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	mockService.On("DeleteBid", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/bids/1", nil)
	req = mux.SetURLVars(req, map[string]string{"bidId": "1"})
	rr := httptest.NewRecorder()

	handler.DeleteBid(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

// Test ApproveBid endpoint
func TestApproveBid_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	mockService.On("ApproveBid", mock.Anything, 1, 1).Return(nil)

	req := httptest.NewRequest("POST", "/bids/1/approve/1", nil)
	req = mux.SetURLVars(req, map[string]string{
		"bidId":      "1",
		"approverId": "1",
	})
	rr := httptest.NewRecorder()

	handler.ApproveBid(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

// Test RejectBid endpoint
func TestRejectBid_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	mockService.On("RejectBid", mock.Anything, 1, 1).Return(nil)

	req := httptest.NewRequest("POST", "/bids/1/reject/1", nil)
	req = mux.SetURLVars(req, map[string]string{
		"bidId":      "1",
		"rejecterId": "1",
	})
	rr := httptest.NewRecorder()

	handler.RejectBid(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

// Test RollbackBidVersion endpoint
func TestRollbackBidVersion_Success(t *testing.T) {
	handler, mockService := setupTestHandler()

	expectedBid := &bid_models.BidModel{
		ID:          1,
		Name:        "Rollback Bid",
		Description: "Rolled back description",
		Status:      "CREATED",
		Version:     2,
	}

	mockService.On("RollbackBidVersion", mock.Anything, 1, 1).Return(expectedBid, nil)

	req := httptest.NewRequest("POST", "/bids/1/rollback/1", nil)
	req = mux.SetURLVars(req, map[string]string{
		"bidId":   "1",
		"version": "1",
	})
	rr := httptest.NewRecorder()

	handler.RollbackBidVersion(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp bid_handler_models.BidResponse
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)

	assert.Equal(t, expectedBid.Name, resp.Name)
	assert.Equal(t, expectedBid.ID, resp.ID)
	mockService.AssertExpectations(t)
}
