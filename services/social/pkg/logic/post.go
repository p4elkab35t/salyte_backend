package logic

import (
	"context"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type PostService struct {
	postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

// CreatePost creates a new post.
func (s *PostService) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	newTime := time.Now()
	post.CreatedAt = &newTime
	post.UpdatedAt = &newTime
	return s.postRepo.CreatePost(ctx, post)
}

// GetPostByID retrieves a post by its ID.
func (s *PostService) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	return s.postRepo.GetPostByID(ctx, id)
}

// GetAllPosts retrieves all posts.
func (s *PostService) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	return s.postRepo.GetAllPosts(ctx)
}

// UpdatePost updates an existing post.
func (s *PostService) UpdatePost(ctx context.Context, post *models.Post) error {
	newTime := time.Now()
	post.UpdatedAt = &newTime
	return s.postRepo.UpdatePost(ctx, post)
}

// DeletePost deletes a post by its ID.
func (s *PostService) DeletePost(ctx context.Context, id string) error {
	return s.postRepo.DeletePost(ctx, id)
}

// GetPostsByUserID retrieves all posts by a specific user.
func (s *PostService) GetPostsByUserID(ctx context.Context, userID string) ([]*models.Post, error) {
	return s.postRepo.GetPostsByUserID(ctx, userID)
}

// GetPostsByCommunityID retrieves all posts in a specific community.
func (s *PostService) GetPostsByCommunityID(ctx context.Context, communityID string) ([]*models.Post, error) {
	return s.postRepo.GetPostsByCommunityID(ctx, communityID)
}
