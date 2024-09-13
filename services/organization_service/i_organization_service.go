package organization_service

import (
	"context"

	"avitoTest/services/organization_service/organization_models"
	"avitoTest/services/user_service/user_models"
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
