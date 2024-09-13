package tender_repository

import (
	"avitoTest/data/entities"
	"context"

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

func (m *MockTenderRepository) CloseTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderRepository) PublishTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}
