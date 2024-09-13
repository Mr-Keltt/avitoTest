package comment_handler

import (
	"avitoTest/services/comment_service"
	"avitoTest/services/comment_service/comment_models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	service comment_service.CommentService
}

func NewCommentHandler(service comment_service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// CreateComment creates a new comment
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req comment_models.CommentCreateModel
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверка, что user_id присутствует
	if req.UserID == 0 {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	// Вызов сервиса для создания комментария
	comment, err := h.service.CreateComment(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetReviews returns comments by authorUsername and optionally organizationId
func (h *CommentHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	authorUsername := r.URL.Query().Get("authorUsername")
	if authorUsername == "" {
		http.Error(w, "authorUsername is required", http.StatusBadRequest)
		return
	}

	organizationIDStr := r.URL.Query().Get("organizationId")
	var organizationID int
	if organizationIDStr != "" {
		var err error
		organizationID, err = strconv.Atoi(organizationIDStr)
		if err != nil {
			http.Error(w, "Invalid organization ID", http.StatusBadRequest)
			return
		}
	}

	comments, err := h.service.GetCommentsByFilters(r.Context(), authorUsername, organizationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

// DeleteComment deletes a comment by its ID
func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentIDStr := mux.Vars(r)["commentId"]
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteComment(r.Context(), commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
