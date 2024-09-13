// File: test/api_tests/tender_handler_test.go

package api_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/tender_handler"
	"avitoTest/api/handlers/tender_handler/tender_handler_models"
	"avitoTest/services/tender_service"
	"avitoTest/services/tender_service/tender_models"
	"avitoTest/services/user_service"
	"avitoTest/services/user_service/user_models"
	"avitoTest/shared/constants"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupMocks() (*tender_service.MockTenderService, *user_service.MockUserService, *tender_handler.TenderHandler) {
	service := new(tender_service.MockTenderService)
	userService := new(user_service.MockUserService)
	handler := tender_handler.NewTenderHandler(service, userService)
	return service, userService, handler
}

func TestCreateTender(t *testing.T) {
	service, userService, handler := setupMocks()

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
	service, _, handler := setupMocks()

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

func TestGetTenders(t *testing.T) {
	service, _, handler := setupMocks()

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
	service, _, handler := setupMocks()

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
	service, _, handler := setupMocks()

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
	service, _, handler := setupMocks()

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

func TestGetTendersByUsername(t *testing.T) {
	service, _, handler := setupMocks()

	req := httptest.NewRequest("GET", "/api/tenders/username/testuser", nil)
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

	service.On("GetTendersByUsername", mock.Anything, "testuser").Return(expectedTenders, nil)

	vars := map[string]string{
		"username": "testuser",
	}
	req = mux.SetURLVars(req, vars)

	handler.GetTendersByUsername(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []tender_handler_models.TenderResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Len(t, response, 2)
	assert.Equal(t, expectedTenders[0].ID, response[0].ID)
	assert.Equal(t, expectedTenders[1].ID, response[1].ID)

	service.AssertExpectations(t)
}

func TestDeleteTender(t *testing.T) {
	service, _, handler := setupMocks()

	req := httptest.NewRequest("DELETE", "/api/tenders/1", nil)
	rr := httptest.NewRecorder()

	service.On("DeleteTender", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"tenderId": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.DeleteTender(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	service.AssertExpectations(t)
}
