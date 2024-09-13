package bid_repository

import (
	"avitoTest/data/entities"
	"context"

	"gorm.io/gorm"
)

type bidRepositoryGorm struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) BidRepository {
	return &bidRepositoryGorm{db: db}
}

func (r *bidRepositoryGorm) Create(ctx context.Context, bid *entities.Bid) error {
	return r.db.WithContext(ctx).Create(bid).Error
}

func (r *bidRepositoryGorm) Update(ctx context.Context, bid *entities.Bid) error {
	return r.db.WithContext(ctx).Save(bid).Error
}

func (r *bidRepositoryGorm) FindByID(ctx context.Context, id int) (*entities.Bid, error) {
	var bid entities.Bid
	if err := r.db.WithContext(ctx).First(&bid, id).Error; err != nil {
		return nil, err
	}
	return &bid, nil
}

func (r *bidRepositoryGorm) FindByTenderID(ctx context.Context, tenderID int) ([]*entities.Bid, error) {
	var bids []*entities.Bid
	if err := r.db.WithContext(ctx).Where("tender_id = ?", tenderID).Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}

func (r *bidRepositoryGorm) FindByCreatorID(ctx context.Context, creatorID int) ([]*entities.Bid, error) {
	var bids []*entities.Bid
	if err := r.db.WithContext(ctx).Where("creator_id = ?", creatorID).Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}

// FindByUsername finds all bids created by a user with the given username.
func (r *bidRepositoryGorm) FindByUsername(ctx context.Context, username string) ([]*entities.Bid, error) {
	// Step 1: Find the user by username
	var user entities.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Step 2: Fetch bids created by the user
	var bids []*entities.Bid
	if err := r.db.WithContext(ctx).Where("creator_id = ?", user.ID).Find(&bids).Error; err != nil {
		return nil, err
	}

	return bids, nil
}

func (r *bidRepositoryGorm) FindLatestVersion(ctx context.Context, bidID int) (*entities.BidVersion, error) {
	var version entities.BidVersion
	if err := r.db.WithContext(ctx).Where("bid_id = ?", bidID).Order("version DESC").First(&version).Error; err != nil {
		return nil, err
	}
	return &version, nil
}

func (r *bidRepositoryGorm) FindVersionByNumber(ctx context.Context, bidID int, versionNumber int) (*entities.BidVersion, error) {
	var version entities.BidVersion
	if err := r.db.WithContext(ctx).Where("bid_id = ? AND version = ?", bidID, versionNumber).First(&version).Error; err != nil {
		return nil, err
	}
	return &version, nil
}

func (r *bidRepositoryGorm) CreateVersion(ctx context.Context, version *entities.BidVersion) error {
	return r.db.WithContext(ctx).Create(version).Error
}

func (r *bidRepositoryGorm) Delete(ctx context.Context, bidID int) error {
	return r.db.WithContext(ctx).Delete(&entities.Bid{}, bidID).Error
}
