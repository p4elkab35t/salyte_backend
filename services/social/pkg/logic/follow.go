package logic

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type FollowService struct {
	followRepo repository.FollowRepository
}

func NewFollowService(followRepo repository.FollowRepository) *FollowService {
	return &FollowService{followRepo: followRepo}
}

// GetProfileFollowers retrieves all followers of a profile.
func (s *FollowService) GetProfileFollowers(ctx context.Context, profileID string) ([]*models.Profile, error) {
	return s.followRepo.GetProfileFollowers(ctx, profileID)
}

// GetProfileFollowing retrieves all profiles a user is following.
func (s *FollowService) GetProfileFollowing(ctx context.Context, profileID string) ([]*models.Profile, error) {
	return s.followRepo.GetProfileFollowing(ctx, profileID)
}

// FollowProfile allows a user to follow another profile.
func (s *FollowService) FollowProfile(ctx context.Context, userID, profileID string) error {
	return s.followRepo.FollowProfile(ctx, userID, profileID)
}

// UnfollowProfile allows a user to unfollow another profile.
func (s *FollowService) UnfollowProfile(ctx context.Context, userID, profileID string) error {
	return s.followRepo.UnfollowProfile(ctx, userID, profileID)
}

// MakeFriends allows two profiles to become friends.
func (s *FollowService) MakeFriends(ctx context.Context, userID, profileID string) error {
	return s.followRepo.MakeFriends(ctx, userID, profileID)
}

// Unfriend removes a friendship between two profiles.
func (s *FollowService) Unfriend(ctx context.Context, userID, profileID string) error {
	return s.followRepo.Unfriend(ctx, userID, profileID)
}

// GetFriends retrieves all friends of a profile.
func (s *FollowService) GetFriends(ctx context.Context, profileID string) ([]*models.Profile, error) {
	return s.followRepo.GetFriends(ctx, profileID)
}

// GetFriendsRequests retrieves all pending friend requests for a profile.
func (s *FollowService) GetFriendsRequests(ctx context.Context, profileID string) ([]*models.Profile, error) {
	return s.followRepo.GetFriendsRequests(ctx, profileID)
}
