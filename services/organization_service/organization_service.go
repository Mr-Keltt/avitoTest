package organization_service

import (
	"context"
	"errors"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/data/repositories/user_repository"
	"avitoTest/services/organization_service/organization_models"
	"avitoTest/services/user_service/user_models"

	"github.com/go-playground/validator/v10"
)

// OrganizationService defines the interface for organization operations.
type OrganizationService interface {
	CreateOrganization(ctx context.Context, org organization_models.OrganizationCreateModel) (*organization_models.OrganizationModel, error)
	UpdateOrganization(ctx context.Context, org organization_models.OrganizationUpdateModel) (*organization_models.OrganizationModel, error)
	GetOrganizations(ctx context.Context) ([]*organization_models.OrganizationModel, error)
	GetOrganizationByID(ctx context.Context, id int) (*organization_models.OrganizationModel, error)
	DeleteOrganization(ctx context.Context, id int) error
	AddResponsible(ctx context.Context, orgID int, userID int) error
	DeleteResponsible(ctx context.Context, orgID int, userID int) error
	GetResponsibles(ctx context.Context, orgID int) ([]*user_models.UserModel, error)
	GetResponsibleByID(ctx context.Context, orgID int, userID int) (*user_models.UserModel, error)
}

type organizationService struct {
	orgRepo  organization_repository.OrganizationRepository
	userRepo user_repository.UserRepository
	validate *validator.Validate
}

// NewOrganizationService creates a new instance of OrganizationService.
func NewOrganizationService(orgRepo organization_repository.OrganizationRepository, userRepo user_repository.UserRepository) OrganizationService {
	return &organizationService{
		orgRepo:  orgRepo,
		userRepo: userRepo,
		validate: validator.New(),
	}
}

// CreateOrganization creates a new organization.
func (s *organizationService) CreateOrganization(ctx context.Context, org organization_models.OrganizationCreateModel) (*organization_models.OrganizationModel, error) {
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

	if err := s.orgRepo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return &organization_models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

// UpdateOrganization updates an existing organization.
func (s *organizationService) UpdateOrganization(ctx context.Context, org organization_models.OrganizationUpdateModel) (*organization_models.OrganizationModel, error) {
	if err := s.validate.Struct(org); err != nil {
		return nil, err
	}

	entity, err := s.orgRepo.FindByID(ctx, org.ID)
	if err != nil {
		return nil, err
	}

	entity.Name = org.Name
	entity.Description = org.Description
	entity.Type = entities.OrganizationType(org.Type)
	entity.UpdatedAt = time.Now()

	if err := s.orgRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return &organization_models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

// GetOrganizations retrieves all organizations.
func (s *organizationService) GetOrganizations(ctx context.Context) ([]*organization_models.OrganizationModel, error) {
	entities, err := s.orgRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var organizations []*organization_models.OrganizationModel
	for _, entity := range entities {
		orgModel := &organization_models.OrganizationModel{
			ID:          entity.ID,
			Name:        entity.Name,
			Description: entity.Description,
			Type:        string(entity.Type),
			CreatedAt:   entity.CreatedAt,
			UpdatedAt:   entity.UpdatedAt,
		}
		organizations = append(organizations, orgModel)
	}

	return organizations, nil
}

// GetOrganizationByID retrieves an organization by its ID.
func (s *organizationService) GetOrganizationByID(ctx context.Context, id int) (*organization_models.OrganizationModel, error) {
	entity, err := s.orgRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, organization_repository.ErrOrganizationNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, err
	}

	return &organization_models.OrganizationModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Type:        string(entity.Type),
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

// DeleteOrganization deletes an organization by its ID.
func (s *organizationService) DeleteOrganization(ctx context.Context, id int) error {
	entity, err := s.orgRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, organization_repository.ErrOrganizationNotFound) {
			return errors.New("organization not found")
		}
		return err
	}

	if err := s.orgRepo.Delete(ctx, entity.ID); err != nil {
		return err
	}

	return nil
}

// AddResponsible adds a user as a responsible for an organization.
func (s *organizationService) AddResponsible(ctx context.Context, orgID int, userID int) error {
	// Validate that both organization and user exist
	org, err := s.orgRepo.FindByID(ctx, orgID)
	if err != nil {
		return errors.New("organization not found")
	}

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Create OrganizationResponsible entity
	orgResponsible := entities.OrganizationResponsible{
		OrganizationID: org.ID,
		UserID:         user.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Add responsible
	return s.orgRepo.AddResponsible(ctx, &orgResponsible)
}

// DeleteResponsible removes a user as a responsible for an organization.
func (s *organizationService) DeleteResponsible(ctx context.Context, orgID int, userID int) error {
	// Validate that both organization and user exist
	_, err := s.orgRepo.FindByID(ctx, orgID)
	if err != nil {
		return errors.New("organization not found")
	}

	_, err = s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Delete responsible
	return s.orgRepo.DeleteResponsible(ctx, orgID, userID)
}

// GetResponsibles retrieves all users responsible for a given organization.
func (s *organizationService) GetResponsibles(ctx context.Context, orgID int) ([]*user_models.UserModel, error) {
	// Fetch responsibles for the given organization
	responsibles, err := s.orgRepo.GetResponsibles(ctx, orgID)
	if err != nil {
		return nil, err
	}

	var userModels []*user_models.UserModel
	for _, user := range responsibles {
		userModel := &user_models.UserModel{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		userModels = append(userModels, userModel)
	}

	return userModels, nil
}

// GetResponsibleByID retrieves a specific responsible user by organization ID and user ID.
func (s *organizationService) GetResponsibleByID(ctx context.Context, orgID int, userID int) (*user_models.UserModel, error) {
	// Validate that the organization exists
	_, err := s.orgRepo.FindByID(ctx, orgID)
	if err != nil {
		return nil, errors.New("organization not found")
	}

	// Fetch the responsible user for the organization by user ID
	responsible, err := s.orgRepo.GetResponsibleByID(ctx, orgID, userID)
	if err != nil {
		if errors.Is(err, organization_repository.ErrResponsibleNotFound) {
			return nil, errors.New("responsible user not found for the organization")
		}
		return nil, err
	}

	// Return the responsible user information
	return &user_models.UserModel{
		ID:        responsible.ID,
		Username:  responsible.Username,
		FirstName: responsible.FirstName,
		LastName:  responsible.LastName,
	}, nil
}
