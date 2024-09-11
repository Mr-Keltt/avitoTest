// File: test/service_tests/tender_service_test/tender_service_test.go

package tender_service_test

import (
	"avitoTest/data/entities"
	"avitoTest/services/tender_service"
	"avitoTest/services/tender_service/tender_models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTenderRepository is a mock implementation of TenderRepository.
type MockTenderRepository struct {
	mock.Mock
}

func (m *MockTenderRepository) Create(ctx context.Context, tender *entities.Tender) error {
	args := m.Called(ctx, tender)
	return args.Error(0)
}

func (m *MockTenderRepository) Update(ctx context.Context, tender *entities.Tender) error {
	args := m.Called(ctx, tender)
	return args.Error(0)
}

func (m *MockTenderRepository) FindByID(ctx context.Context, id int) (*entities.Tender, error) {
	args := m.Called(ctx, id)
	if tender, ok := args.Get(0).(*entities.Tender); ok {
		return tender, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenderRepository) GetAll(ctx context.Context) ([]*entities.Tender, error) {
	args := m.Called(ctx)
	if tenders, ok := args.Get(0).([]*entities.Tender); ok {
		return tenders, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenderRepository) CreateVersion(ctx context.Context, version *entities.TenderVersion) error {
	args := m.Called(ctx, version)
	return args.Error(0)
}

func (m *MockTenderRepository) FindVersionByNumber(ctx context.Context, tenderID int, versionNumber int) (*entities.TenderVersion, error) {
	args := m.Called(ctx, tenderID, versionNumber)
	if version, ok := args.Get(0).(*entities.TenderVersion); ok {
		return version, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenderRepository) FindLatestVersion(ctx context.Context, tenderID int) (*entities.TenderVersion, error) {
	args := m.Called(ctx, tenderID)
	if version, ok := args.Get(0).(*entities.TenderVersion); ok {
		return version, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenderRepository) FindUserOrganizationResponsibility(ctx context.Context, userID, orgID int) (*entities.OrganizationResponsible, error) {
	args := m.Called(ctx, userID, orgID)
	if responsible, ok := args.Get(0).(*entities.OrganizationResponsible); ok {
		return responsible, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateTender_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	tenderCreate := tender_models.TenderCreateModel{
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		OrganizationID: 1,
		CreatorID:      1,
	}

	expectedEntity := &entities.Tender{
		ID:             1,
		OrganizationID: tenderCreate.OrganizationID,
		CreatorID:      tenderCreate.CreatorID,
		CreatedAt:      time.Now(),
	}

	expectedVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    expectedEntity.ID,
		Name:        tenderCreate.Name,
		Description: tenderCreate.Description,
		ServiceType: tenderCreate.ServiceType,
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindUserOrganizationResponsibility", mock.Anything, tenderCreate.CreatorID, tenderCreate.OrganizationID).Return(&entities.OrganizationResponsible{}, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Tender")).Return(nil).Run(func(args mock.Arguments) {
		tender := args.Get(1).(*entities.Tender)
		tender.ID = expectedEntity.ID
		tender.CreatedAt = expectedEntity.CreatedAt
	})
	mockRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil).Run(func(args mock.Arguments) {
		version := args.Get(1).(*entities.TenderVersion)
		version.ID = expectedVersion.ID
		version.UpdatedAt = expectedVersion.UpdatedAt
	})

	result, err := service.CreateTender(context.Background(), tenderCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedVersion.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTender_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	tenderUpdate := tender_models.TenderUpdateModel{
		ID:          1,
		Name:        "Updated Tender",
		Description: "Updated Description",
		ServiceType: "Renovation",
	}

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, tenderUpdate.ID).Return(existingEntity, nil)
	mockRepo.On("FindLatestVersion", mock.Anything, tenderUpdate.ID).Return(latestVersion, nil)
	mockRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil)

	result, err := service.UpdateTender(context.Background(), tenderUpdate)

	assert.NoError(t, err)
	assert.Equal(t, tenderUpdate.Name, result.Name)
	assert.Equal(t, tenderUpdate.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestPublishTender_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)
	mockRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil)

	err := service.PublishTender(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCloseTender_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)
	mockRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil)

	err := service.CloseTender(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRollbackTenderVersion_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
	}

	previousVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("FindVersionByNumber", mock.Anything, 1, 1).Return(previousVersion, nil)
	mockRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil)

	result, err := service.RollbackTenderVersion(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, previousVersion.Name, result.Name)
	assert.Equal(t, previousVersion.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestGetTenderByID_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)

	result, err := service.GetTenderByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, latestVersion.Name, result.Name)
	assert.Equal(t, latestVersion.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTenders_Success(t *testing.T) {
	mockRepo := new(MockTenderRepository)
	service := tender_service.NewTenderService(mockRepo)

	existingEntities := []*entities.Tender{
		{
			ID:             1,
			OrganizationID: 1,
			CreatorID:      1,
			CreatedAt:      time.Now(),
		},
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntities[0].ID,
		Name:        "Tender 1",
		Description: "Description 1",
		ServiceType: "Construction",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("GetAll", mock.Anything).Return(existingEntities, nil)
	mockRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)

	result, err := service.GetAllTenders(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, latestVersion.Name, result[0].Name)
	assert.Equal(t, latestVersion.Description, result[0].Description)
	mockRepo.AssertExpectations(t)
}
