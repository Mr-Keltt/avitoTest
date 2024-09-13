package api_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"avitoTest/api/handlers/user_handler"
	"avitoTest/api/handlers/user_handler/user_handler_models"
	"avitoTest/services/user_service"
	"avitoTest/services/user_service/user_models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupMocks() (*user_service.MockUserService, *user_handler.UserHandler) {
	service := new(user_service.MockUserService)
	handler := user_handler.NewUserHandler(service)
	return service, handler
}

func TestCreateUser(t *testing.T) {
	service, handler := setupMocks()

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
	service, handler := setupMocks()

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
	service, handler := setupMocks()

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
	service, handler := setupMocks()

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
