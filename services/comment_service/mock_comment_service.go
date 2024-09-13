package comment_service

import (
	"avitoTest/services/comment_service/comment_models"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCommentService struct {
	mock.Mock
}

func (m *MockCommentService) CreateComment(ctx context.Context, model comment_models.CommentCreateModel) (*comment_models.CommentModel, error) {
	args := m.Called(ctx, model)
	if comment, ok := args.Get(0).(*comment_models.CommentModel); ok {
		return comment, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCommentService) GetCommentsByFilters(ctx context.Context, authorUsername string, organizationID int) ([]*comment_models.CommentModel, error) {
	args := m.Called(ctx, authorUsername, organizationID)
	if comments, ok := args.Get(0).([]*comment_models.CommentModel); ok {
		return comments, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCommentService) DeleteComment(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
