package user_service

import (
	"context"

	user_models "avitoTest/services/user_service/user_models"
)

type UserService interface {
	CreateUser(ctx context.Context, user user_models.UserCreateModel) (*user_models.UserModel, error)
	GetUsers(ctx context.Context) ([]*user_models.UserModel, error)
	GetUserByID(ctx context.Context, id int) (*user_models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*user_models.UserModel, error)
	UpdateUser(ctx context.Context, user user_models.UserUpdateModel) (*user_models.UserModel, error)
	DeleteUser(ctx context.Context, id int) error
}
