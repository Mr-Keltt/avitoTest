// services/organization_service/organization_service.go
package organization_service

import (
	"context"
	"errors"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service/models"

	"github.com/go-playground/validator/v10"
)

type OrganizationService interface {
	CreateOrganization(ctx context.Context, org models.OrganizationCreateModel) (*models.OrganizationModel, error)
	UpdateOrganization(ctx context.Context, org models.OrganizationUpdateModel) (*models.OrganizationModel, error)
	GetOrganizationByID(ctx context.Context, id int) (*models.OrganizationModel, error)
	DeleteOrganization(ctx context.Context, id int) error
}

type organizationService struct {
	repo     organization_repository.OrganizationRepository
	validate *validator.Validate
}

func NewOrganizationService(repo organization_repository.OrganizationRepository) OrganizationService {
	return &organizationService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *organizationService) CreateOrganization(ctx context.Context, org models.OrganizationCreateModel) (*models.OrganizationModel, error) {
	if err := s.validate.Struct(org); err != nil {
		return nil, err
	}

	entity := &entities.Organization{
		Name:        org.Name,
		Description: org.Description,
		Type:        entities.OrganizationType(org.Type),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return &models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *organizationService) UpdateOrganization(ctx context.Context, org models.OrganizationUpdateModel) (*models.OrganizationModel, error) {
	if err := s.validate.Struct(org); err != nil {
		return nil, err
	}

	entity, err := s.repo.FindByID(ctx, org.ID)
	if err != nil {
		return nil, err
	}

	entity.Name = org.Name
	entity.Description = org.Description
	entity.Type = entities.OrganizationType(org.Type)
	entity.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return &models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *organizationService) GetOrganizationByID(ctx context.Context, id int) (*models.OrganizationModel, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, organization_repository.ErrOrganizationNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, err
	}

	return &models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *organizationService) DeleteOrganization(ctx context.Context, id int) error {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, organization_repository.ErrOrganizationNotFound) {
			return errors.New("organization not found")
		}
		return err
	}

	if err := s.repo.Delete(ctx, entity.ID); err != nil {
		return err
	}

	return nil
}
