package bid_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/bid_repository"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/bid_service/bid_models"
	"avitoTest/shared"
	"context"
	"errors"
	"time"
)

type BidService interface {
	CreateBid(ctx context.Context, bid bid_models.BidCreateModel) (*bid_models.BidModel, error)
	UpdateBid(ctx context.Context, bid bid_models.BidUpdateModel) (*bid_models.BidModel, error)
	GetBidByID(ctx context.Context, bidID int) (*bid_models.BidModel, error)
	GetBidsByTenderID(ctx context.Context, tenderID int) ([]*bid_models.BidModel, error)
	GetBidsByUserID(ctx context.Context, userID int) ([]*bid_models.BidModel, error)
	ApproveBid(ctx context.Context, bidID, approverID int) error
	RejectBid(ctx context.Context, bidID, rejecterID int) error
	RollbackBidVersion(ctx context.Context, bidID int, version int) (*bid_models.BidModel, error)
	DeleteBid(ctx context.Context, bidID int) error
}

type bidService struct {
	bidRepo bid_repository.BidRepository
	orgRepo organization_repository.OrganizationRepository
}

func NewBidService(bidRepo bid_repository.BidRepository, orgRepo organization_repository.OrganizationRepository) BidService {
	return &bidService{
		bidRepo: bidRepo,
		orgRepo: orgRepo,
	}
}

