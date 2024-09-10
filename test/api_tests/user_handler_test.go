package api_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/user_handler"
	"avitoTest/api/handlers/user_handler/user_handler_models"
	"avitoTest/services/user_service/user_models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, user user_models.UserCreateModel) (*user_models.UserModel, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*user_models.UserModel), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, user user_models.UserUpdateModel) (*user_models.UserModel, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*user_models.UserModel), args.Error(1)
}

func (m *MockUserService) GetUsers(ctx context.Context) ([]*user_models.UserModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*user_models.UserModel), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id int) (*user_models.UserModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*user_models.UserModel), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	service := new(MockUserService)
	handler := user_handler.NewUserHandler(service)

	reqBody := &user_handler_models.CreateUserRequest{
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/users/", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	expectedResponse := &user_models.UserModel{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service.On("CreateUser", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	handler.CreateUser(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response user_handler_models.UserResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Username, response.Username)
	assert.Equal(t, expectedResponse.FirstName, response.FirstName)
	assert.Equal(t, expectedResponse.LastName, response.LastName)

	service.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	service := new(MockUserService)
	handler := user_handler.NewUserHandler(service)

	reqBody := &user_handler_models.UpdateUserRequest{
		Username:  "jdoe_updated",
		FirstName: "Johnathan",
		LastName:  "Doe",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("PATCH", "/api/users/1", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	expectedResponse := &user_models.UserModel{
		ID:        1,
		Username:  "jdoe_updated",
		FirstName: "Johnathan",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service.On("UpdateUser", mock.Anything, mock.Anything).Return(expectedResponse, nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.UpdateUser(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response user_handler_models.UserResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Username, response.Username)
	assert.Equal(t, expectedResponse.FirstName, response.FirstName)
	assert.Equal(t, expectedResponse.LastName, response.LastName)

	service.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	service := new(MockUserService)
	handler := user_handler.NewUserHandler(service)

	req := httptest.NewRequest("GET", "/api/users/1", nil)
	rr := httptest.NewRecorder()

	expectedResponse := &user_models.UserModel{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service.On("GetUserByID", mock.Anything, 1).Return(expectedResponse, nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.GetUserByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response user_handler_models.UserResponse
	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Username, response.Username)
	assert.Equal(t, expectedResponse.FirstName, response.FirstName)
	assert.Equal(t, expectedResponse.LastName, response.LastName)

	service.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	service := new(MockUserService)
	handler := user_handler.NewUserHandler(service)

	req := httptest.NewRequest("DELETE", "/api/users/1", nil)
	rr := httptest.NewRecorder()

	service.On("DeleteUser", mock.Anything, 1).Return(nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	handler.DeleteUser(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	service.AssertExpectations(t)
}
