package bid_repository

import (
	"avitoTest/data/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockBidRepository is a mock implementation of BidRepository.
type MockBidRepository struct {
	mock.Mock
}

// Create mocks the creation of a Bid.
func (m *MockBidRepository) Create(ctx context.Context, bid *entities.Bid) error {
	args := m.Called(ctx, bid)
	return args.Error(0)
}

// Update mocks the updating of a Bid.
func (m *MockBidRepository) Update(ctx context.Context, bid *entities.Bid) error {
	args := m.Called(ctx, bid)
	return args.Error(0)
}

// FindByID mocks finding a Bid by its ID.
func (m *MockBidRepository) FindByID(ctx context.Context, id int) (*entities.Bid, error) {
	args := m.Called(ctx, id)
	if bid, ok := args.Get(0).(*entities.Bid); ok {
		return bid, args.Error(1)
	}
	return nil, args.Error(1)
}

// FindByTenderID mocks finding Bids by Tender ID.
func (m *MockBidRepository) FindByTenderID(ctx context.Context, tenderID int) ([]*entities.Bid, error) {
	args := m.Called(ctx, tenderID)
	if bids, ok := args.Get(0).([]*entities.Bid); ok {
		return bids, args.Error(1)
	}
	return nil, args.Error(1)
}

// FindByCreatorID mocks finding Bids by the Creator ID.
func (m *MockBidRepository) FindByCreatorID(ctx context.Context, creatorID int) ([]*entities.Bid, error) {
	args := m.Called(ctx, creatorID)
	if bids, ok := args.Get(0).([]*entities.Bid); ok {
		return bids, args.Error(1)
	}
	return nil, args.Error(1)
}

// FindLatestVersion mocks finding the latest version of a Bid.
func (m *MockBidRepository) FindLatestVersion(ctx context.Context, bidID int) (*entities.BidVersion, error) {
	args := m.Called(ctx, bidID)
	if version, ok := args.Get(0).(*entities.BidVersion); ok {
		return version, args.Error(1)
	}
	return nil, args.Error(1)
}

// FindVersionByNumber mocks finding a specific version of a Bid by its version number.
func (m *MockBidRepository) FindVersionByNumber(ctx context.Context, bidID int, versionNumber int) (*entities.BidVersion, error) {
	args := m.Called(ctx, bidID, versionNumber)
	if version, ok := args.Get(0).(*entities.BidVersion); ok {
		return version, args.Error(1)
	}
	return nil, args.Error(1)
}

// CreateVersion mocks the creation of a new BidVersion.
func (m *MockBidRepository) CreateVersion(ctx context.Context, version *entities.BidVersion) error {
	args := m.Called(ctx, version)
	return args.Error(0)
}

// Delete mocks the deletion of a Bid.
func (m *MockBidRepository) Delete(ctx context.Context, bidID int) error {
	args := m.Called(ctx, bidID)
	return args.Error(0)
}
