package service_test

import (
	"avitoTest/data/entities"
	"avitoTest/services/bid_service"
	"avitoTest/services/bid_service/bid_models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func setupMocks() (*MockBidRepository, bid_service.BidService) {
	mockBidRepo := new(MockBidRepository)
	service := bid_service.NewBidService(mockBidRepo)
	return mockBidRepo, service
}

func TestCreateBid_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	bidCreate := bid_models.BidCreateModel{
		Name:           "Bid 1",
		Description:    "Test Description",
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
	}

	expectedEntity := &entities.Bid{
		ID:             1,
		TenderID:       bidCreate.TenderID,
		OrganizationID: bidCreate.OrganizationID,
		CreatorID:      bidCreate.CreatorID,
		Status:         bidCreate.Status,
		CreatedAt:      time.Now(),
	}

	// Mock the Create method for Bid
	mockBidRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil).Run(func(args mock.Arguments) {
		bid := args.Get(1).(*entities.Bid)
		bid.ID = expectedEntity.ID
		bid.CreatedAt = expectedEntity.CreatedAt
	})

	// Mock the CreateVersion method for BidVersion
	expectedVersion := &entities.BidVersion{
		BidID:       expectedEntity.ID,
		Name:        bidCreate.Name,
		Description: bidCreate.Description,
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockBidRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.BidVersion")).Return(nil).Run(func(args mock.Arguments) {
		version := args.Get(1).(*entities.BidVersion)
		version.BidID = expectedVersion.BidID
		version.Name = expectedVersion.Name
		version.Description = expectedVersion.Description
		version.Version = expectedVersion.Version
		version.UpdatedAt = expectedVersion.UpdatedAt
	})

	// Execute the CreateBid method in the service
	result, err := service.CreateBid(context.Background(), bidCreate)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, bidCreate.Name, result.Name)
	assert.Equal(t, bidCreate.Description, result.Description)
	mockBidRepo.AssertExpectations(t)
}

func TestCreateBid_ValidationFail(t *testing.T) {
	// Setup mock repository and service
	mockBidRepo, service := setupMocks()

	// Prepare invalid BidCreateModel (missing required fields, e.g., Name is empty)
	bidCreate := bid_models.BidCreateModel{
		Name:           "", // Invalid: name is required
		Description:    "Test Description",
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
	}

	// Act: Call CreateBid and expect validation to fail
	_, err := service.CreateBid(context.Background(), bidCreate)

	// Assert: Ensure that error is returned and it's a validation error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name is required") // Validation message should be correct

	// No need to mock or assert repository methods because they shouldn't be called on validation failure
	mockBidRepo.AssertNotCalled(t, "Create")
	mockBidRepo.AssertNotCalled(t, "CreateVersion")
}

func TestUpdateBid_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	bidUpdate := bid_models.BidUpdateModel{
		ID:          1,
		Name:        "Updated Bid",
		Description: "Updated Description",
	}

	existingBid := &entities.Bid{
		ID:             bidUpdate.ID,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.BidVersion{
		BidID:     bidUpdate.ID,
		Version:   1,
		UpdatedAt: time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, bidUpdate.ID).Return(existingBid, nil)
	mockBidRepo.On("FindLatestVersion", mock.Anything, bidUpdate.ID).Return(latestVersion, nil)
	mockBidRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil)
	mockBidRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.BidVersion")).Return(nil)

	result, err := service.UpdateBid(context.Background(), bidUpdate)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Bid", result.Name)
	mockBidRepo.AssertExpectations(t)
}

func TestUpdateBid_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	bidUpdate := bid_models.BidUpdateModel{
		ID:          1,
		Name:        "Updated Bid",
		Description: "Updated Description",
	}

	mockBidRepo.On("FindByID", mock.Anything, bidUpdate.ID).Return(nil, errors.New("bid not found"))

	_, err := service.UpdateBid(context.Background(), bidUpdate)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidByID_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	expectedBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.BidVersion{
		BidID:     expectedBid.ID,
		Version:   1,
		Name:      "Test Bid",
		UpdatedAt: time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(expectedBid, nil)
	mockBidRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)

	result, err := service.GetBidByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, latestVersion.Name, result.Name)
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidByID_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	_, err := service.GetBidByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidsByTenderID_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	expectedBids := []*entities.Bid{
		{
			ID:             1,
			TenderID:       1,
			OrganizationID: 1,
			CreatorID:      1,
			Status:         "CREATED",
			CreatedAt:      time.Now(),
		},
	}

	latestVersion := &entities.BidVersion{
		BidID:     1,
		Version:   1,
		Name:      "Test Bid",
		UpdatedAt: time.Now(),
	}

	mockBidRepo.On("FindByTenderID", mock.Anything, 1).Return(expectedBids, nil)
	mockBidRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)

	result, err := service.GetBidsByTenderID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, latestVersion.Name, result[0].Name)
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidsByTenderID_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	mockBidRepo.On("FindByTenderID", mock.Anything, 1).Return(nil, errors.New("no bids found"))

	_, err := service.GetBidsByTenderID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "no bids found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestApproveBid_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		ApprovalCount:  2,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)
	mockBidRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil)

	err := service.ApproveBid(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, 3, existingBid.ApprovalCount)
	mockBidRepo.AssertExpectations(t)
}

func TestApproveBid_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	err := service.ApproveBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestRejectBid_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)
	mockBidRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil)

	err := service.RejectBid(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, "REJECTED", existingBid.Status)
	mockBidRepo.AssertExpectations(t)
}

func TestRejectBid_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	err := service.RejectBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestRollbackBidVersion_Success(t *testing.T) {
	mockBidRepo, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	rollbackVersion := &entities.BidVersion{
		BidID:       1,
		Name:        "Version 1",
		Description: "Description",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)
	mockBidRepo.On("FindVersionByNumber", mock.Anything, 1, 1).Return(rollbackVersion, nil)
	mockBidRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.BidVersion")).Return(nil)

	result, err := service.RollbackBidVersion(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, "Version 1", result.Name)
	mockBidRepo.AssertExpectations(t)
}

func TestRollbackBidVersion_NotFound(t *testing.T) {
	mockBidRepo, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	_, err := service.RollbackBidVersion(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}
