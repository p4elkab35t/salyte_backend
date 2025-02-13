package logic

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type InteractionService struct {
	interactionRepo repository.InteractionRepository
}

func NewInteractionService(interactionRepo repository.InteractionRepository) *InteractionService {
	return &InteractionService{interactionRepo: interactionRepo}
}

// GetPostComments retrieves all comments for a specific post.
func (s *InteractionService) GetPostComments(ctx context.Context, postID string) ([]*models.Comment, error) {
	return s.interactionRepo.GetPostComments(ctx, postID)
}

// GetPostLikes retrieves all likes for a specific post.
func (s *InteractionService) GetPostLikes(ctx context.Context, postID string) ([]*models.Profile, error) {
	return s.interactionRepo.GetPostLikes(ctx, postID)
}

// LikePost allows a user to like a post.
func (s *InteractionService) LikePost(ctx context.Context, userID, postID string) error {
	return s.interactionRepo.LikePost(ctx, userID, postID)
}

// UnlikePost allows a user to unlike a post.
func (s *InteractionService) UnlikePost(ctx context.Context, userID, postID string) error {
	return s.interactionRepo.UnlikePost(ctx, userID, postID)
}
