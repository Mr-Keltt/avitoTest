package comment_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/comment_repository"
	"avitoTest/services/comment_service/comment_models"
	"context"
	"errors"
	"time"
)

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{commentRepo: commentRepo}
}

func (s *commentService) CreateComment(ctx context.Context, model comment_models.CommentCreateModel) (*comment_models.CommentModel, error) {
	// Проверяем, что UserID не равен 0
	if model.UserID == 0 {
		return nil, errors.New("user ID is required")
	}

	comment := &entities.Comment{
		UserID:            model.UserID,
		OrganizationID:    model.OrganizationID,
		CompanyName:       model.CompanyName,
		TenderName:        model.TenderName,
		TenderDescription: model.TenderDescription,
		BidDescription:    model.BidDescription,
		ServiceType:       model.ServiceType,
		Content:           model.Content,
		CreatedAt:         time.Now(),
	}

	if err := s.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}

	return &comment_models.CommentModel{
		ID:                comment.ID,
		UserID:            comment.UserID,
		OrganizationID:    comment.OrganizationID,
		CompanyName:       comment.CompanyName,
		TenderName:        comment.TenderName,
		TenderDescription: comment.TenderDescription,
		BidDescription:    comment.BidDescription,
		ServiceType:       comment.ServiceType,
		Content:           comment.Content,
		CreatedAt:         comment.CreatedAt.Format(time.RFC3339),
	}, nil
}

// GetCommentsByFilters returns comments by authorUsername and organizationID (if specified)
func (s *commentService) GetCommentsByFilters(ctx context.Context, authorUsername string, organizationID int) ([]*comment_models.CommentModel, error) {
	comments, err := s.commentRepo.FindByFilters(ctx, authorUsername, organizationID)
	if err != nil {
		return nil, err
	}

	var commentModels []*comment_models.CommentModel
	for _, comment := range comments {
		commentModels = append(commentModels, &comment_models.CommentModel{
			ID:                comment.ID,
			UserID:            comment.UserID,
			OrganizationID:    comment.OrganizationID,
			CompanyName:       comment.CompanyName,
			TenderName:        comment.TenderName,
			TenderDescription: comment.TenderDescription,
			BidDescription:    comment.BidDescription,
			ServiceType:       comment.ServiceType,
			Content:           comment.Content,
			CreatedAt:         comment.CreatedAt.Format(time.RFC3339),
		})
	}

	return commentModels, nil
}

func (s *commentService) DeleteComment(ctx context.Context, id int) error {
	return s.commentRepo.Delete(ctx, id)
}
