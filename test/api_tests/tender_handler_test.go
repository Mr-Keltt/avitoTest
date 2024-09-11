// File: test/api_tests/tender_handler_test.go

package api_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/tender_handler"
	"avitoTest/api/handlers/tender_handler/tender_handler_models"
	"avitoTest/services/tender_service/tender_models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTenderService is a mock implementation of the TenderService interface
type MockTenderService struct {
	mock.Mock
}

func (m *MockTenderService) CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tender)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tender)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) PublishTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderService) CloseTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderService) GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tenderID, version)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) GetAllTenders(ctx context.Context) ([]*tender_models.TenderModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*tender_models.TenderModel), args.Error(1)
}

func TestCreateTender(t *testing.T) {
	service := new(MockTenderService)
	handler := tender_handler.NewTenderHandler(service)

	reqBody := &tender_handler_models.CreateTenderRequest{
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		OrganizationID: 1,
		CreatorID:      1,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/tenders/", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		Status:         tender_models.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        1,
	}

	service.On("CreateTender", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	handler.CreateTender(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.ServiceType, response.ServiceType)

	service.AssertExpectations(t)
}

func TestGetTenderByID(t *testing.T) {
	service := new(MockTenderService)
	handler := tender_handler.NewTenderHandler(service)

	req := httptest.NewRequest("GET", "/api/tenders/1", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		Status:         tender_models.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        1,
	}

	service.On("GetTenderByID", mock.Anything, 1).Return(expectedResponse, nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.GetTenderByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)

	service.AssertExpectations(t)
}

func TestPublishTender(t *testing.T) {
	service := new(MockTenderService)
	handler := tender_handler.NewTenderHandler(service)

	req := httptest.NewRequest("POST", "/api/tenders/1/publish", nil)
	rr := httptest.NewRecorder()

	service.On("PublishTender", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.PublishTender(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	service.AssertExpectations(t)
}

func TestCloseTender(t *testing.T) {
	service := new(MockTenderService)
	handler := tender_handler.NewTenderHandler(service)

	req := httptest.NewRequest("POST", "/api/tenders/1/close", nil)
	rr := httptest.NewRecorder()

	service.On("CloseTender", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.CloseTender(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	service.AssertExpectations(t)
}

func TestRollbackTenderVersion(t *testing.T) {
	service := new(MockTenderService)
	handler := tender_handler.NewTenderHandler(service)

	req := httptest.NewRequest("PUT", "/api/tenders/1/rollback/2", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1 Version 2",
		Description:    "Description Version 2",
		ServiceType:    "Construction",
		Status:         tender_models.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        2,
	}

	service.On("RollbackTenderVersion", mock.Anything, 1, 2).Return(expectedResponse, nil)

	vars := map[string]string{
		"tenderId": "1",
		"version":  "2",
	}
	req = mux.SetURLVars(req, vars)

	handler.RollbackTenderVersion(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)

	service.AssertExpectations(t)
}
