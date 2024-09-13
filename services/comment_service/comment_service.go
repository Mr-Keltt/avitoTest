package comment_service

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/comment_repository"
	"avitoTest/services/comment_service/comment_models"
	"context"
	"errors"
	"time"
)

type CommentService interface {
	CreateComment(ctx context.Context, model comment_models.CommentCreateModel) (*comment_models.CommentModel, error)
	GetCommentsByUsername(ctx context.Context, username string) ([]*comment_models.CommentModel, error)
	GetCommentsByOrganizationID(ctx context.Context, organizationID int) ([]*comment_models.CommentModel, error)
	DeleteComment(ctx context.Context, id int) error
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (s *commentService) CreateComment(ctx context.Context, model comment_models.CommentCreateModel) (*comment_models.CommentModel, error) {
	if model.UserID == 0 {
		return nil, errors.New("user ID is required")
	}

	if model.OrganizationID == 0 {
		return nil, errors.New("organization ID is required")
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

func (s *commentService) GetCommentsByUsername(ctx context.Context, username string) ([]*comment_models.CommentModel, error) {
	comments, err := s.commentRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	var result []*comment_models.CommentModel
	for _, comment := range comments {
		result = append(result, &comment_models.CommentModel{
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

	return result, nil
}

func (s *commentService) GetCommentsByOrganizationID(ctx context.Context, organizationID int) ([]*comment_models.CommentModel, error) {
	comments, err := s.commentRepo.FindByOrganizationID(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	var result []*comment_models.CommentModel
	for _, comment := range comments {
		result = append(result, &comment_models.CommentModel{
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

	return result, nil
}

func (s *commentService) DeleteComment(ctx context.Context, id int) error {
	return s.commentRepo.Delete(ctx, id)
}
