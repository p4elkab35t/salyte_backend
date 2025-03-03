package logic

import (
	"context"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

type ProfileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) *ProfileService {
	return &ProfileService{profileRepo: profileRepo}
}

// CreateProfile creates a new profile.
func (s *ProfileService) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	now := time.Now()
	profile.CreatedAt = &now
	profile.UpdatedAt = &now
	return s.profileRepo.CreateProfile(ctx, profile)
}

// GetProfileByID retrieves a profile by its ID.
func (s *ProfileService) GetProfileByID(ctx context.Context, id string) (*models.Profile, error) {
	return s.profileRepo.GetProfileByID(ctx, id)
}

// GetProfileByUserID retrieves a profile by its user ID.
func (s *ProfileService) GetProfileByUserID(ctx context.Context, userID string) (*models.Profile, error) {
	return s.profileRepo.GetProfileByUserID(ctx, userID)
}

// GetAllProfiles retrieves all profiles.
func (s *ProfileService) GetAllProfiles(ctx context.Context) ([]*models.Profile, error) {
	return s.profileRepo.GetAllProfiles(ctx)
}

// UpdateProfile updates an existing profile.
func (s *ProfileService) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	now := time.Now()
	profile.UpdatedAt = &now
	return s.profileRepo.UpdateProfile(ctx, profile)
}

// DeleteProfile deletes a profile by its ID.
func (s *ProfileService) DeleteProfile(ctx context.Context, id string) error {
	return s.profileRepo.DeleteProfile(ctx, id)
}

// GetSettings retrieves the settings for a profile.
func (s *ProfileService) GetSettings(ctx context.Context, profileID string) (*models.Setting, error) {
	return s.profileRepo.GetSettings(ctx, profileID)
}

// UpdateSettings updates the settings for a profile.
func (s *ProfileService) UpdateSettings(ctx context.Context, settings *models.Setting) error {
	now := time.Now()
	settings.UpdatedAt = now
	return s.profileRepo.UpdateSettings(ctx, settings)
}
