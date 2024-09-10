package user_service_test

import (
	"context"
	"testing"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/user_repository"
	"avitoTest/services/user_service"
	models "avitoTest/services/user_service/user_models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	args := m.Called(ctx)
	if users, ok := args.Get(0).([]entities.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*entities.User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	userCreate := models.UserCreateModel{
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

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*entities.User)
		user.ID = expectedEntity.ID
		user.CreatedAt = expectedEntity.CreatedAt
		user.UpdatedAt = expectedEntity.UpdatedAt
	})

	result, err := service.CreateUser(context.Background(), userCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Username, result.Username)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ValidationFail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	userCreate := models.UserCreateModel{
		Username:  "",
		FirstName: "John",
		LastName:  "Doe",
	}

	_, err := service.CreateUser(context.Background(), userCreate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed on the 'required' tag")
}

func TestGetUsers_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	expectedEntities := []entities.User{
		{
			ID:        1,
			Username:  "jdoe",
			FirstName: "John",
			LastName:  "Doe",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Username:  "asmith",
			FirstName: "Alice",
			LastName:  "Smith",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAll", mock.Anything).Return(expectedEntities, nil)

	result, err := service.GetUsers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedEntities[0].ID, result[0].ID)
	assert.Equal(t, expectedEntities[1].ID, result[1].ID)
	mockRepo.AssertExpectations(t)
}

func TestGetUsers_Failure(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("GetAll", mock.Anything).Return(nil, user_repository.ErrUserNotFound)

	_, err := service.GetUsers(context.Background())

	assert.Error(t, err)
	assert.Equal(t, user_repository.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	expectedEntity := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(expectedEntity, nil)

	result, err := service.GetUserByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Username, result.Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(nil, user_repository.ErrUserNotFound)

	_, err := service.GetUserByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	userUpdate := models.UserUpdateModel{
		ID:        1,
		Username:  "jdoe_updated",
		FirstName: "John",
		LastName:  "Doe",
	}

	existingEntity := &entities.User{
		ID:        1,
		Username:  "jdoe",
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.User")).Return(nil)

	result, err := service.UpdateUser(context.Background(), userUpdate)

	assert.NoError(t, err)
	assert.Equal(t, userUpdate.Username, result.Username)
	assert.Equal(t, userUpdate.FirstName, result.FirstName)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(&entities.User{ID: 1}, nil)
	mockRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteUser(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(nil, user_repository.ErrUserNotFound)

	err := service.DeleteUser(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}
