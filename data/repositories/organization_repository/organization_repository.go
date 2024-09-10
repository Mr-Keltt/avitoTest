package organization_repository

import (
	"avitoTest/data/entities"
	"context"
)

type OrganizationRepository interface {
	Create(ctx context.Context, org *entities.Organization) error
	GetAll(ctx context.Context) ([]entities.Organization, error)
	FindByID(ctx context.Context, id int) (*entities.Organization, error)
	Update(ctx context.Context, org *entities.Organization) error
	Delete(ctx context.Context, id int) error
	AddResponsible(ctx context.Context, orgResponsible *entities.OrganizationResponsible) error
	DeleteResponsible(ctx context.Context, orgID int, userID int) error
	GetResponsibles(ctx context.Context, orgID int) ([]entities.User, error)
}
