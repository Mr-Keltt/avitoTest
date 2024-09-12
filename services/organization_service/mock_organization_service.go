package organization_service

import (
	"avitoTest/services/organization_service/organization_models"
	"avitoTest/services/user_service/user_models"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockOrganizationService is a mock implementation of the OrganizationService interface
type MockOrganizationService struct {
	mock.Mock
}

func (m *MockOrganizationService) CreateOrganization(ctx context.Context, org organization_models.OrganizationCreateModel) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, org)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) UpdateOrganization(ctx context.Context, org organization_models.OrganizationUpdateModel) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, org)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) GetOrganizations(ctx context.Context) ([]*organization_models.OrganizationModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) GetOrganizationByID(ctx context.Context, id int) (*organization_models.OrganizationModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*organization_models.OrganizationModel), args.Error(1)
}

func (m *MockOrganizationService) DeleteOrganization(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockOrganizationService) AddResponsible(ctx context.Context, orgID int, userID int) error {
	args := m.Called(ctx, orgID, userID)
	return args.Error(0)
}

func (m *MockOrganizationService) DeleteResponsible(ctx context.Context, orgID int, userID int) error {
	args := m.Called(ctx, orgID, userID)
	return args.Error(0)
}

func (m *MockOrganizationService) GetResponsibles(ctx context.Context, orgID int) ([]*user_models.UserModel, error) {
	args := m.Called(ctx, orgID)

	if responsibles, ok := args.Get(0).([]*user_models.UserModel); ok {
		return responsibles, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockOrganizationService) GetResponsibleByID(ctx context.Context, orgID int, userID int) (*user_models.UserModel, error) {
	args := m.Called(ctx, orgID, userID)
	if user, ok := args.Get(0).(*user_models.UserModel); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
