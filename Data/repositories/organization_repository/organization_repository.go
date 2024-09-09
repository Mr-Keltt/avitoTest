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
}
