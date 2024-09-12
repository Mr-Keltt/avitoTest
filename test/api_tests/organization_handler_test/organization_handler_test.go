package api_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/organization_handler"
	"avitoTest/api/handlers/organization_handler/organization_handler_models"
	"avitoTest/services/organization_service"
	"avitoTest/services/organization_service/organization_models"
	"avitoTest/services/user_service/user_models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupMocks() (*organization_service.MockOrganizationService, *organization_handler.OrganizationHandler) {
	service := new(organization_service.MockOrganizationService)
	handler := organization_handler.NewOrganizationHandler(service)
	return service, handler
}

func TestCreateOrganization(t *testing.T) {
	service, handler := setupMocks()

	reqBody := &organization_handler_models.CreateOrganizationRequest{
		Name:        "Test Organization",
		Description: "A test organization",
		Type:        "LLC",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/organizations/", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	expectedResponse := &organization_models.OrganizationModel{
		ID:          1,
		Name:        "Test Organization",
		Description: "A test organization",
		Type:        "LLC",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	service.On("CreateOrganization", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	handler.CreateOrganization(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response organization_handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestUpdateOrganization(t *testing.T) {
	service, handler := setupMocks()

	reqBody := &organization_handler_models.UpdateOrganizationRequest{
		Name:        "Updated Organization",
		Description: "An updated organization",
		Type:        "LLC",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("PATCH", "/api/organizations/1", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	expectedResponse := &organization_models.OrganizationModel{
		ID:          1,
		Name:        "Updated Organization",
		Description: "An updated organization",
		Type:        "LLC",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	service.On("UpdateOrganization", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.UpdateOrganization(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response organization_handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestGetOrganizationByID(t *testing.T) {
	service, handler := setupMocks()

	req := httptest.NewRequest("GET", "/api/organizations/1", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &organization_models.OrganizationModel{
		ID:          1,
		Name:        "Test Organization",
		Description: "A test organization",
		Type:        "LLC",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	service.On("GetOrganizationByID", mock.Anything, 1).Return(expectedResponse, nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.GetOrganizationByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response organization_handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestDeleteOrganization(t *testing.T) {
	service, handler := setupMocks()

	req := httptest.NewRequest("DELETE", "/api/organizations/1", nil)
	rr := httptest.NewRecorder()

	service.On("DeleteOrganization", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.DeleteOrganization(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	service.AssertExpectations(t)
}

func TestAddResponsible(t *testing.T) {
	service, handler := setupMocks()

	req := httptest.NewRequest("POST", "/api/organizations/1/responsibles/1", nil)
	rr := httptest.NewRecorder()

	service.On("AddResponsible", mock.Anything, 1, 1).Return(nil)

	vars := map[string]string{
		"org_id":  "1",
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.AddResponsible(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	service.AssertExpectations(t)
}

func TestDeleteResponsible(t *testing.T) {
	service, handler := setupMocks()

	req := httptest.NewRequest("DELETE", "/api/organizations/1/responsibles/1", nil)
	rr := httptest.NewRecorder()

	service.On("DeleteResponsible", mock.Anything, 1, 1).Return(nil)

	vars := map[string]string{
		"org_id":  "1",
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.DeleteResponsible(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	service.AssertExpectations(t)
}

func TestGetResponsibles(t *testing.T) {
	service, handler := setupMocks()

	req := httptest.NewRequest("GET", "/api/organizations/1/responsibles", nil)
	rr := httptest.NewRecorder()

	expectedUsers := []*user_models.UserModel{
		{
			ID:        1,
			Username:  "jdoe",
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	service.On("GetResponsibles", mock.Anything, 1).Return(expectedUsers, nil)

	vars := map[string]string{
		"org_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.GetResponsibles(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []*user_models.UserModel
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, len(expectedUsers), len(response))
	assert.Equal(t, expectedUsers[0].ID, response[0].ID)
	assert.Equal(t, expectedUsers[0].Username, response[0].Username)

	service.AssertExpectations(t)
}

func TestGetResponsibleByID(t *testing.T) {
	service, handler := setupMocks()

	// Set up mock response
	expectedUser := &user_models.UserModel{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
	}

	// Set up the request and response
	req := httptest.NewRequest("GET", "/api/organizations/1/responsibles/1", nil)
	rr := httptest.NewRecorder()

	// Mock the service call
	service.On("GetResponsibleByID", mock.Anything, 1, 1).Return(expectedUser, nil)

	// Set URL variables
	vars := map[string]string{
		"org_id":  "1",
		"user_id": "1",
	}
	req = mux.SetURLVars(req, vars)

	// Call the handler
	handler.GetResponsibleByID(rr, req)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body
	var response user_models.UserModel
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedUser.ID, response.ID)
	assert.Equal(t, expectedUser.Username, response.Username)
	assert.Equal(t, expectedUser.FirstName, response.FirstName)
	assert.Equal(t, expectedUser.LastName, response.LastName)

	// Assert that the expectations were met
	service.AssertExpectations(t)
}
