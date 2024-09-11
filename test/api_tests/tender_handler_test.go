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
	"avitoTest/services/user_service/user_models"
	"avitoTest/shared/constants"

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

func (m *MockTenderService) GetAllTenders(ctx context.Context, serviceTypeFilter string) ([]*tender_models.TenderModel, error) {
	args := m.Called(ctx, serviceTypeFilter)
	return args.Get(0).([]*tender_models.TenderModel), args.Error(1)
}

func TestCreateTender(t *testing.T) {
	service := new(MockTenderService)
	userService := new(MockUserService)
	handler := tender_handler.NewTenderHandler(service, userService)

	// Mock GetUserByUsername for CreatorUsername
	userModel := &user_models.UserModel{ID: 1, Username: "CreatorUsername 1"}
	userService.On("GetUserByUsername", mock.Anything, "CreatorUsername 1").Return(userModel, nil)

	// Create request body with the correct ServiceType
	reqBody := &tender_handler_models.CreateTenderRequest{
		Name:            "Tender 1",
		Description:     "Description 1",
		ServiceType:     "Consulting", // Ensure this matches one of the constants
		OrganizationID:  1,
		CreatorUsername: "CreatorUsername 1",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/tenders/", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Expected response from the service mock
	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    string(constants.ServiceTypeConsulting), // Ensure ServiceType matches the constant
		Status:         constants.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        1,
	}

	// Mock CreateTender to return the expected response
	service.On("CreateTender", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	// Call the handler
	handler.CreateTender(rr, req)

	// Assert that the response code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response body
	var response tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	// Assert that the returned tender matches the expected values
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.ServiceType, response.ServiceType) // Check ServiceType

	// Verify that all expectations are met
	service.AssertExpectations(t)
	userService.AssertExpectations(t)
}

func TestGetTenderByID(t *testing.T) {
	tender_service := new(MockTenderService)
	user_service := new(MockUserService)
	handler := tender_handler.NewTenderHandler(tender_service, user_service)

	req := httptest.NewRequest("GET", "/api/tenders/1", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		Status:         constants.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        1,
	}

	tender_service.On("GetTenderByID", mock.Anything, 1).Return(expectedResponse, nil)

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

	tender_service.AssertExpectations(t)
}

func TestGetTenders(t *testing.T) {
	service := new(MockTenderService)
	userService := new(MockUserService)
	handler := tender_handler.NewTenderHandler(service, userService)

	req := httptest.NewRequest("GET", "/api/tenders", nil)
	rr := httptest.NewRecorder()

	expectedTenders := []*tender_models.TenderModel{
		{
			ID:             1,
			Name:           "Tender 1",
			Description:    "Description 1",
			ServiceType:    "Construction",
			Status:         constants.TenderStatusCreated,
			OrganizationID: 1,
			CreatedAt:      time.Now(),
			Version:        1,
		},
		{
			ID:             2,
			Name:           "Tender 2",
			Description:    "Description 2",
			ServiceType:    "Consulting",
			Status:         constants.TenderStatusPublished,
			OrganizationID: 2,
			CreatedAt:      time.Now(),
			Version:        1,
		},
	}

	service.On("GetAllTenders", mock.Anything, "").Return(expectedTenders, nil)

	handler.GetTenders(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Len(t, response, 2)
	assert.Equal(t, expectedTenders[0].ID, response[0].ID)
	assert.Equal(t, expectedTenders[1].ID, response[1].ID)

	service.AssertExpectations(t)
}

func TestPublishTender(t *testing.T) {
	tender_service := new(MockTenderService)
	user_service := new(MockUserService)
	handler := tender_handler.NewTenderHandler(tender_service, user_service)

	req := httptest.NewRequest("POST", "/api/tenders/1/publish", nil)
	rr := httptest.NewRecorder()

	tender_service.On("PublishTender", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.PublishTender(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	tender_service.AssertExpectations(t)
}

func TestCloseTender(t *testing.T) {
	tender_service := new(MockTenderService)
	user_service := new(MockUserService)
	handler := tender_handler.NewTenderHandler(tender_service, user_service)

	req := httptest.NewRequest("POST", "/api/tenders/1/close", nil)
	rr := httptest.NewRecorder()

	tender_service.On("CloseTender", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.CloseTender(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	tender_service.AssertExpectations(t)
}

func TestRollbackTenderVersion(t *testing.T) {
	tender_service := new(MockTenderService)
	user_service := new(MockUserService)
	handler := tender_handler.NewTenderHandler(tender_service, user_service)

	req := httptest.NewRequest("PUT", "/api/tenders/1/rollback/2", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &tender_models.TenderModel{
		ID:             1,
		Name:           "Tender 1 Version 2",
		Description:    "Description Version 2",
		ServiceType:    "Construction",
		Status:         constants.TenderStatusCreated,
		OrganizationID: 1,
		CreatedAt:      time.Now(),
		Version:        2,
	}

	tender_service.On("RollbackTenderVersion", mock.Anything, 1, 2).Return(expectedResponse, nil)

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

	tender_service.AssertExpectations(t)
}
