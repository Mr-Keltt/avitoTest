// File: services/tender_service/tender_service.go

package tender_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/tender_repository"
	"avitoTest/services/tender_service/tender_models"
	"context"
	"errors"
	"time"
)

var (
	ErrTenderNotFound        = errors.New("tender not found")
	ErrUnauthorized          = errors.New("user not authorized")
	ErrTenderVersionNotFound = errors.New("tender version not found")
)

type TenderService interface {
	CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error)
	UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error)
	PublishTender(ctx context.Context, tenderID int) error
	CloseTender(ctx context.Context, tenderID int) error
	GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error)
	RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error)
	GetAllTenders(ctx context.Context) ([]*tender_models.TenderModel, error)
}

type tenderService struct {
	repo tender_repository.TenderRepository
}

func NewTenderService(repo tender_repository.TenderRepository) TenderService {
	return &tenderService{repo: repo}
}

// CreateTender creates a new tender and the initial version
func (s *tenderService) CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error) {
	if !s.isUserResponsibleForOrganization(ctx, tender.CreatorID, tender.OrganizationID) {
		return nil, ErrUnauthorized
	}

	entity := &entities.Tender{
		OrganizationID: tender.OrganizationID,
		CreatorID:      tender.CreatorID,
		CreatedAt:      time.Now(),
	}

	if err := s.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	// Create initial version
	version := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		Version:     1, // Initial version
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateVersion(ctx, version); err != nil {
		return nil, err
	}

	return &tender_models.TenderModel{
		ID:             entity.ID,
		OrganizationID: entity.OrganizationID,
		Name:           version.Name,
		Description:    version.Description,
		ServiceType:    version.ServiceType,
		Status:         tender_models.TenderStatusCreated, // Set the initial status
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// UpdateTender updates an existing tender and creates a new version
func (s *tenderService) UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error) {
	entity, err := s.repo.FindByID(ctx, tender.ID)
	if err != nil {
		if errors.Is(err, tender_repository.ErrTenderNotFound) {
			return nil, ErrTenderNotFound
		}
		return nil, err
	}

	// Find the latest version of the tender
	latestVersion, err := s.repo.FindLatestVersion(ctx, tender.ID)
	if err != nil {
		return nil, err
	}

	// Increment version number
	newVersion := latestVersion.Version + 1

	// Create new version
	version := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        tender.Name,
		Description: tender.Description,
		ServiceType: tender.ServiceType,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateVersion(ctx, version); err != nil {
		return nil, err
	}

	return &tender_models.TenderModel{
		ID:             entity.ID,
		OrganizationID: entity.OrganizationID,
		Name:           version.Name,
		Description:    version.Description,
		ServiceType:    version.ServiceType,
		Status:         latestVersionStatus(newVersion), // Determine status based on new version
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// PublishTender sets the status of the tender to PUBLISHED
func (s *tenderService) PublishTender(ctx context.Context, tenderID int) error {
	entity, err := s.repo.FindByID(ctx, tenderID)
	if err != nil {
		return err
	}

	// Find the latest version of the tender
	latestVersion, err := s.repo.FindLatestVersion(ctx, tenderID)
	if err != nil {
		return err
	}

	// Update the status indirectly by creating a new version with the updated status
	newVersion := latestVersion.Version + 1
	version := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        latestVersion.Name,
		Description: latestVersion.Description,
		ServiceType: latestVersion.ServiceType,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateVersion(ctx, version); err != nil {
		return err
	}

	return nil
}

// CloseTender sets the status of the tender to CLOSED
func (s *tenderService) CloseTender(ctx context.Context, tenderID int) error {
	entity, err := s.repo.FindByID(ctx, tenderID)
	if err != nil {
		return err
	}

	// Find the latest version of the tender
	latestVersion, err := s.repo.FindLatestVersion(ctx, tenderID)
	if err != nil {
		return err
	}

	// Update the status indirectly by creating a new version with the updated status
	newVersion := latestVersion.Version + 1
	version := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        latestVersion.Name,
		Description: latestVersion.Description,
		ServiceType: latestVersion.ServiceType,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateVersion(ctx, version); err != nil {
		return err
	}

	return nil
}

// RollbackTenderVersion rolls back the tender to a specific version
func (s *tenderService) RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error) {
	entity, err := s.repo.FindByID(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	tenderVersion, err := s.repo.FindVersionByNumber(ctx, tenderID, version)
	if err != nil {
		return nil, ErrTenderVersionNotFound
	}

	// Create a new version with rolled-back data
	newVersion := tenderVersion.Version + 1
	rollbackVersion := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        tenderVersion.Name,
		Description: tenderVersion.Description,
		ServiceType: tenderVersion.ServiceType,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateVersion(ctx, rollbackVersion); err != nil {
		return nil, err
	}

	return &tender_models.TenderModel{
		ID:             entity.ID,
		OrganizationID: entity.OrganizationID,
		Name:           rollbackVersion.Name,
		Description:    rollbackVersion.Description,
		ServiceType:    rollbackVersion.ServiceType,
		Status:         latestVersionStatus(newVersion), // Determine status based on new version
		CreatedAt:      entity.CreatedAt,
		Version:        rollbackVersion.Version,
	}, nil
}

// GetAllTenders retrieves all tenders
func (s *tenderService) GetAllTenders(ctx context.Context) ([]*tender_models.TenderModel, error) {
	entities, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var tenders []*tender_models.TenderModel
	for _, entity := range entities {
		// Find the latest version of each tender
		latestVersion, err := s.repo.FindLatestVersion(ctx, entity.ID)
		if err != nil {
			return nil, err
		}

		tenderModel := &tender_models.TenderModel{
			ID:             entity.ID,
			OrganizationID: entity.OrganizationID,
			Name:           latestVersion.Name,
			Description:    latestVersion.Description,
			ServiceType:    latestVersion.ServiceType,
			Status:         latestVersionStatus(latestVersion.Version), // Determine status
			CreatedAt:      entity.CreatedAt,
			Version:        latestVersion.Version,
		}
		tenders = append(tenders, tenderModel)
	}

	return tenders, nil
}

// GetTenderByID retrieves a specific tender by ID
func (s *tenderService) GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error) {
	entity, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, tender_repository.ErrTenderNotFound) {
			return nil, ErrTenderNotFound
		}
		return nil, err
	}

	// Find the latest version of the tender
	latestVersion, err := s.repo.FindLatestVersion(ctx, id)
	if err != nil {
		return nil, err
	}

	return &tender_models.TenderModel{
		ID:             entity.ID,
		OrganizationID: entity.OrganizationID,
		Name:           latestVersion.Name,
		Description:    latestVersion.Description,
		ServiceType:    latestVersion.ServiceType,
		Status:         latestVersionStatus(latestVersion.Version), // Determine status
		CreatedAt:      entity.CreatedAt,
		Version:        latestVersion.Version,
	}, nil
}

// isUserResponsibleForOrganization checks if the user is responsible for the organization
func (s *tenderService) isUserResponsibleForOrganization(ctx context.Context, userID, orgID int) bool {
	responsible, err := s.repo.FindUserOrganizationResponsibility(ctx, userID, orgID)
	if err != nil {
		return false
	}
	return responsible != nil
}

// latestVersionStatus determines the status based on the version number or other criteria.
func latestVersionStatus(version int) tender_models.TenderStatus {
	// Logic to determine status based on version
	if version == 1 {
		return tender_models.TenderStatusCreated
	}
	// Example: Return different statuses depending on other conditions
	return tender_models.TenderStatusPublished
}
