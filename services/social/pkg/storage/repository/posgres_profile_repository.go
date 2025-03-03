package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	// "github.com/p4elkab35t/salyte_backend/services/social/pkg/storage"
)

func NewPostgresProfileRepositorySQL(db *pgxpool.Pool) ProfileRepository {
	return &PostgresRepositorySQL{db: db}
}

func (r *PostgresRepositorySQL) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	query := `
		INSERT INTO Profile (user_id, username, bio, profile_picture_url)
		VALUES ($1, $2, $3, $4)
		RETURNING profile_id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, profile.UserID, profile.Username, profile.Bio, profile.ProfilePictureURL).
		Scan(&profile.ProfileID, &profile.CreatedAt, &profile.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}
	return profile, nil
}

func (r *PostgresRepositorySQL) GetProfileByID(ctx context.Context, id string) (*models.Profile, error) {
	query := `
		SELECT profile_id, user_id, username, bio, profile_picture_url, visibility, created_at, updated_at
		FROM Profile
		WHERE profile_id = $1
	`
	var profile models.Profile
	err := r.db.QueryRow(ctx, query, id).Scan(
		&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		// if errors.Is(err, pgx.ErrNoRows) {
		// 	return nil, nil
		// }
		return nil, fmt.Errorf("failed to get profile by ID: %w", err)
	}
	return &profile, nil
}

func (r *PostgresRepositorySQL) GetProfileByUserID(ctx context.Context, userID string) (*models.Profile, error) {
	query := `
		SELECT profile_id, user_id, username, bio, profile_picture_url, visibility, created_at, updated_at
		FROM Profile
		WHERE user_id = $1
	`
	var profile models.Profile
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		// if errors.Is(err, pgx.ErrNoRows) {
		// 	return nil, nil
		// }
		return nil, fmt.Errorf("failed to get profile by user ID: %w", err)
	}
	return &profile, nil
}

func (r *PostgresRepositorySQL) GetAllProfiles(ctx context.Context) ([]*models.Profile, error) {
	query := `
		SELECT profile_id, user_id, username, bio, profile_picture_url, visibility, created_at, updated_at
		FROM Profile
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all profiles: %w", err)
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(
			&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}

func (r *PostgresRepositorySQL) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	query := `
		UPDATE Profile
		SET username = $1, bio = $2, profile_picture_url = $3, visibility = $4, updated_at = $5
		WHERE profile_id = $6
	`
	_, err := r.db.Exec(ctx, query, profile.Username, profile.Bio, profile.ProfilePictureURL, profile.Visibility, time.Now(), profile.ProfileID)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) DeleteProfile(ctx context.Context, id string) error {
	query := `
		DELETE FROM Profile
		WHERE profile_id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) GetSettings(ctx context.Context, profileID string) (*models.Setting, error) {
	query := `
		SELECT setting_id, profile_id, dark_mode_enabled, language, created_at, updated_at
		FROM Settings
		WHERE profile_id = $1
	`
	var setting models.Setting
	err := r.db.QueryRow(ctx, query, profileID).Scan(
		&setting.SettingID, &setting.ProfileID, &setting.DarkModeEnabled, &setting.Language, &setting.CreatedAt, &setting.UpdatedAt,
	)
	if err != nil {
		// if errors.Is(err, pgx.ErrNoRows) {
		// 	return nil, nil
		// }
		return nil, fmt.Errorf("failed to get settings: %w", err)
	}
	return &setting, nil
}

func (r *PostgresRepositorySQL) UpdateSettings(ctx context.Context, settings *models.Setting) error {
	query := `
		UPDATE Settings
		SET dark_mode_enabled = $1, language = $2, updated_at = $3
		WHERE profile_id = $4
	`
	_, err := r.db.Exec(ctx, query, settings.DarkModeEnabled, settings.Language, time.Now(), settings.ProfileID)
	if err != nil {
		return fmt.Errorf("failed to update settings: %w", err)
	}
	return nil
}

func NewPostgresFollowRepositorySQL(db *pgxpool.Pool) FollowRepository {
	return &PostgresRepositorySQL{db: db}
}

func (r *PostgresRepositorySQL) GetProfileFollowers(ctx context.Context, profileID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM Followers f
		JOIN Profile p ON f.follower_profile_id = p.profile_id
		WHERE f.followed_profile_id = $1
	`
	rows, err := r.db.Query(ctx, query, profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile followers: %w", err)
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(
			&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}

func (r *PostgresRepositorySQL) GetProfileFollowing(ctx context.Context, profileID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM Followers f
		JOIN Profile p ON f.followed_profile_id = p.profile_id
		WHERE f.follower_profile_id = $1
	`
	rows, err := r.db.Query(ctx, query, profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile following: %w", err)
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(
			&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}

func (r *PostgresRepositorySQL) FollowProfile(ctx context.Context, userID, profileID string) error {
	query := `
		INSERT INTO Followers (follower_profile_id, followed_profile_id)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(ctx, query, userID, profileID)
	if err != nil {
		return fmt.Errorf("failed to follow profile: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) UnfollowProfile(ctx context.Context, userID, profileID string) error {
	query := `
		DELETE FROM Followers
		WHERE follower_profile_id = $1 AND followed_profile_id = $2
	`
	_, err := r.db.Exec(ctx, query, userID, profileID)
	if err != nil {
		return fmt.Errorf("failed to unfollow profile: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) MakeFriends(ctx context.Context, userID, profileID string) error {
	query := `
		INSERT INTO Interchange (profile_id, friend_profile_id, status)
		VALUES ($1, $2, 'accepted')
	`
	_, err := r.db.Exec(ctx, query, userID, profileID)
	if err != nil {
		return fmt.Errorf("failed to make friends: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) Unfriend(ctx context.Context, userID, profileID string) error {
	query := `
		DELETE FROM Interchange
		WHERE profile_id = $1 AND friend_profile_id = $2
	`
	_, err := r.db.Exec(ctx, query, userID, profileID)
	if err != nil {
		return fmt.Errorf("failed to unfriend: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) GetFriends(ctx context.Context, profileID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM Interchange i
		JOIN Profile p ON i.friend_profile_id = p.profile_id
		WHERE i.profile_id = $1 AND i.status = 'accepted'
	`
	rows, err := r.db.Query(ctx, query, profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get friends: %w", err)
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(
			&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}

func (r *PostgresRepositorySQL) GetFriendsRequests(ctx context.Context, profileID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM Interchange i
		JOIN Profile p ON i.profile_id = p.profile_id
		WHERE i.friend_profile_id = $1 AND i.status = 'pending'
	`
	rows, err := r.db.Query(ctx, query, profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to get friend requests: %w", err)
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(
			&profile.ProfileID, &profile.UserID, &profile.Username, &profile.Bio, &profile.ProfilePictureURL, &profile.Visibility, &profile.CreatedAt, &profile.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}
