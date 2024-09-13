package comment_repository

import (
	"avitoTest/data/entities"
	"context"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entities.Comment) error
	FindByFilters(ctx context.Context, authorUsername string, organizationID int) ([]*entities.Comment, error)
	Delete(ctx context.Context, id int) error
}
