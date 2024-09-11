package user_service_test

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

func (m *MockUserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	args := m.Called(ctx, username)
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

// Test for GetUserByUsername
func TestGetUserByUsername_Success(t *testing.T) {
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

	mockRepo.On("FindByUsername", mock.Anything, "jdoe").Return(expectedEntity, nil)

	result, err := service.GetUserByUsername(context.Background(), "jdoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Username, result.Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := user_service.NewUserService(mockRepo)

	mockRepo.On("FindByUsername", mock.Anything, "unknown").Return(nil, user_repository.ErrUserNotFound)

	_, err := service.GetUserByUsername(context.Background(), "unknown")

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}
