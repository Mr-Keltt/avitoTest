package service_test

import (
	"avitoTest/data/entities"
	"avitoTest/services/tender_service"
	"avitoTest/services/tender_service/tender_models"
	"avitoTest/shared/constants"
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

func (m *MockTenderRepository) GetAllByServiceType(ctx context.Context, serviceType string) ([]*entities.Tender, error) {
	args := m.Called(ctx, serviceType)
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

func (m *MockTenderRepository) GetAllByCreatorID(ctx context.Context, creatorID int) ([]*entities.Tender, error) {
	args := m.Called(ctx, creatorID)
	if tenders, ok := args.Get(0).([]*entities.Tender); ok {
		return tenders, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenderRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockUserRepository is a mock implementation of UserRepository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*entities.User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

// Добавляем недостающий метод FindByUsername для корректного выполнения тестов
func (m *MockUserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	args := m.Called(ctx, username)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	args := m.Called(ctx)
	if users, ok := args.Get(0).([]entities.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Test for CreateTender
func TestCreateTender_Success(t *testing.T) {
	mockTenderRepo := new(MockTenderRepository)
	mockUserRepo := new(MockUserRepository)
	service := tender_service.NewTenderService(mockTenderRepo, mockUserRepo)

	tenderCreate := tender_models.TenderCreateModel{
		Name:           "Tender 1",
		Description:    "Description 1",
		ServiceType:    "Construction",
		OrganizationID: 1,
		CreatorID:      1,
		Status:         constants.TenderStatusCreated,
	}

	expectedEntity := &entities.Tender{
		ID:             1,
		OrganizationID: tenderCreate.OrganizationID,
		CreatorID:      tenderCreate.CreatorID,
		ServiceType:    tenderCreate.ServiceType,
		CreatedAt:      time.Now(),
	}

	expectedVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    expectedEntity.ID,
		Name:        tenderCreate.Name,
		Description: tenderCreate.Description,
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockTenderRepo.On("FindUserOrganizationResponsibility", mock.Anything, tenderCreate.CreatorID, tenderCreate.OrganizationID).Return(&entities.OrganizationResponsible{}, nil)
	mockTenderRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Tender")).Return(nil).Run(func(args mock.Arguments) {
		tender := args.Get(1).(*entities.Tender)
		tender.ID = expectedEntity.ID
		tender.CreatedAt = expectedEntity.CreatedAt
	})
	mockTenderRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil).Run(func(args mock.Arguments) {
		version := args.Get(1).(*entities.TenderVersion)
		version.ID = expectedVersion.ID
		version.UpdatedAt = expectedVersion.UpdatedAt
	})

	result, err := service.CreateTender(context.Background(), tenderCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedVersion.Name, result.Name)
	mockTenderRepo.AssertExpectations(t)
}

// Test for UpdateTender
func TestUpdateTender_Success(t *testing.T) {
	mockTenderRepo := new(MockTenderRepository)
	mockUserRepo := new(MockUserRepository)
	service := tender_service.NewTenderService(mockTenderRepo, mockUserRepo)

	tenderUpdate := tender_models.TenderUpdateModel{
		ID:          1,
		Name:        "Updated Tender",
		Description: "Updated Description",
	}

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		ServiceType:    "Renovation",
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockTenderRepo.On("FindByID", mock.Anything, tenderUpdate.ID).Return(existingEntity, nil)
	mockTenderRepo.On("FindLatestVersion", mock.Anything, tenderUpdate.ID).Return(latestVersion, nil)
	mockTenderRepo.On("CreateVersion", mock.Anything, mock.AnythingOfType("*entities.TenderVersion")).Return(nil)

	result, err := service.UpdateTender(context.Background(), tenderUpdate)

	assert.NoError(t, err)
	assert.Equal(t, tenderUpdate.Name, result.Name)
	assert.Equal(t, tenderUpdate.Description, result.Description)
	mockTenderRepo.AssertExpectations(t)
}

// Test for GetTendersByUsername
func TestGetTendersByUsername_Success(t *testing.T) {
	mockTenderRepo := new(MockTenderRepository)
	mockUserRepo := new(MockUserRepository)
	service := tender_service.NewTenderService(mockTenderRepo, mockUserRepo)

	username := "test_user"
	expectedUser := &entities.User{ID: 1, Username: username}
	tenders := []*entities.Tender{
		{
			ID:             1,
			OrganizationID: 1,
			CreatorID:      1,
			ServiceType:    "Construction",
			Status:         "created",
			CreatedAt:      time.Now(),
		},
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    tenders[0].ID,
		Name:        "Tender 1",
		Description: "Description 1",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockUserRepo.On("FindByUsername", mock.Anything, username).Return(expectedUser, nil)
	mockTenderRepo.On("GetAllByCreatorID", mock.Anything, expectedUser.ID).Return(tenders, nil)
	mockTenderRepo.On("FindLatestVersion", mock.Anything, tenders[0].ID).Return(latestVersion, nil)

	result, err := service.GetTendersByUsername(context.Background(), username)

	assert.NoError(t, err)
	assert.Equal(t, len(tenders), len(result))
	assert.Equal(t, tenders[0].ID, result[0].ID)
	mockTenderRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// Test for DeleteTender
func TestDeleteTender_Success(t *testing.T) {
	mockTenderRepo := new(MockTenderRepository)
	mockUserRepo := new(MockUserRepository)
	service := tender_service.NewTenderService(mockTenderRepo, mockUserRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		Status:         "created",
		CreatedAt:      time.Now(),
	}

	mockTenderRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockTenderRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteTender(context.Background(), 1)

	assert.NoError(t, err)
	mockTenderRepo.AssertExpectations(t)
}

// Test for GetTenderByID
func TestGetTenderByID_Success(t *testing.T) {
	mockTenderRepo := new(MockTenderRepository)
	mockUserRepo := new(MockUserRepository)
	service := tender_service.NewTenderService(mockTenderRepo, mockUserRepo)

	existingEntity := &entities.Tender{
		ID:             1,
		OrganizationID: 1,
		CreatorID:      1,
		ServiceType:    "Construction",
		CreatedAt:      time.Now(),
	}

	latestVersion := &entities.TenderVersion{
		ID:          1,
		TenderID:    existingEntity.ID,
		Name:        "Tender 1",
		Description: "Description 1",
		Version:     1,
		UpdatedAt:   time.Now(),
	}

	mockTenderRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockTenderRepo.On("FindLatestVersion", mock.Anything, 1).Return(latestVersion, nil)

	result, err := service.GetTenderByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, latestVersion.Name, result.Name)
	assert.Equal(t, latestVersion.Description, result.Description)
	mockTenderRepo.AssertExpectations(t)
}
