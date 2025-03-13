package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

type PostRepository interface {
	// Create new post
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	// Get post by id
	GetPostByID(ctx context.Context, id string) (*models.Post, error)
	// Get all posts
	GetAllPosts(ctx context.Context, page, limit int) ([]*models.Post, error)
	// Update post by id
	UpdatePost(ctx context.Context, post *models.Post) error
	// Delete post by id
	DeletePost(ctx context.Context, id string) error
	// Get posts by user_id
	GetPostsByUserID(ctx context.Context, user_id string) ([]*models.Post, error)
	// Get posts by community_id
	GetPostsByCommunityID(ctx context.Context, community_id string) ([]*models.Post, error)
}

type CommentRepository interface {
	// Create new comment
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	// Get comment by id
	GetCommentByID(ctx context.Context, id string) (*models.Comment, error)
	// Get all comments
	GetAllComments(ctx context.Context) ([]*models.Comment, error)
	// Update comment by id
	UpdateComment(ctx context.Context, comment *models.Comment) error
	// Delete comment by id
	DeleteComment(ctx context.Context, id string) error
	// Get comment by post_id
	GetCommentsByPostID(ctx context.Context, post_id string) ([]*models.Comment, error)
}

type InteractionRepository interface {
	// Get post comments
	GetPostComments(ctx context.Context, post_id string) ([]*models.Comment, error)
	// Get post likes
	GetPostLikes(ctx context.Context, post_id string) ([]*models.Profile, error)
	// Like post
	LikePost(ctx context.Context, user_id, post_id string) error
	// Unlike post
	UnlikePost(ctx context.Context, user_id, post_id string) error
}
