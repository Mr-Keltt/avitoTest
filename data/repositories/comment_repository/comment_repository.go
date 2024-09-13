package comment_repository

import (
	"avitoTest/data/entities"
	"context"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entities.Comment) error
	FindByUsername(ctx context.Context, username string) ([]*entities.Comment, error)
	FindByOrganizationID(ctx context.Context, organizationID int) ([]*entities.Comment, error)
	Delete(ctx context.Context, id int) error
}
