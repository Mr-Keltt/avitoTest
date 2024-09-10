package user_repository

import (
	"avitoTest/data/entities"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetAll(ctx context.Context) ([]entities.User, error)
	FindByID(ctx context.Context, id int) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id int) error
}
