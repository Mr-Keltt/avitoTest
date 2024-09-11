// File: data/repositories/tender_repository/tender_repository_gorm.go

package tender_repository

import (
	"avitoTest/data/entities"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrTenderNotFound = errors.New("tender not found")

type tenderRepositoryGorm struct {
	db *gorm.DB
}

// NewTenderRepository creates a new GORM-based repository for tenders.
func NewTenderRepository(db *gorm.DB) TenderRepository {
	return &tenderRepositoryGorm{db: db}
}

// Create adds a new tender to the database.
func (r *tenderRepositoryGorm) Create(ctx context.Context, tender *entities.Tender) error {
	return r.db.WithContext(ctx).Create(tender).Error
}

// Update modifies an existing tender in the database.
func (r *tenderRepositoryGorm) Update(ctx context.Context, tender *entities.Tender) error {
	return r.db.WithContext(ctx).Save(tender).Error
}

// FindByID retrieves a tender by its ID.
func (r *tenderRepositoryGorm) FindByID(ctx context.Context, id int) (*entities.Tender, error) {
	var tender entities.Tender
	if err := r.db.WithContext(ctx).Preload("Versions").First(&tender, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tender not found")
		}
		return nil, err
	}
	return &tender, nil
}

// GetAll retrieves all tenders from the database.
func (r *tenderRepositoryGorm) GetAll(ctx context.Context) ([]*entities.Tender, error) {
	var tenders []*entities.Tender
	if err := r.db.WithContext(ctx).Preload("Versions").Find(&tenders).Error; err != nil {
		return nil, err
	}
	return tenders, nil
}

// CreateVersion adds a new version of a tender to the database.
func (r *tenderRepositoryGorm) CreateVersion(ctx context.Context, version *entities.TenderVersion) error {
	return r.db.WithContext(ctx).Create(version).Error
}

// FindVersionByNumber retrieves a specific version of a tender by tender ID and version number.
func (r *tenderRepositoryGorm) FindVersionByNumber(ctx context.Context, tenderID int, versionNumber int) (*entities.TenderVersion, error) {
	var version entities.TenderVersion
	if err := r.db.WithContext(ctx).
		Where("tender_id = ? AND version = ?", tenderID, versionNumber).
		First(&version).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("tender version not found")
		}
		return nil, err
	}
	return &version, nil
}

// FindLatestVersion retrieves the latest version of a tender.
func (r *tenderRepositoryGorm) FindLatestVersion(ctx context.Context, tenderID int) (*entities.TenderVersion, error) {
	var version entities.TenderVersion
	if err := r.db.WithContext(ctx).
		Where("tender_id = ?", tenderID).
		Order("version DESC").
		First(&version).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("latest tender version not found")
		}
		return nil, err
	}
	return &version, nil
}

// FindUserOrganizationResponsibility checks if a user is a responsible person for a given organization.
func (r *tenderRepositoryGorm) FindUserOrganizationResponsibility(ctx context.Context, userID, orgID int) (*entities.OrganizationResponsible, error) {
	var responsible entities.OrganizationResponsible
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		First(&responsible).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No responsibility found
		}
		return nil, err
	}
	return &responsible, nil
}
