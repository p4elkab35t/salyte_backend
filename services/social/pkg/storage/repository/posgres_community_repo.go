package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

func NewPostgresCommunityRepositorySQL(db *pgxpool.Pool) CommunityRepository {
	return &PostgresRepositorySQL{db: db}
}

// CreateCommunity creates a new community in the database.
func (r *PostgresRepositorySQL) CreateCommunity(ctx context.Context, community *models.Community) (*models.Community, error) {
	query := `
		INSERT INTO Communities (name, description, profile_picture_url, visibility)
		VALUES ($1, $2, $3, $4)
		RETURNING community_id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, community.Name, community.Description, community.ProfilePictureURL, community.Visibility).
		Scan(&community.CommunityID, &community.CreatedAt, &community.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create community: %w", err)
	}
	return community, nil
}

// GetCommunityByID retrieves a community by its ID.
func (r *PostgresRepositorySQL) GetCommunityByID(ctx context.Context, id string) (*models.Community, error) {
	query := `
		SELECT community_id, name, description, profile_picture_url, visibility, created_at, updated_at
		FROM Communities
		WHERE community_id = $1
	`
	var community models.Community
	err := r.db.QueryRow(ctx, query, id).Scan(
		&community.CommunityID, &community.Name, &community.Description, &community.ProfilePictureURL, &community.Visibility, &community.CreatedAt, &community.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get community by ID: %w", err)
	}
	return &community, nil
}

// GetCommunitiesByName retrieves a community by its name.
func (r *PostgresRepositorySQL) GetCommunitiesByName(ctx context.Context, name string) (*models.Community, error) {
	query := `
		SELECT community_id, name, description, profile_picture_url, visibility, created_at, updated_at
		FROM Communities
		WHERE name = $1
	`
	var community models.Community
	err := r.db.QueryRow(ctx, query, name).Scan(
		&community.CommunityID, &community.Name, &community.Description, &community.ProfilePictureURL, &community.Visibility, &community.CreatedAt, &community.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get community by name: %w", err)
	}
	return &community, nil
}

// GetAllCommunities retrieves all communities from the database.
func (r *PostgresRepositorySQL) GetAllCommunities(ctx context.Context) ([]*models.Community, error) {
	query := `
		SELECT community_id, name, description, profile_picture_url, visibility, created_at, updated_at
		FROM Communities
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all communities: %w", err)
	}
	defer rows.Close()

	var communities []*models.Community
	for rows.Next() {
		var community models.Community
		if err := rows.Scan(
			&community.CommunityID, &community.Name, &community.Description, &community.ProfilePictureURL, &community.Visibility, &community.CreatedAt, &community.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan community: %w", err)
		}
		communities = append(communities, &community)
	}
	return communities, nil
}

// UpdateCommunity updates an existing community in the database.
func (r *PostgresRepositorySQL) UpdateCommunity(ctx context.Context, community *models.Community) error {
	query := `
		UPDATE Communities
		SET name = $1, description = $2, profile_picture_url = $3, visibility = $4, updated_at = $5
		WHERE community_id = $6
	`
	_, err := r.db.Exec(ctx, query, community.Name, community.Description, community.ProfilePictureURL, community.Visibility, time.Now(), community.CommunityID)
	if err != nil {
		return fmt.Errorf("failed to update community: %w", err)
	}
	return nil
}

// GetCommunityFollowers retrieves all followers of a community.
func (r *PostgresRepositorySQL) GetCommunityFollowers(ctx context.Context, communityID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM CommunityMembers cm
		JOIN Profile p ON cm.profile_id = p.profile_id
		WHERE cm.community_id = $1
	`
	rows, err := r.db.Query(ctx, query, communityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get community followers: %w", err)
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

// GetCommunitiesByUserID retrieves all communities a user is a member of.
func (r *PostgresRepositorySQL) GetCommunitiesByUserID(ctx context.Context, userID string) ([]*models.Community, error) {
	query := `
		SELECT c.community_id, c.name, c.description, c.profile_picture_url, c.visibility, c.created_at, c.updated_at
		FROM CommunityMembers cm
		JOIN Communities c ON cm.community_id = c.community_id
		WHERE cm.profile_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get communities by user ID: %w", err)
	}
	defer rows.Close()

	var communities []*models.Community
	for rows.Next() {
		var community models.Community
		if err := rows.Scan(
			&community.CommunityID, &community.Name, &community.Description, &community.ProfilePictureURL, &community.Visibility, &community.CreatedAt, &community.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan community: %w", err)
		}
		communities = append(communities, &community)
	}
	return communities, nil
}

// FollowCommunity allows a user to follow a community.
func (r *PostgresRepositorySQL) FollowCommunity(ctx context.Context, userID, communityID string) error {
	query := `
		INSERT INTO CommunityMembers (community_id, profile_id, role)
		VALUES ($1, $2, 'member')
	`
	_, err := r.db.Exec(ctx, query, communityID, userID)
	if err != nil {
		return fmt.Errorf("failed to follow community: %w", err)
	}
	return nil
}

// UnfollowCommunity allows a user to unfollow a community.
func (r *PostgresRepositorySQL) UnfollowCommunity(ctx context.Context, userID, communityID string) error {
	query := `
		DELETE FROM CommunityMembers
		WHERE community_id = $1 AND profile_id = $2
	`
	_, err := r.db.Exec(ctx, query, communityID, userID)
	if err != nil {
		return fmt.Errorf("failed to unfollow community: %w", err)
	}
	return nil
}
