package organization_service_test

import (
	"context"
	"testing"
	"time"

	"avitoTest/data/entities"
	"avitoTest/data/repositories/organization_repository"
	"avitoTest/services/organization_service"
	"avitoTest/services/organization_service/service_models"

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

func TestCreateOrganization_Success(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	orgCreate := service_models.OrganizationCreateModel{
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

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Organization")).Return(nil).Run(func(args mock.Arguments) {
		org := args.Get(1).(*entities.Organization)
		org.ID = expectedEntity.ID
		org.CreatedAt = expectedEntity.CreatedAt
		org.UpdatedAt = expectedEntity.UpdatedAt
	})

	result, err := service.CreateOrganization(context.Background(), orgCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrganization_ValidationFail(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	orgCreate := service_models.OrganizationCreateModel{
		Name:        "",
		Description: "Test Description",
		Type:        "LLC",
	}

	_, err := service.CreateOrganization(context.Background(), orgCreate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed on the 'required' tag")
}

func TestGetOrganizations_Success(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

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

	mockRepo.On("GetAll", mock.Anything).Return(expectedEntities, nil)

	result, err := service.GetOrganizations(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedEntities[0].ID, result[0].ID)
	assert.Equal(t, expectedEntities[1].ID, result[1].ID)
	mockRepo.AssertExpectations(t)
}

func TestGetOrganizations_Failure(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	mockRepo.On("GetAll", mock.Anything).Return(nil, organization_repository.ErrOrganizationNotFound)

	_, err := service.GetOrganizations(context.Background())

	assert.Error(t, err)
	assert.Equal(t, organization_repository.ErrOrganizationNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestGetOrganizationByID_Success(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	expectedEntity := &entities.Organization{
		ID:          1,
		Name:        "My Organization",
		Description: "Test Description",
		Type:        entities.OrganizationType("LLC"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(expectedEntity, nil)

	result, err := service.GetOrganizationByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, expectedEntity.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetOrganizationByID_NotFound(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(nil, organization_repository.ErrOrganizationNotFound)

	_, err := service.GetOrganizationByID(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "organization not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateOrganization_Success(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	orgUpdate := service_models.OrganizationUpdateModel{
		ID:          1,
		Name:        "Updated Organization",
		Description: "Updated Description",
		Type:        "LLC",
	}

	existingEntity := &entities.Organization{
		ID:          1,
		Name:        "Old Organization",
		Description: "Old Description",
		Type:        entities.OrganizationType("LLC"),
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, 1).Return(existingEntity, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*entities.Organization")).Return(nil)

	result, err := service.UpdateOrganization(context.Background(), orgUpdate)

	assert.NoError(t, err)
	assert.Equal(t, orgUpdate.Name, result.Name)
	assert.Equal(t, orgUpdate.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestDeleteOrganization_Success(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(&entities.Organization{ID: 1}, nil)
	mockRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteOrganization(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteOrganization_NotFound(t *testing.T) {
	mockRepo := new(MockOrganizationRepository)
	service := organization_service.NewOrganizationService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, 1).Return(nil, organization_repository.ErrOrganizationNotFound)

	err := service.DeleteOrganization(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "organization not found", err.Error())
	mockRepo.AssertExpectations(t)
}
