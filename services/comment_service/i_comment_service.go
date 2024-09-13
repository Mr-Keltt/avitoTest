package comment_service

import (
	"avitoTest/services/comment_service/comment_models"
	"context"
)

type CommentService interface {
	CreateComment(ctx context.Context, model comment_models.CommentCreateModel) (*comment_models.CommentModel, error)
	GetCommentsByFilters(ctx context.Context, authorUsername string, organizationID int) ([]*comment_models.CommentModel, error)
	DeleteComment(ctx context.Context, id int) error
}
