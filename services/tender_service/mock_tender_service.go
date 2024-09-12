package tender_service

import (
	"avitoTest/services/tender_service/tender_models"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockTenderService is a mock implementation of the TenderService interface
type MockTenderService struct {
	mock.Mock
}

func (m *MockTenderService) CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tender)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tender)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) PublishTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderService) CloseTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderService) GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error) {
	args := m.Called(ctx, tenderID, version)
	return args.Get(0).(*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) GetAllTenders(ctx context.Context, serviceTypeFilter string) ([]*tender_models.TenderModel, error) {
	args := m.Called(ctx, serviceTypeFilter)
	return args.Get(0).([]*tender_models.TenderModel), args.Error(1)
}

func (m *MockTenderService) DeleteTender(ctx context.Context, tenderID int) error {
	args := m.Called(ctx, tenderID)
	return args.Error(0)
}

func (m *MockTenderService) GetTendersByUsername(ctx context.Context, username string) ([]*tender_models.TenderModel, error) {
	args := m.Called(ctx, username)
	return args.Get(0).([]*tender_models.TenderModel), args.Error(1)
}
