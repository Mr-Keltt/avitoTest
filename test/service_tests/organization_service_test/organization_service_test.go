package organization_service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/services/organization_service/organization_models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockOrganizationRepository is a mock implementation of OrganizationRepository.
type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) Create(ctx context.Context, org *entities.Organization) error {
	args := m.Called(ctx, org)
	return args.Error(0)
}

func (m *MockOrganizationRepository) GetAll(ctx context.Context) ([]entities.Organization, error) {
	args := m.Called(ctx)
	if orgs, ok := args.Get(0).([]entities.Organization); ok {
		return orgs, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationRepository) FindByID(ctx context.Context, id int) (*entities.Organization, error) {
	args := m.Called(ctx, id)
	if org, ok := args.Get(0).(*entities.Organization); ok {
		return org, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationRepository) Update(ctx context.Context, org *entities.Organization) error {
	args := m.Called(ctx, org)
	return args.Error(0)
}

func (m *MockOrganizationRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// New methods for responsible management
func (m *MockOrganizationRepository) AddResponsible(ctx context.Context, orgResponsible *entities.OrganizationResponsible) error {
	args := m.Called(ctx, orgResponsible)
	return args.Error(0)
}

func (m *MockOrganizationRepository) DeleteResponsible(ctx context.Context, orgID int, userID int) error {
	args := m.Called(ctx, orgID, userID)
	return args.Error(0)
}

func (m *MockOrganizationRepository) GetResponsibles(ctx context.Context, orgID int) ([]entities.User, error) {
	args := m.Called(ctx, orgID)
	if responsibles, ok := args.Get(0).([]entities.User); ok {
		return responsibles, args.Error(1)
	}
	return nil, args.Error(1)
}

// Add a new method to MockOrganizationRepository for GetResponsibleByID
func (m *MockOrganizationRepository) GetResponsibleByID(ctx context.Context, orgID int, userID int) (*entities.User, error) {
	args := m.Called(ctx, orgID, userID)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
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

func setupMocks() (*MockOrganizationRepository, *MockUserRepository, organization_service.OrganizationService) {
	mockOrgRepo := new(MockOrganizationRepository)
	mockUserRepo := new(MockUserRepository)
	service := organization_service.NewOrganizationService(mockOrgRepo, mockUserRepo)
	return mockOrgRepo, mockUserRepo, service
}

func TestCreateOrganization_Success(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	orgCreate := organization_models.OrganizationCreateModel{
		Name:        "My Organization",
		Description: "Test Description",
		Type:        "LLC",
	}

	expectedEntity := &entities.Organization{
		ID:          1,
		Name:        orgCreate.Name,
		Description: orgCreate.Description,
		Type:        entities.OrganizationType(orgCreate.Type),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockOrgRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Organization")).Return(nil).Run(func(args mock.Arguments) {
		org := args.Get(1).(*entities.Organization)
		org.ID = expectedEntity.ID
		org.CreatedAt = expectedEntity.CreatedAt
		org.UpdatedAt = expectedEntity.UpdatedAt
	})

	result, err := service.CreateOrganization(context.Background(), orgCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Name, result.Name)
	mockOrgRepo.AssertExpectations(t)
}

func TestCreateOrganization_ValidationFail(t *testing.T) {
	_, _, service := setupMocks()

	orgCreate := organization_models.OrganizationCreateModel{
		Name:        "",
		Description: "Test Description",
		Type:        "LLC",
	}

	_, err := service.CreateOrganization(context.Background(), orgCreate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed on the 'required' tag")
}

func TestGetOrganizations_Success(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	expectedEntities := []entities.Organization{
		{
			ID:          1,
			Name:        "Org 1",
			Description: "Description 1",
			Type:        entities.OrganizationType("LLC"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Org 2",
			Description: "Description 2",
			Type:        entities.OrganizationType("Corporation"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	mockOrgRepo.On("GetAll", mock.Anything).Return(expectedEntities, nil)

	result, err := service.GetOrganizations(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedEntities[0].ID, result[0].ID)
	assert.Equal(t, expectedEntities[1].ID, result[1].ID)
	mockOrgRepo.AssertExpectations(t)
}

func TestGetOrganizations_Failure(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	mockOrgRepo.On("GetAll", mock.Anything).Return(nil, organization_repository.ErrOrganizationNotFound)

	_, err := service.GetOrganizations(context.Background())

	assert.Error(t, err)
	assert.Equal(t, organization_repository.ErrOrganizationNotFound, err)
	mockOrgRepo.AssertExpectations(t)
}

// Tests for AddResponsible, DeleteResponsible, and GetResponsibles methods

func TestAddResponsible_Success(t *testing.T) {
	mockOrgRepo, mockUserRepo, service := setupMocks()

	// Mock organization and user existence
	mockOrgRepo.On("FindByID", mock.Anything, 1).Return(&entities.Organization{ID: 1}, nil)
	mockUserRepo.On("FindByID", mock.Anything, 1).Return(&entities.User{ID: 1}, nil)
	mockOrgRepo.On("AddResponsible", mock.Anything, mock.AnythingOfType("*entities.OrganizationResponsible")).Return(nil)

	err := service.AddResponsible(context.Background(), 1, 1)

	assert.NoError(t, err)
	mockOrgRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestDeleteResponsible_Success(t *testing.T) {
	mockOrgRepo, mockUserRepo, service := setupMocks()

	// Mock organization and user existence
	mockOrgRepo.On("FindByID", mock.Anything, 1).Return(&entities.Organization{ID: 1}, nil)
	mockUserRepo.On("FindByID", mock.Anything, 1).Return(&entities.User{ID: 1}, nil)
	mockOrgRepo.On("DeleteResponsible", mock.Anything, 1, 1).Return(nil)

	err := service.DeleteResponsible(context.Background(), 1, 1)

	assert.NoError(t, err)
	mockOrgRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestGetResponsibles_Success(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	expectedUsers := []entities.User{
		{
			ID:        1,
			Username:  "user1",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			ID:        2,
			Username:  "user2",
			FirstName: "Jane",
			LastName:  "Smith",
		},
	}

	mockOrgRepo.On("GetResponsibles", mock.Anything, 1).Return(expectedUsers, nil)

	result, err := service.GetResponsibles(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedUsers[0].ID, result[0].ID)
	assert.Equal(t, expectedUsers[1].ID, result[1].ID)
	mockOrgRepo.AssertExpectations(t)
}

// Test case for GetResponsibleByID
func TestGetResponsibleByID_Success(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	expectedUser := &entities.User{
		ID:        1,
		Username:  "user1",
		FirstName: "John",
		LastName:  "Doe",
	}

	// Mock the organization and user existence
	mockOrgRepo.On("FindByID", mock.Anything, 1).Return(&entities.Organization{ID: 1}, nil)
	mockOrgRepo.On("GetResponsibleByID", mock.Anything, 1, 1).Return(expectedUser, nil)

	result, err := service.GetResponsibleByID(context.Background(), 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Username, result.Username)
	assert.Equal(t, expectedUser.FirstName, result.FirstName)
	assert.Equal(t, expectedUser.LastName, result.LastName)
	mockOrgRepo.AssertExpectations(t)
}

func TestGetResponsibleByID_OrganizationNotFound(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	// Mock organization not found
	mockOrgRepo.On("FindByID", mock.Anything, 1).Return(nil, organization_repository.ErrOrganizationNotFound)

	_, err := service.GetResponsibleByID(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "organization not found", err.Error())
	mockOrgRepo.AssertExpectations(t)
}

func TestGetResponsibleByID_ResponsibleNotFound(t *testing.T) {
	mockOrgRepo, _, service := setupMocks()

	// Mock organization existence
	mockOrgRepo.On("FindByID", mock.Anything, 1).Return(&entities.Organization{ID: 1}, nil)
	// Mock responsible user not found
	mockOrgRepo.On("GetResponsibleByID", mock.Anything, 1, 1).Return(nil, errors.New("responsible user not found"))

	_, err := service.GetResponsibleByID(context.Background(), 1, 1)

	assert.Error(t, err)
	assert.Equal(t, "responsible user not found", err.Error())
	mockOrgRepo.AssertExpectations(t)
}
