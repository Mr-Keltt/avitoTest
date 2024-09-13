package bid_repository

import (
	"avitoTest/data/entities"
	"context"
)

type BidRepository interface {
	Create(ctx context.Context, bid *entities.Bid) error
	Update(ctx context.Context, bid *entities.Bid) error
	FindByID(ctx context.Context, id int) (*entities.Bid, error)
	FindByTenderID(ctx context.Context, tenderID int) ([]*entities.Bid, error)
	FindByCreatorID(ctx context.Context, creatorID int) ([]*entities.Bid, error)
	FindByUsername(ctx context.Context, username string) ([]*entities.Bid, error)
	FindLatestVersion(ctx context.Context, bidID int) (*entities.BidVersion, error)
	FindVersionByNumber(ctx context.Context, bidID int, versionNumber int) (*entities.BidVersion, error)
	CreateVersion(ctx context.Context, version *entities.BidVersion) error
	Delete(ctx context.Context, bidID int) error
}
