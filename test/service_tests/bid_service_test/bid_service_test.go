package bid_service_test

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/bid_repository"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/data/repositories/tender_repository"
	"avitoTest/data/repositories/user_repository"
	"avitoTest/services/bid_service"
	"avitoTest/services/bid_service/bid_models"
	"avitoTest/services/user_service/user_models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupMocks() (*bid_repository.MockBidRepository, *organization_repository.MockOrganizationRepository, *user_repository.MockUserRepository, *tender_repository.MockTenderRepository, bid_service.BidService) {
	mockBidRepo := new(bid_repository.MockBidRepository)
	mockOrgRepo := new(organization_repository.MockOrganizationRepository)
	mockUserRepo := new(user_repository.MockUserRepository)
	mockTenderRepo := new(tender_repository.MockTenderRepository)
	service := bid_service.NewBidService(mockBidRepo, mockOrgRepo, mockUserRepo, mockTenderRepo)
	return mockBidRepo, mockOrgRepo, mockUserRepo, mockTenderRepo, service
}

func TestCreateBid_Success(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

	bidCreate := bid_models.BidCreateModel{
		Name:           "Bid 1",
		Description:    "Test Description",
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
	}

	expectedEntity := &entities.Bid{
		ID:             1,
		TenderID:       bidCreate.TenderID,
		OrganizationID: bidCreate.OrganizationID,
		CreatorID:      bidCreate.CreatorID,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil).Run(func(args mock.Arguments) {
		bid := args.Get(1).(*entities.Bid)
		bid.ID = expectedEntity.ID
		bid.CreatedAt = expectedEntity.CreatedAt
	})

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

	result, err := service.CreateBid(context.Background(), bidCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, bidCreate.Name, result.Name)
	assert.Equal(t, bidCreate.Description, result.Description)
	mockBidRepo.AssertExpectations(t)
}

