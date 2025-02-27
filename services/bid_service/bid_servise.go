package bid_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/bid_repository"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/data/repositories/tender_repository"
	"avitoTest/data/repositories/user_repository"
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
	GetBidsByUsername(ctx context.Context, username string) ([]*bid_models.BidModel, error)
	ApproveBid(ctx context.Context, bidID, approverID int) error
	RejectBid(ctx context.Context, bidID, rejecterID int) error
	RollbackBidVersion(ctx context.Context, bidID int, version int) (*bid_models.BidModel, error)
	DeleteBid(ctx context.Context, bidID int) error
}

type bidService struct {
	bidRepo    bid_repository.BidRepository
	orgRepo    organization_repository.OrganizationRepository
	userRepo   user_repository.UserRepository
	tenderRepo tender_repository.TenderRepository
}

func NewBidService(
	bidRepo bid_repository.BidRepository,
	orgRepo organization_repository.OrganizationRepository,
	userRepo user_repository.UserRepository,
	tenderRepo tender_repository.TenderRepository) BidService {
	return &bidService{
		bidRepo:    bidRepo,
		orgRepo:    orgRepo,
		userRepo:   userRepo,
		tenderRepo: tenderRepo,
	}
}

// CreateBid creates a new bid and its initial version
func (s *bidService) CreateBid(ctx context.Context, bid bid_models.BidCreateModel) (*bid_models.BidModel, error) {
	// Validate required fields (example validation, customize as needed)
	if bid.Name == "" {
		return nil, errors.New("name is required")
	}

	// Log the TenderID for additional debugging
	shared.Logger.Infof("Creating bid with TenderID: %d", bid.TenderID)

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

func (s *bidService) GetBidsByUsername(ctx context.Context, username string) ([]*bid_models.BidModel, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	bids, err := s.bidRepo.FindByCreatorID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	var bidModels []*bid_models.BidModel
	for _, bid := range bids {
		version, err := s.bidRepo.FindLatestVersion(ctx, bid.ID)
		if err != nil {
			return nil, err
		}

		bidModel := &bid_models.BidModel{
			ID:             bid.ID,
			TenderID:       bid.TenderID,
			OrganizationID: bid.OrganizationID,
			CreatorID:      bid.CreatorID,
			Status:         bid.Status,
			CreatedAt:      bid.CreatedAt,
			Name:           version.Name,
			Description:    version.Description,
		}

		bidModels = append(bidModels, bidModel)
	}

	return bidModels, nil
}

func (s *bidService) ApproveBid(ctx context.Context, bidID, approverID int) error {
	// Fetch the bid by ID
	bid, err := s.bidRepo.FindByID(ctx, bidID)
	if err != nil {
		return err
	}

	// Log the organization ID to see if it's 0
	shared.Logger.Infof("Bid organization ID: %d", bid.OrganizationID)

	// Validate that the organization exists
	if bid.OrganizationID <= 0 {
		return errors.New("invalid organization ID")
	}

	if err := s.validateOrganizationExists(ctx, bid.OrganizationID); err != nil {
		return err
	}

	// Check if the user is responsible for this organization
	isResponsible, err := s.isUserResponsibleForOrganization(ctx, bid.OrganizationID, approverID)
	if err != nil {
		return err
	}
	if !isResponsible {
		return errors.New("user is not responsible for the organization")
	}

	// Get the responsibles for this organization
	responsibles, err := s.orgRepo.GetResponsibles(ctx, bid.OrganizationID)
	if err != nil {
		return err
	}

	if bid.Status == "REJECTED" {
		return errors.New("bid already REJECTED")
	}

	// Calculate the quorum and increment approval count
	quorum := min(3, len(responsibles))
	bid.ApprovalCount++

	// Update bid status if quorum is met
	if bid.ApprovalCount >= quorum {
		bid.Status = "APPROVED"

		// Call CloseTender method instead of manually setting the tender's status
		err := s.tenderRepo.CloseTender(ctx, bid.TenderID)

		if err != nil {
			return err
		}
	} else {
		shared.Logger.Infof("Approval count: %d, Quorum: %d\n", bid.ApprovalCount, quorum)
	}

	// Update the bid in the repository
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

// validateTenderExists checks whether a tender exists before updating its status.
func (s *bidService) validateOrganizationExists(ctx context.Context, organizationID int) error {
	organization, _ := s.bidRepo.FindByID(ctx, organizationID)
	if organization == nil {
		return errors.New("organization not found")
	}
	return nil
}

// validateTenderExists checks whether a tender exists before updating its status.
func (s *bidService) validateTenderExists(ctx context.Context, tenderID int) error {
	tender, _ := s.tenderRepo.FindByID(ctx, tenderID)
	if tender == nil {
		return errors.New("tender not found")
	}
	return nil
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
