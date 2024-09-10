package user_service

import (
	"context"
	"errors"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/user_repository"
	models "avitoTest/services/user_service/models"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.UserCreateModel) (*models.UserModel, error)
	UpdateUser(ctx context.Context, user models.UserUpdateModel) (*models.UserModel, error)
	GetUsers(ctx context.Context) ([]*models.UserModel, error)
	GetUserByID(ctx context.Context, id int) (*models.UserModel, error)
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repo     user_repository.UserRepository
	validate *validator.Validate
}

func NewUserService(repo user_repository.UserRepository) UserService {
	return &userService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *userService) CreateUser(ctx context.Context, user models.UserCreateModel) (*models.UserModel, error) {
	if err := s.validate.Struct(user); err != nil {
		return nil, err
	}

	entity := &entities.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return &models.UserModel{
		ID:        entity.ID,
		Username:  entity.Username,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, user models.UserUpdateModel) (*models.UserModel, error) {
	if err := s.validate.Struct(user); err != nil {
		return nil, err
	}

	entity, err := s.repo.FindByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	entity.Username = user.Username
	entity.FirstName = user.FirstName
	entity.LastName = user.LastName
	entity.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return &models.UserModel{
		ID:        entity.ID,
		Username:  entity.Username,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*models.UserModel, error) {
	entities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var users []*models.UserModel
	for _, entity := range entities {
		userModel := &models.UserModel{
			ID:        entity.ID,
			Username:  entity.Username,
			FirstName: entity.FirstName,
			LastName:  entity.LastName,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
		users = append(users, userModel)
	}

	return users, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*models.UserModel, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &models.UserModel{
		ID:        entity.ID,
		Username:  entity.Username,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := s.repo.Delete(ctx, entity.ID); err != nil {
		return err
	}

	return nil
}