func TestCreateBid_ValidationFail(t *testing.T) {
	// Setup mock repository and service
	mockBidRepo, _, _, _, service := setupMocks()

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
	mockBidRepo, _, _, _, service := setupMocks()

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
	mockBidRepo, _, _, _, service := setupMocks()

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
	mockBidRepo, _, _, _, service := setupMocks()

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
	mockBidRepo, _, _, _, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	_, err := service.GetBidByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidsByTenderID_Success(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

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
	mockBidRepo, _, _, _, service := setupMocks()

	mockBidRepo.On("FindByTenderID", mock.Anything, 1).Return(nil, errors.New("no bids found"))

	_, err := service.GetBidsByTenderID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "no bids found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestGetBidsByUsername_Success(t *testing.T) {
	mockBidRepo, _, mockUserRepo, _, service := setupMocks()

	expectedUser := &entities.User{
		ID:       1,
		Username: "testuser",
	}
	mockUserRepo.On("FindByUsername", mock.Anything, "testuser").Return(expectedUser, nil)

	expectedBids := []*entities.Bid{
		{
			ID:             1,
			TenderID:       1,
			OrganizationID: 1,
			CreatorID:      expectedUser.ID,
			Status:         "CREATED",
			CreatedAt:      time.Now(),
		},
		{
			ID:             2,
			TenderID:       2,
			OrganizationID: 2,
			CreatorID:      expectedUser.ID,
			Status:         "PENDING",
			CreatedAt:      time.Now(),
		},
	}
	mockBidRepo.On("FindByCreatorID", mock.Anything, expectedUser.ID).Return(expectedBids, nil)

	latestVersion1 := &entities.BidVersion{
		BidID:       1,
		Version:     1,
		Name:        "Test Bid 1",
		Description: "Test Description 1",
		UpdatedAt:   time.Now(),
	}
	latestVersion2 := &entities.BidVersion{
		BidID:       2,
		Version:     1,
		Name:        "Test Bid 2",
		Description: "Test Description 2",
		UpdatedAt:   time.Now(),
	}
	mockBidRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion1, nil)
	mockBidRepo.On("FindLatestVersion", mock.Anything, 2).Return(latestVersion2, nil)

	result, err := service.GetBidsByUsername(context.Background(), "testuser")

	assert.NoError(t, err)
	assert.Len(t, result, 2)

	t.Logf("Result[0]: %+v", result[0])
	t.Logf("Result[1]: %+v", result[1])

	assert.Equal(t, latestVersion1.Name, result[0].Name)
	assert.Equal(t, latestVersion1.Description, result[0].Description)

	assert.Equal(t, latestVersion2.Name, result[1].Name)
	assert.Equal(t, latestVersion2.Description, result[1].Description)

	mockBidRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestGetBidsByUsername_UserNotFound(t *testing.T) {
	mockBidRepo, _, mockUserRepo, _, service := setupMocks()

	mockUserRepo.On("FindByUsername", mock.Anything, "unknownuser").Return(nil, errors.New("user not found"))

	_, err := service.GetBidsByUsername(context.Background(), "unknownuser")

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockBidRepo.AssertNotCalled(t, "FindByCreatorID")
	mockUserRepo.AssertExpectations(t)
}

func TestGetBidsByUsername_NoBidsFound(t *testing.T) {
	mockBidRepo, _, mockUserRepo, _, service := setupMocks()

	expectedUser := &entities.User{
		ID:       1,
		Username: "testuser",
	}
	mockUserRepo.On("FindByUsername", mock.Anything, "testuser").Return(expectedUser, nil)

	mockBidRepo.On("FindByCreatorID", mock.Anything, expectedUser.ID).Return(nil, errors.New("no bids found"))

	_, err := service.GetBidsByUsername(context.Background(), "testuser")

	assert.Error(t, err)
	assert.Equal(t, "no bids found", err.Error())
	mockBidRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// func TestApproveBid_Success(t *testing.T) {
// 	mockBidRepo, mockOrgRepo, _, _, service := setupMocks()

// 	existingBid := &entities.Bid{
// 		ID:             1,
// 		TenderID:       1,
// 		OrganizationID: 1,
// 		ApprovalCount:  2, // Start with 2 approvals.
// 		Status:         "CREATED",
// 		CreatedAt:      time.Now(),
// 	}

// 	// Mock returning the existing bid
// 	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)

// 	// Mock returning true for user responsibility check
// 	mockOrgRepo.On("GetResponsibles", mock.Anything, 1).Return([]*user_models.UserModel{
// 		{ID: 1, Username: "user1"}, // Approver
// 		{ID: 2, Username: "user2"},
// 		{ID: 3, Username: "user3"},
// 	}, nil)

// 	// Mock returning true for isUserResponsibleForOrganization
// 	mockBidRepo.On("isUserResponsibleForOrganization", mock.Anything, 1, 1).Return(true, nil)

// 	// Mock updating the bid after approval
// 	mockBidRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil).Run(func(args mock.Arguments) {
// 		bid := args.Get(1).(*entities.Bid)
// 		bid.Status = "APPROVED"
// 		bid.ApprovalCount = 3
// 	})

// 	// Call the service method
// 	err := service.ApproveBid(context.Background(), 1, 1)

// 	// Verify no errors and the expected updates
// 	assert.NoError(t, err)
// 	assert.Equal(t, 3, existingBid.ApprovalCount)
// 	assert.Equal(t, "APPROVED", existingBid.Status)
// 	mockBidRepo.AssertExpectations(t)
// 	mockOrgRepo.AssertExpectations(t)
// }

func TestApproveBid_NotFound(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	err := service.ApproveBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestApproveBid_UserNotResponsible(t *testing.T) {
	mockBidRepo, mockOrgRepo, _, _, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)

	mockOrgRepo.On("GetResponsibles", mock.Anything, 1).Return([]*user_models.UserModel{}, nil)

	err := service.ApproveBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "user is not responsible for the organization", err.Error())
	mockBidRepo.AssertExpectations(t)
	mockOrgRepo.AssertExpectations(t)
}

// func TestRejectBid_Success(t *testing.T) {
// 	mockBidRepo, mockOrgRepo, _, _, service := setupMocks()

// 	existingBid := &entities.Bid{
// 		ID:             1,
// 		TenderID:       1,
// 		OrganizationID: 1,
// 		Status:         "CREATED",
// 	}

// 	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)
// 	mockOrgRepo.On("GetResponsibles", mock.Anything, 1).Return([]*user_models.UserModel{
// 		{ID: 1},
// 	}, nil)

// 	mockBidRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Bid")).Return(nil)

// 	err := service.RejectBid(context.Background(), 1, 1)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "REJECTED", existingBid.Status)
// 	mockBidRepo.AssertExpectations(t)
// 	mockOrgRepo.AssertExpectations(t)
// }

func TestRejectBid_NotFound(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	err := service.RejectBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestRejectBid_UserNotResponsible(t *testing.T) {
	mockBidRepo, mockOrgRepo, _, _, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)

	mockOrgRepo.On("GetResponsibles", mock.Anything, 1).Return([]*user_models.UserModel{}, nil)

	err := service.RejectBid(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "user is not responsible for the organization", err.Error())
	mockBidRepo.AssertExpectations(t)
	mockOrgRepo.AssertExpectations(t)
}

func TestRollbackBidVersion_Success(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		Status:         "CREATED",
	}

	rollbackVersion := &entities.BidVersion{
		BidID:       1,
		Name:        "Version 1",
		Description: "Description",
		Version:     1,
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
	mockBidRepo, _, _, _, service := setupMocks()

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(nil, errors.New("bid not found"))

	_, err := service.RollbackBidVersion(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "bid not found", err.Error())
	mockBidRepo.AssertExpectations(t)
}

func TestDeleteBid_Success(t *testing.T) {
	mockBidRepo, _, _, _, service := setupMocks()

	existingBid := &entities.Bid{
		ID:             1,
		TenderID:       1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "CREATED",
		CreatedAt:      time.Now(),
	}

	mockBidRepo.On("FindByID", mock.Anything, 1).Return(existingBid, nil)
	mockBidRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteBid(context.Background(), 1)

	assert.NoError(t, err)
	mockBidRepo.AssertExpectations(t)
}
