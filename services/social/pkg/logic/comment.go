package logic

import (
	"context"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type CommentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

// CreateComment creates a new comment.
func (s *CommentService) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	newTime := time.Now()
	comment.CreatedAt = &newTime
	comment.UpdatedAt = &newTime
	return s.commentRepo.CreateComment(ctx, comment)
}

// GetCommentByID retrieves a comment by its ID.
func (s *CommentService) GetCommentByID(ctx context.Context, id string) (*models.Comment, error) {
	return s.commentRepo.GetCommentByID(ctx, id)
}

// GetAllComments retrieves all comments.
func (s *CommentService) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	return s.commentRepo.GetAllComments(ctx)
}

// UpdateComment updates an existing comment.
func (s *CommentService) UpdateComment(ctx context.Context, comment *models.Comment) error {
	newTime := time.Now()
	comment.UpdatedAt = &newTime
	return s.commentRepo.UpdateComment(ctx, comment)
}

// DeleteComment deletes a comment by its ID.
func (s *CommentService) DeleteComment(ctx context.Context, id string) error {
	return s.commentRepo.DeleteComment(ctx, id)
}

// GetCommentsByPostID retrieves all comments for a specific post.
func (s *CommentService) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.Comment, error) {
	return s.commentRepo.GetCommentsByPostID(ctx, postID)
}
