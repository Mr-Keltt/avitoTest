// File: services/tender_service/tender_service.go

package tender_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/tender_repository"
	"avitoTest/services/tender_service/tender_models"
	"avitoTest/shared/constants"
	"avitoTest/shared/errors/tendert_erorrs"
	"avitoTest/shared/validators"
	"context"
	"errors"
	"time"
)

type TenderService interface {
	CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error)
	UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error)
	PublishTender(ctx context.Context, tenderID int) error
	CloseTender(ctx context.Context, tenderID int) error
	GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error)
	RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error)
	GetAllTenders(ctx context.Context, serviceTypeFilter string) ([]*tender_models.TenderModel, error)
}

type tenderService struct {
	repo tender_repository.TenderRepository
}

func NewTenderService(repo tender_repository.TenderRepository) TenderService {
	return &tenderService{repo: repo}
}

// CreateTender создает новый тендер и проверяет статус
func (s *tenderService) CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error) {
	if !s.isUserResponsibleForOrganization(ctx, tender.CreatorID, tender.OrganizationID) {
		return nil, tendert_erorrs.ErrUnauthorized
	}

	// Validate the ServiceType using predefined constants
	if tender.ServiceType != string(constants.ServiceTypeConstruction) &&
		tender.ServiceType != string(constants.ServiceTypeIT) &&
		tender.ServiceType != string(constants.ServiceTypeConsulting) {
		return nil, errors.New("invalid service type")
	}

	// Check for valid tender status
	if tender.Status != constants.TenderStatusCreated && tender.Status != constants.TenderStatusPublished {
		return nil, tendert_erorrs.ErrInvalidStatus
	}

	entity := &entities.Tender{
		OrganizationID: tender.OrganizationID,
		CreatorID:      tender.CreatorID,
		Status:         string(tender.Status),
		ServiceType:    tender.ServiceType,
		CreatedAt:      time.Now(),
	}

	// Create the tender
	if err := s.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	// Create the initial version of the tender
	version := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        tender.Name,
		Description: tender.Description,
		Version:     1,
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
		ServiceType:    entity.ServiceType,
		Status:         tender.Status,
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// UpdateTender updates an existing tender and creates a new version
func (s *tenderService) UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error) {
	entity, err := s.repo.FindByID(ctx, tender.ID)
	if err != nil {
		if errors.Is(err, tendert_erorrs.ErrTenderNotFound) {
			return nil, tendert_erorrs.ErrTenderNotFound
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
		ServiceType:    entity.ServiceType,
		Status:         constants.TenderStatus(entity.Status),
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// PublishTender обновляет статус тендера на PUBLISHED, не изменяя версию
func (s *tenderService) PublishTender(ctx context.Context, tenderID int) error {
	entity, err := s.repo.FindByID(ctx, tenderID)
	if err != nil {
		return err
	}

	//Update the status without changing the version
	entity.Status = string(constants.TenderStatusPublished)

	if err := s.repo.Update(ctx, entity); err != nil {
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

	// Find the latest version of the tender
	latestVersion, err := s.repo.FindLatestVersion(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	tenderVersion, err := s.repo.FindVersionByNumber(ctx, tenderID, version)
	if err != nil {
		return nil, tendert_erorrs.ErrTenderVersionNotFound
	}

	// Create a new version with rolled-back data
	newVersion := latestVersion.Version + 1
	rollbackVersion := &entities.TenderVersion{
		TenderID:    entity.ID,
		Name:        tenderVersion.Name,
		Description: tenderVersion.Description,
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
		ServiceType:    entity.ServiceType,
		Status:         constants.TenderStatus(entity.Status),
		CreatedAt:      entity.CreatedAt,
		Version:        rollbackVersion.Version,
	}, nil
}

// GetAllTenders retrieves all tenders
func (s *tenderService) GetAllTenders(ctx context.Context, serviceTypeFilter string) ([]*tender_models.TenderModel, error) {
	var entities []*entities.Tender
	var err error

	// Если есть фильтр по типу услуг, используем его
	if serviceTypeFilter != "" {
		if !validators.IsValidServiceType(serviceTypeFilter) {
			return nil, errors.New("invalid service type")
		}
		entities, err = s.repo.GetAllByServiceType(ctx, serviceTypeFilter)
	} else {
		entities, err = s.repo.GetAll(ctx)
	}

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
			Status:         constants.TenderStatus(entity.Status),
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
		if errors.Is(err, tendert_erorrs.ErrTenderNotFound) {
			return nil, tendert_erorrs.ErrTenderNotFound
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
		Status:         constants.TenderStatus(entity.Status),
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
