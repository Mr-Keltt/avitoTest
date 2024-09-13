package commen_thandler_test

import (
	"avitoTest/api/handlers/comment_handler"
	"avitoTest/services/comment_service"
	"avitoTest/services/comment_service/comment_models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*comment_service.MockCommentService, *comment_handler.CommentHandler) {
	mockService := new(comment_service.MockCommentService)
	handler := comment_handler.NewCommentHandler(mockService)
	return mockService, handler
}

func TestCreateComment_Success(t *testing.T) {
	mockService, handler := setup()

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

	expectedResponse := &comment_models.CommentModel{
		ID:                1,
		UserID:            commentCreate.UserID,
		OrganizationID:    commentCreate.OrganizationID,
		CompanyName:       commentCreate.CompanyName,
		TenderName:        commentCreate.TenderName,
		TenderDescription: commentCreate.TenderDescription,
		BidDescription:    commentCreate.BidDescription,
		ServiceType:       commentCreate.ServiceType,
		Content:           commentCreate.Content,
		CreatedAt:         "2024-09-13T15:20:30Z",
	}

	mockService.On("CreateComment", mock.Anything, commentCreate).Return(expectedResponse, nil)

	requestBody, _ := json.Marshal(commentCreate)
	req := httptest.NewRequest(http.MethodPost, "/api/comments", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateComment(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response comment_models.CommentModel
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Content, response.Content)
	mockService.AssertExpectations(t)
}

func TestCreateComment_ValidationFail(t *testing.T) {
	_, handler := setup()

	commentCreate := comment_models.CommentCreateModel{
		UserID:         0,
		OrganizationID: 1,
		Content:        "This is a comment",
	}

	requestBody, _ := json.Marshal(commentCreate)
	req := httptest.NewRequest(http.MethodPost, "/api/comments", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetReviews_Success(t *testing.T) {
	mockService, handler := setup()

	expectedComments := []*comment_models.CommentModel{
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
			CreatedAt:         "2024-09-13T15:20:30Z",
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
			CreatedAt:         "2024-09-13T15:20:30Z",
		},
	}

	mockService.On("GetCommentsByFilters", mock.Anything, "testuser", 1).Return(expectedComments, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/bids/1/reviews?authorUsername=testuser&organizationId=1", nil)
	w := httptest.NewRecorder()

	handler.GetReviews(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []*comment_models.CommentModel
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "First comment", response[0].Content)
	mockService.AssertExpectations(t)
}

func TestDeleteComment_Success(t *testing.T) {
	mockService, handler := setup()

	mockService.On("DeleteComment", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/comments/1", nil)
	w := httptest.NewRecorder()

	handler.DeleteComment(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteComment_NotFound(t *testing.T) {
	mockService, handler := setup()

	mockService.On("DeleteComment", mock.Anything, 1).Return(errors.New("comment not found"))

	req := httptest.NewRequest(http.MethodDelete, "/api/comments/1", nil)
	w := httptest.NewRecorder()

	handler.DeleteComment(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
