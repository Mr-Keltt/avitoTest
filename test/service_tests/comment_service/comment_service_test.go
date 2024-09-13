package comment_service_test

import (
	"avitoTest/data/entities"
	"avitoTest/data/repositories/comment_repository"
	"avitoTest/services/comment_service"
	"avitoTest/services/comment_service/comment_models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Мок репозиторий
func setupMocks() (*comment_repository.MockCommentRepository, comment_service.CommentService) {
	mockCommentRepo := new(comment_repository.MockCommentRepository)
	service := comment_service.NewCommentService(mockCommentRepo)
	return mockCommentRepo, service
}

func TestCreateComment_Success(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	commentCreate := comment_models.CommentCreateModel{
		UserID:            1,
		OrganizationID:    1,
		CompanyName:       "Test Company",
		TenderName:        "Test Tender",
		TenderDescription: "Test Tender Description",
		BidDescription:    "Test Bid Description",
		ServiceType:       "Consulting",
		Content:           "This is a comment.",
	}

	expectedEntity := &entities.Comment{
		ID:                1,
		UserID:            commentCreate.UserID,
		OrganizationID:    commentCreate.OrganizationID,
		CompanyName:       commentCreate.CompanyName,
		TenderName:        commentCreate.TenderName,
		TenderDescription: commentCreate.TenderDescription,
		BidDescription:    commentCreate.BidDescription,
		ServiceType:       commentCreate.ServiceType,
		Content:           commentCreate.Content,
		CreatedAt:         time.Now(),
	}

	mockCommentRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Comment")).Return(nil).Run(func(args mock.Arguments) {
		comment := args.Get(1).(*entities.Comment)
		comment.ID = expectedEntity.ID
		comment.CreatedAt = expectedEntity.CreatedAt
	})

	result, err := service.CreateComment(context.Background(), commentCreate)

	assert.NoError(t, err)
	assert.Equal(t, expectedEntity.ID, result.ID)
	assert.Equal(t, commentCreate.Content, result.Content)
	mockCommentRepo.AssertExpectations(t)
}

func TestCreateComment_ValidationFail(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	commentCreate := comment_models.CommentCreateModel{
		UserID:         0, // Invalid: UserID is required
		OrganizationID: 1,
		Content:        "Test comment",
	}

	_, err := service.CreateComment(context.Background(), commentCreate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user ID is required")

	mockCommentRepo.AssertNotCalled(t, "Create")
}

func TestGetCommentsByFilters_Success(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	expectedComments := []*entities.Comment{
		{
			ID:                1,
			UserID:            1,
			OrganizationID:    1,
			CompanyName:       "Test Company",
			TenderName:        "Test Tender",
			TenderDescription: "Test Tender Description",
			BidDescription:    "Test Bid Description",
			ServiceType:       "Consulting",
			Content:           "First comment",
			CreatedAt:         time.Now(),
		},
		{
			ID:                2,
			UserID:            1,
			OrganizationID:    1,
			CompanyName:       "Test Company",
			TenderName:        "Test Tender",
			TenderDescription: "Test Tender Description",
			BidDescription:    "Test Bid Description",
			ServiceType:       "Consulting",
			Content:           "Second comment",
			CreatedAt:         time.Now(),
		},
	}

	mockCommentRepo.On("FindByFilters", mock.Anything, "testuser", 1).Return(expectedComments, nil)

	result, err := service.GetCommentsByFilters(context.Background(), "testuser", 1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "First comment", result[0].Content)
	assert.Equal(t, "Second comment", result[1].Content)
	mockCommentRepo.AssertExpectations(t)
}

func TestGetCommentsByFilters_NoOrganization_Success(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	expectedComments := []*entities.Comment{
		{
			ID:                1,
			UserID:            1,
			OrganizationID:    1,
			CompanyName:       "Test Company",
			TenderName:        "Test Tender",
			TenderDescription: "Test Tender Description",
			BidDescription:    "Test Bid Description",
			ServiceType:       "Consulting",
			Content:           "First comment",
			CreatedAt:         time.Now(),
		},
	}

	mockCommentRepo.On("FindByFilters", mock.Anything, "testuser", 0).Return(expectedComments, nil)

	result, err := service.GetCommentsByFilters(context.Background(), "testuser", 0)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "First comment", result[0].Content)
	mockCommentRepo.AssertExpectations(t)
}

func TestGetCommentsByFilters_UserNotFound(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	mockCommentRepo.On("FindByFilters", mock.Anything, "unknownuser", 0).Return(nil, errors.New("no comments found"))

	_, err := service.GetCommentsByFilters(context.Background(), "unknownuser", 0)

	assert.Error(t, err)
	assert.Equal(t, "no comments found", err.Error())
	mockCommentRepo.AssertExpectations(t)
}

func TestDeleteComment_Success(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	mockCommentRepo.On("Delete", mock.Anything, 1).Return(nil)

	err := service.DeleteComment(context.Background(), 1)

	assert.NoError(t, err)
	mockCommentRepo.AssertExpectations(t)
}

func TestDeleteComment_NotFound(t *testing.T) {
	mockCommentRepo, service := setupMocks()

	mockCommentRepo.On("Delete", mock.Anything, 1).Return(errors.New("comment not found"))

	err := service.DeleteComment(context.Background(), 1)

	assert.Error(t, err)
	assert.Equal(t, "comment not found", err.Error())
	mockCommentRepo.AssertExpectations(t)
}
