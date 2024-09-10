package organization_repository

import (
	"avitoTest/data/entities"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrOrganizationNotFound = errors.New("organization not found")
var ErrResponsibleNotFound = errors.New("responsible not found")

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

func (r *OrganizationRepositoryGorm) AddResponsible(ctx context.Context, orgResponsible *entities.OrganizationResponsible) error {
	return r.db.WithContext(ctx).Create(orgResponsible).Error
}

func (r *OrganizationRepositoryGorm) DeleteResponsible(ctx context.Context, orgID int, userID int) error {
	return r.db.WithContext(ctx).
		Where("organization_id = ? AND user_id = ?", orgID, userID).
		Delete(&entities.OrganizationResponsible{}).Error
}

func (r *OrganizationRepositoryGorm) GetResponsibles(ctx context.Context, orgID int) ([]entities.User, error) {
	var responsibles []entities.User
	err := r.db.WithContext(ctx).
		Model(&entities.OrganizationResponsible{}).
		Where("organization_id = ?", orgID).
		Joins("JOIN users ON users.id = organization_responsibles.user_id").
		Select("users.*").
		Find(&responsibles).Error

	if err != nil {
		return nil, err
	}

	return responsibles, nil
}

// GetResponsibleByID retrieves a responsible user by organization ID and user ID.
func (r *OrganizationRepositoryGorm) GetResponsibleByID(ctx context.Context, orgID int, userID int) (*entities.User, error) {
	var responsible entities.User
	err := r.db.WithContext(ctx).
		Model(&entities.OrganizationResponsible{}).
		Where("organization_id = ? AND user_id = ?", orgID, userID).
		Joins("JOIN users ON users.id = organization_responsibles.user_id").
		Select("users.*").
		First(&responsible).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("responsible user not found")
		}
		return nil, err
	}

	return &responsible, nil
}
