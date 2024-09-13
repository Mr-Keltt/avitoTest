package comment_repository

import (
	"avitoTest/data/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) Create(ctx context.Context, comment *entities.Comment) error {
	args := m.Called(ctx, comment)
	return args.Error(0)
}

func (m *MockCommentRepository) FindByFilters(ctx context.Context, authorUsername string, organizationID int) ([]*entities.Comment, error) {
	args := m.Called(ctx, authorUsername, organizationID)
	if comments, ok := args.Get(0).([]*entities.Comment); ok {
		return comments, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCommentRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
