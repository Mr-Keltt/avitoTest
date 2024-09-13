package bid_service

import (
	"avitoTest/services/bid_service/bid_models"
	"context"

	"github.com/stretchr/testify/mock"
)

// Mock the BidService
type MockBidService struct {
	mock.Mock
}

func (m *MockBidService) CreateBid(ctx context.Context, bid bid_models.BidCreateModel) (*bid_models.BidModel, error) {
	args := m.Called(ctx, bid)
	return args.Get(0).(*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) UpdateBid(ctx context.Context, bid bid_models.BidUpdateModel) (*bid_models.BidModel, error) {
	args := m.Called(ctx, bid)
	return args.Get(0).(*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) GetBidByID(ctx context.Context, bidID int) (*bid_models.BidModel, error) {
	args := m.Called(ctx, bidID)
	return args.Get(0).(*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) GetBidsByTenderID(ctx context.Context, tenderID int) ([]*bid_models.BidModel, error) {
	args := m.Called(ctx, tenderID)
	return args.Get(0).([]*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) GetBidsByUserID(ctx context.Context, userID int) ([]*bid_models.BidModel, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) GetBidsByUsername(ctx context.Context, username string) ([]*bid_models.BidModel, error) {
	args := m.Called(ctx, username)
	return args.Get(0).([]*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) ApproveBid(ctx context.Context, bidID, approverID int) error {
	args := m.Called(ctx, bidID, approverID)
	return args.Error(0)
}

func (m *MockBidService) RejectBid(ctx context.Context, bidID, rejecterID int) error {
	args := m.Called(ctx, bidID, rejecterID)
	return args.Error(0)
}

func (m *MockBidService) RollbackBidVersion(ctx context.Context, bidID int, version int) (*bid_models.BidModel, error) {
	args := m.Called(ctx, bidID, version)
	return args.Get(0).(*bid_models.BidModel), args.Error(1)
}

func (m *MockBidService) DeleteBid(ctx context.Context, bidID int) error {
	args := m.Called(ctx, bidID)
	return args.Error(0)
}
