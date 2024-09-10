package api_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/organization_handler"
	"avitoTest/api/handlers/organization_handler/handler_models"
	"avitoTest/services/organization_service/organization_models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockOrganizationService is a mock implementation of the OrganizationService interface
type MockOrganizationService struct {
	mock.Mock
}

func (m *MockOrganizationService) CreateOrganization(ctx context.Context, org organization_models.OrganizationCreateModel) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, org)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) UpdateOrganization(ctx context.Context, org organization_models.OrganizationUpdateModel) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, org)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) GetOrganizations(ctx context.Context) ([]*organization_models.OrganizationModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) GetOrganizationByID(ctx context.Context, id int) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) DeleteOrganization(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateOrganization(t *testing.T) {
	service := new(MockOrganizationService)
	handler := organization_handler.NewOrganizationHandler(service)

	reqBody := &handler_models.CreateOrganizationRequest{
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

	var response handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestUpdateOrganization(t *testing.T) {
	service := new(MockOrganizationService)
	handler := organization_handler.NewOrganizationHandler(service)

	reqBody := &handler_models.UpdateOrganizationRequest{
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

	var response handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestGetOrganizationByID(t *testing.T) {
	service := new(MockOrganizationService)
	handler := organization_handler.NewOrganizationHandler(service)

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

	var response handler_models.OrganizationResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.Type, response.Type)

	service.AssertExpectations(t)
}

func TestDeleteOrganization(t *testing.T) {
	service := new(MockOrganizationService)
	handler := organization_handler.NewOrganizationHandler(service)

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
