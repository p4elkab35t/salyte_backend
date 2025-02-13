package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

type ProfileRepository interface {
	// Create new profile
	CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	// Get profile by id
	GetProfileByID(ctx context.Context, id string) (*models.Profile, error)
	// Get profile by user_id
	GetProfileByUserID(ctx context.Context, user_id string) (*models.Profile, error)
	// Get all profiles
	GetAllProfiles(ctx context.Context) ([]*models.Profile, error)
	// Update profile by id
	UpdateProfile(ctx context.Context, profile *models.Profile) error
	// Delete profile by id
	DeleteProfile(ctx context.Context, id string) error
	// Get settings
	GetSettings(ctx context.Context, profile_id string) (*models.Setting, error)
	// Update settings
	UpdateSettings(ctx context.Context, settings *models.Setting) error
}

type FollowRepository interface {
	// Get profile followers
	GetProfileFollowers(ctx context.Context, profile_id string) ([]*models.Profile, error)
	// Get profile following
	GetProfileFollowing(ctx context.Context, profile_id string) ([]*models.Profile, error)
	// Follow profile
	FollowProfile(ctx context.Context, user_id, profile_id string) error
	// Unfollow profile
	UnfollowProfile(ctx context.Context, user_id, profile_id string) error
	// Make friends
	MakeFriends(ctx context.Context, user_id, profile_id string) error
	// Unfriend
	Unfriend(ctx context.Context, user_id, profile_id string) error
	// Get friends
	GetFriends(ctx context.Context, profile_id string) ([]*models.Profile, error)
	// Get friends requests
	GetFriendsRequests(ctx context.Context, profile_id string) ([]*models.Profile, error)
}
