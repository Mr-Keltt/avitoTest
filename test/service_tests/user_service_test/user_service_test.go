package service_test

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/user_repository"
	"avitoTest/services/user_service"
	"avitoTest/services/user_service/user_models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupUserMocks() (*user_repository.MockUserRepository, user_service.UserService) {
	mockUserRepo := new(user_repository.MockUserRepository)
	service := user_service.NewUserService(mockUserRepo)
	return mockUserRepo, service
}

func TestCreateUser_Success(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	userCreate := user_models.UserCreateModel{
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
	}

	expectedEntity := &entities.User{
		ID:        1,
		Username:  userCreate.Username,
		FirstName: userCreate.FirstName,
		LastName:  userCreate.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*entities.User)
		user.ID = expectedEntity.ID
		user.CreatedAt = expectedEntity.CreatedAt
	})

	result, err := service.CreateUser(context.Background(), userCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, userCreate.Username, result.Username)
	mockUserRepo.AssertExpectations(t)
}

func TestCreateUser_ValidationFail(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	userCreate := user_models.UserCreateModel{
		Username:  "",
		FirstName: "John",
		LastName:  "Doe",
	}

	_, err := service.CreateUser(context.Background(), userCreate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation for 'Username' failed on the 'required' tag")

	mockUserRepo.AssertNotCalled(t, "Create")
}

// func TestGetUsers_Success(t *testing.T) {
// 	mockUserRepo, service := setupUserMocks()

// 	expectedUsers := []*entities.User{
// 		{
// 			ID:        1,
// 			Username:  "jdoe",
// 			FirstName: "John",
// 			LastName:  "Doe",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 		{
// 			ID:        2,
// 			Username:  "asmith",
// 			FirstName: "Alice",
// 			LastName:  "Smith",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		},
// 	}

// 	mockUserRepo.On("GetAll", mock.Anything).Return(expectedUsers, nil)

// 	result, err := service.GetUsers(context.Background())

// 	assert.NoError(t, err)

// 	assert.Len(t, result, 2)

// 	assert.Equal(t, "jdoe", result[0].Username)
// 	assert.Equal(t, "asmith", result[1].Username)

// 	mockUserRepo.AssertExpectations(t)
// }

func TestGetUserByID_Success(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	expectedUser := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserRepo.On("FindByID", mock.Anything, 1).Return(expectedUser, nil)

	result, err := service.GetUserByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Username, result.Username)
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	mockUserRepo.On("FindByID", mock.Anything, 1).Return(nil, user_repository.ErrUserNotFound)

	_, err := service.GetUserByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestGetUserByUsername_Success(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	expectedUser := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserRepo.On("FindByUsername", mock.Anything, "jdoe").Return(expectedUser, nil)

	result, err := service.GetUserByUsername(context.Background(), "jdoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Username, result.Username)
	mockUserRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	expectedUser := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockUserRepo.On("FindByID", mock.Anything, 1).Return(expectedUser, nil)

	mockUserRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteUser(context.Background(), 1)

	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	mockUserRepo, service := setupUserMocks()

	mockUserRepo.On("FindByID", mock.Anything, 1).Return(nil, user_repository.ErrUserNotFound)

	err := service.DeleteUser(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockUserRepo.AssertExpectations(t)
}
