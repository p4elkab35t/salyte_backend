package logic

import (
	"context"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type CommunityService struct {
	communityRepo repository.CommunityRepository
}

func NewCommunityService(communityRepo repository.CommunityRepository) *CommunityService {
	return &CommunityService{communityRepo: communityRepo}
}

// CreateCommunity creates a new community.
func (s *CommunityService) CreateCommunity(ctx context.Context, community *models.Community) (*models.Community, error) {
	community.CreatedAt = time.Now()
	community.UpdatedAt = time.Now()
	return s.communityRepo.CreateCommunity(ctx, community)
}

// GetCommunityByID retrieves a community by its ID.
func (s *CommunityService) GetCommunityByID(ctx context.Context, id string) (*models.Community, error) {
	return s.communityRepo.GetCommunityByID(ctx, id)
}

// GetCommunitiesByName retrieves a community by its name.
func (s *CommunityService) GetCommunitiesByName(ctx context.Context, name string) (*models.Community, error) {
	return s.communityRepo.GetCommunitiesByName(ctx, name)
}

// GetAllCommunities retrieves all communities.
func (s *CommunityService) GetAllCommunities(ctx context.Context) ([]*models.Community, error) {
	return s.communityRepo.GetAllCommunities(ctx)
}

// UpdateCommunity updates an existing community.
func (s *CommunityService) UpdateCommunity(ctx context.Context, community *models.Community) error {
	community.UpdatedAt = time.Now()
	return s.communityRepo.UpdateCommunity(ctx, community)
}

// GetCommunityFollowers retrieves all followers of a community.
func (s *CommunityService) GetCommunityFollowers(ctx context.Context, communityID string) ([]*models.Profile, error) {
	return s.communityRepo.GetCommunityFollowers(ctx, communityID)
}

// GetCommunitiesByUserID retrieves all communities a user is a member of.
func (s *CommunityService) GetCommunitiesByUserID(ctx context.Context, userID string) ([]*models.Community, error) {
	return s.communityRepo.GetCommunitiesByUserID(ctx, userID)
}

// FollowCommunity allows a user to follow a community.
func (s *CommunityService) FollowCommunity(ctx context.Context, userID, communityID string) error {
	return s.communityRepo.FollowCommunity(ctx, userID, communityID)
}

// UnfollowCommunity allows a user to unfollow a community.
func (s *CommunityService) UnfollowCommunity(ctx context.Context, userID, communityID string) error {
	return s.communityRepo.UnfollowCommunity(ctx, userID, communityID)
}