// CreateBid creates a new bid and its initial version
func (s *bidService) CreateBid(ctx context.Context, bid bid_models.BidCreateModel) (*bid_models.BidModel, error) {
	// Validate required fields (example validation, customize as needed)
	if bid.Name == "" {
		return nil, errors.New("name is required")
	}
	if bid.TenderID <= 0 {
		return nil, errors.New("invalid tender ID")
	}
	if bid.OrganizationID <= 0 {
		return nil, errors.New("invalid organization ID")
	}

	entity := &entities.Bid{
		TenderID:       bid.TenderID,
		OrganizationID: bid.OrganizationID,
		CreatorID:      bid.CreatorID,
		Status:         "CREATED",
		ApprovalCount:  0,
		CreatedAt:      time.Now(),
	}

	if err := s.bidRepo.Create(ctx, entity); err != nil {
		return nil, err
	}

	version := &entities.BidVersion{
		BidID:       entity.ID,
		Name:        bid.Name,
		Description: bid.Description,
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	if err := s.bidRepo.CreateVersion(ctx, version); err != nil {
		return nil, err
	}

	return &bid_models.BidModel{
		ID:             entity.ID,
		Name:           version.Name,
		Description:    version.Description,
		TenderID:       entity.TenderID,
		OrganizationID: entity.OrganizationID,
		CreatorID:      entity.CreatorID,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// UpdateBid updates the bid and increments its version
func (s *bidService) UpdateBid(ctx context.Context, bid bid_models.BidUpdateModel) (*bid_models.BidModel, error) {
	entity, err := s.bidRepo.FindByID(ctx, bid.ID)
	if err != nil {
		return nil, err
	}

	// Update the bid details
	entity.Status = "UPDATED"
	if err := s.bidRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	// Find the latest version and increment it
	latestVersion, err := s.bidRepo.FindLatestVersion(ctx, entity.ID)
	if err != nil {
		return nil, err
	}

	newVersion := latestVersion.Version + 1
	version := &entities.BidVersion{
		BidID:       entity.ID,
		Name:        bid.Name,
		Description: bid.Description,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.bidRepo.CreateVersion(ctx, version); err != nil {
		return nil, err
	}

	return &bid_models.BidModel{
		ID:             entity.ID,
		Name:           version.Name,
		Description:    version.Description,
		TenderID:       entity.TenderID,
		OrganizationID: entity.OrganizationID,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt,
		Version:        version.Version,
	}, nil
}

// GetBidByID retrieves the bid and its latest version by ID
func (s *bidService) GetBidByID(ctx context.Context, bidID int) (*bid_models.BidModel, error) {
	entity, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return nil, err
	}

	latestVersion, err := s.bidRepo.FindLatestVersion(ctx, bidID)
	if err != nil {
		return nil, err
	}

	return &bid_models.BidModel{
		ID:             entity.ID,
		Name:           latestVersion.Name,
		Description:    latestVersion.Description,
		TenderID:       entity.TenderID,
		OrganizationID: entity.OrganizationID,
		CreatorID:      entity.CreatorID,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt,
		Version:        latestVersion.Version,
	}, nil
}

// GetBidsByTenderID retrieves all bids for a specific tender
func (s *bidService) GetBidsByTenderID(ctx context.Context, tenderID int) ([]*bid_models.BidModel, error) {
	entities, err := s.bidRepo.FindByTenderID(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	var bids []*bid_models.BidModel
	for _, entity := range entities {
		latestVersion, err := s.bidRepo.FindLatestVersion(ctx, entity.ID)
		if err != nil {
			return nil, err
		}

		bidModel := &bid_models.BidModel{
			ID:             entity.ID,
			Name:           latestVersion.Name,
			Description:    latestVersion.Description,
			Status:         entity.Status,
			TenderID:       entity.TenderID,
			OrganizationID: entity.OrganizationID,
			CreatedAt:      entity.CreatedAt,
			Version:        latestVersion.Version,
		}
		bids = append(bids, bidModel)
	}
	return bids, nil
}

// GetBidsByUserID retrieves all bids created by a specific user
func (s *bidService) GetBidsByUserID(ctx context.Context, userID int) ([]*bid_models.BidModel, error) {
	entities, err := s.bidRepo.FindByCreatorID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var bids []*bid_models.BidModel
	for _, entity := range entities {
		latestVersion, err := s.bidRepo.FindLatestVersion(ctx, entity.ID)
		if err != nil {
			return nil, err
		}

		bidModel := &bid_models.BidModel{
			ID:             entity.ID,
			Name:           latestVersion.Name,
			Description:    latestVersion.Description,
			Status:         entity.Status,
			TenderID:       entity.TenderID,
			OrganizationID: entity.OrganizationID,
			CreatedAt:      entity.CreatedAt,
			Version:        latestVersion.Version,
		}
		bids = append(bids, bidModel)
	}
	return bids, nil
}

func (s *bidService) ApproveBid(ctx context.Context, bidID, approverID int) error {
	// Найдем предложение
	bid, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return err
	}

	// Проверим, что пользователь является ответственным за организацию
	isResponsible, err := s.isUserResponsibleForOrganization(ctx, bid.OrganizationID, approverID)
	if err != nil {
		return err
	}
	if !isResponsible {
		return errors.New("user is not responsible for the organization")
	}

	// Определим кворум
	responsibles, err := s.orgRepo.GetResponsibles(ctx, bid.OrganizationID)
	if err != nil {
		return err
	}

	quorum := min(3, len(responsibles))
	bid.ApprovalCount++

	// Проверим, если достигнут кворум, меняем статус
	if bid.ApprovalCount >= quorum {
		bid.Status = "APPROVED"
		bid.Tender.Status = "CLOSED"
	} else {
		shared.Logger.Infof("Approval count: %d, Quorum: %d\n", bid.ApprovalCount, quorum)
	}

	// Обновим статус предложения
	if err := s.bidRepo.Update(ctx, bid); err != nil {
		return err
	}

	return nil
}

func (s *bidService) RejectBid(ctx context.Context, bidID, rejecterID int) error {
	//Let's find an offer
	bid, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return err
	}

	//Let's check that the user is responsible for the organization
	isResponsible, err := s.isUserResponsibleForOrganization(ctx, bid.OrganizationID, rejecterID)
	if err != nil {
		return err
	}
	if !isResponsible {
		return errors.New("user is not responsible for the organization")
	}

	//If there is at least one deviation, the status changes to "REJECTED"
	bid.Status = "REJECTED"
	bid.Tender.Status = "CLOSED"

	// Обновим статус предложения
	if err := s.bidRepo.Update(ctx, bid); err != nil {
		return err
	}

	return nil
}

// RollbackBidVersion rolls back the bid to a specific version
func (s *bidService) RollbackBidVersion(ctx context.Context, bidID int, versionNumber int) (*bid_models.BidModel, error) {
	entity, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return nil, err
	}

	tenderVersion, err := s.bidRepo.FindVersionByNumber(ctx, bidID, versionNumber)
	if err != nil {
		return nil, errors.New("bid version not found")
	}

	// Create a new version with the rolled-back details
	newVersion := tenderVersion.Version + 1
	rollbackVersion := &entities.BidVersion{
		BidID:       entity.ID,
		Name:        tenderVersion.Name,
		Description: tenderVersion.Description,
		Version:     newVersion,
		UpdatedAt:   time.Now(),
	}

	if err := s.bidRepo.CreateVersion(ctx, rollbackVersion); err != nil {
		return nil, err
	}

	return &bid_models.BidModel{
		ID:             entity.ID,
		Name:           rollbackVersion.Name,
		Description:    rollbackVersion.Description,
		Status:         entity.Status,
		TenderID:       entity.TenderID,
		OrganizationID: entity.OrganizationID,
		CreatedAt:      entity.CreatedAt,
		Version:        rollbackVersion.Version,
	}, nil
}

// DeleteBid removes a bid by its ID
func (s *bidService) DeleteBid(ctx context.Context, bidID int) error {
	_, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return err
	}

	return s.bidRepo.Delete(ctx, bidID)
}

// isUserResponsibleForOrganization checks whether the user is responsible for the organization.
func (s *bidService) isUserResponsibleForOrganization(ctx context.Context, organizationID int, userID int) (bool, error) {
	responsibles, err := s.orgRepo.GetResponsibles(ctx, organizationID)
	if err != nil {
		return false, err
	}

	for _, responsible := range responsibles {
		if responsible.ID == userID {
			return true, nil
		}
	}
	return false, nil
}

// min - auxiliary function for determining the minimum value
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
