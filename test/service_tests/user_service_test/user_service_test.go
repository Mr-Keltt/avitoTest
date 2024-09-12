package service_test

import (
	"context"
	"testing"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/user_repository"
	"avitoTest/services/user_service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test for GetUserByUsername
func TestGetUserByUsername_Success(t *testing.T) {
	mockRepo := new(user_repository.MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	expectedEntity := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByUsername", mock.Anything, "jdoe").Return(expectedEntity, nil)

	result, err := service.GetUserByUsername(context.Background(), "jdoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Username, result.Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	mockRepo := new(user_repository.MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("FindByUsername", mock.Anything, "unknown").Return(nil, user_repository.ErrUserNotFound)

	_, err := service.GetUserByUsername(context.Background(), "unknown")

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}
