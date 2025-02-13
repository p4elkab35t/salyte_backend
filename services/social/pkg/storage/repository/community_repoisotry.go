package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

type CommunityRepository interface {
	// Create new community
	CreateCommunity(ctx context.Context, community *models.Community) (*models.Community, error)
	// Get community by id
	GetCommunityByID(ctx context.Context, id string) (*models.Community, error)
	// Get community by name
	GetCommunitiesByName(ctx context.Context, name string) (*models.Community, error)
	// Get all communities
	GetAllCommunities(ctx context.Context) ([]*models.Community, error)
	// Update community by id
	UpdateCommunity(ctx context.Context, community *models.Community) error
	// Get community followers
	GetCommunityFollowers(ctx context.Context, community_id string) ([]*models.Profile, error)
	// Get communities by user_id
	GetCommunitiesByUserID(ctx context.Context, user_id string) ([]*models.Community, error)
	// Follow community
	FollowCommunity(ctx context.Context, user_id, community_id string) error
	// Unfollow community
	UnfollowCommunity(ctx context.Context, user_id, community_id string) error
}
