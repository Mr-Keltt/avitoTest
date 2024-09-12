package user_service

import (
	user_models "avitoTest/services/user_service/user_models"
	"context"

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

func (m *MockUserService) GetUserByUsername(ctx context.Context, username string) (*user_models.UserModel, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*user_models.UserModel), args.Error(1)
}
