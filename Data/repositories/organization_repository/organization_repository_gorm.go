package organization_repository

import (
	"avitoTest/data/entities"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrOrganizationNotFound = errors.New("organization not found")

type OrganizationRepositoryGorm struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &OrganizationRepositoryGorm{db: db}
}

func (r *OrganizationRepositoryGorm) Create(ctx context.Context, org *entities.Organization) error {
	return r.db.WithContext(ctx).Create(org).Error
}

func (r *OrganizationRepositoryGorm) GetAll(ctx context.Context) ([]entities.Organization, error) {
	var organizations []entities.Organization
	if err := r.db.WithContext(ctx).Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *OrganizationRepositoryGorm) FindByID(ctx context.Context, id int) (*entities.Organization, error) {
	var org entities.Organization
	if err := r.db.WithContext(ctx).First(&org, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrganizationNotFound
		}
		return nil, err
	}
	return &org, nil
}

func (r *OrganizationRepositoryGorm) Update(ctx context.Context, org *entities.Organization) error {
	return r.db.WithContext(ctx).Save(org).Error
}

func (r *OrganizationRepositoryGorm) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entities.Organization{}, id).Error
}
