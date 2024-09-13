// File: services/tender_service/tender_service.go

package tender_service

import (
	"avitoTest/services/tender_service/tender_models"
	"context"
)

type TenderService interface {
	GetAllTenders(ctx context.Context, serviceTypeFilter string) ([]*tender_models.TenderModel, error)
	GetTendersByUsername(ctx context.Context, username string) ([]*tender_models.TenderModel, error)
	GetTenderByID(ctx context.Context, id int) (*tender_models.TenderModel, error)
	CreateTender(ctx context.Context, tender tender_models.TenderCreateModel) (*tender_models.TenderModel, error)
	UpdateTender(ctx context.Context, tender tender_models.TenderUpdateModel) (*tender_models.TenderModel, error)
	PublishTender(ctx context.Context, tenderID int) error
	CloseTender(ctx context.Context, tenderID int) error
	RollbackTenderVersion(ctx context.Context, tenderID int, version int) (*tender_models.TenderModel, error)
	DeleteTender(ctx context.Context, tenderID int) error
}
