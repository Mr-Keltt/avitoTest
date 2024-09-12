package organization_repository

import (
	"avitoTest/data/entities"
	"context"

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

func (m *MockOrganizationRepository) GetResponsibleByID(ctx context.Context, orgID int, userID int) (*entities.User, error) {
	args := m.Called(ctx, orgID, userID)
	if user, ok := args.Get(0).(*entities.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}
