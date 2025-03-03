package repository

import (
	"context"
	// "errors"
	"fmt"
	"time"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
)

func NewPostgresPostRepositorySQL(db *pgxpool.Pool) PostRepository {
	return &PostgresRepositorySQL{db: db}
}

func NewPostgresInteractionRepositorySQL(db *pgxpool.Pool) InteractionRepository {
	return &PostgresRepositorySQL{db: db}
}

func NewPostgresCommentRepositorySQL(db *pgxpool.Pool) CommentRepository {
	return &PostgresRepositorySQL{db: db}
}

func (r *PostgresRepositorySQL) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	fmt.Println(post.ProfileID)
	query := `
		INSERT INTO Posts (profile_id, community_id, content, media_url, visibility)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING post_id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, post.ProfileID, post.CommunityID, post.Content, post.MediaURL, post.Visibility).
		Scan(&post.PostID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return post, nil
}

func (r *PostgresRepositorySQL) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	query := `
		SELECT post_id, profile_id, community_id, content, media_url, visibility, created_at, updated_at
		FROM Posts
		WHERE post_id = $1
	`
	var post models.Post
	err := r.db.QueryRow(ctx, query, id).Scan(
		&post.PostID, &post.ProfileID, &post.CommunityID, &post.Content, &post.MediaURL, &post.Visibility, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		// if errors.Is(err, pgx.ErrNoRows) {
		// 	return nil, nil
		// }
		return nil, fmt.Errorf("failed to get post by ID: %w", err)
	}
	return &post, nil
}

func (r *PostgresRepositorySQL) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	query := `
		SELECT post_id, profile_id, community_id, content, media_url, visibility, created_at, updated_at
		FROM Posts
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all posts: %w", err)
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.PostID, &post.ProfileID, &post.CommunityID, &post.Content, &post.MediaURL, &post.Visibility, &post.CreatedAt, &post.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostgresRepositorySQL) UpdatePost(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE Posts
		SET content = $1, media_url = $2, visibility = $3, updated_at = $4
		WHERE post_id = $5
	`
	_, err := r.db.Exec(ctx, query, post.Content, post.MediaURL, post.Visibility, time.Now(), post.PostID)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) DeletePost(ctx context.Context, id string) error {
	query := `
		DELETE FROM Posts
		WHERE post_id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) GetPostsByUserID(ctx context.Context, userID string) ([]*models.Post, error) {
	query := `
		SELECT post_id, profile_id, community_id, content, media_url, visibility, created_at, updated_at
		FROM Posts
		WHERE profile_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts by user ID: %w", err)
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.PostID, &post.ProfileID, &post.CommunityID, &post.Content, &post.MediaURL, &post.Visibility, &post.CreatedAt, &post.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostgresRepositorySQL) GetPostsByCommunityID(ctx context.Context, communityID string) ([]*models.Post, error) {
	query := `
		SELECT post_id, profile_id, community_id, content, media_url, visibility, created_at, updated_at
		FROM Posts
		WHERE community_id = $1
	`
	rows, err := r.db.Query(ctx, query, communityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts by community ID: %w", err)
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.PostID, &post.ProfileID, &post.CommunityID, &post.Content, &post.MediaURL, &post.Visibility, &post.CreatedAt, &post.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostgresRepositorySQL) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	query := `
		INSERT INTO Comments (profile_id, post_id, content)
		VALUES ($1, $2, $3)
		RETURNING comment_id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query, comment.ProfileID, comment.PostID, comment.Content).
		Scan(&comment.CommentID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}
	return comment, nil
}

func (r *PostgresRepositorySQL) GetCommentByID(ctx context.Context, id string) (*models.Comment, error) {
	query := `
		SELECT comment_id, profile_id, post_id, content, created_at, updated_at
		FROM Comments
		WHERE comment_id = $1
	`
	var comment models.Comment
	err := r.db.QueryRow(ctx, query, id).Scan(
		&comment.CommentID, &comment.ProfileID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
	)
	if err != nil {
		// if errors.Is(err, pgx.ErrNoRows) {
		// 	return nil, nil
		// }
		return nil, fmt.Errorf("failed to get comment by ID: %w", err)
	}
	return &comment, nil
}

func (r *PostgresRepositorySQL) GetAllComments(ctx context.Context) ([]*models.Comment, error) {
	query := `
		SELECT comment_id, profile_id, post_id, content, created_at, updated_at
		FROM Comments
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all comments: %w", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.CommentID, &comment.ProfileID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (r *PostgresRepositorySQL) UpdateComment(ctx context.Context, comment *models.Comment) error {
	query := `
		UPDATE Comments
		SET content = $1, updated_at = $2
		WHERE comment_id = $3
	`
	_, err := r.db.Exec(ctx, query, comment.Content, time.Now(), comment.CommentID)
	if err != nil {
		return fmt.Errorf("failed to update comment: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) DeleteComment(ctx context.Context, id string) error {
	query := `
		DELETE FROM Comments
		WHERE comment_id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.Comment, error) {
	query := `
		SELECT comment_id, profile_id, post_id, content, created_at, updated_at
		FROM Comments
		WHERE post_id = $1
	`
	rows, err := r.db.Query(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments by post ID: %w", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.CommentID, &comment.ProfileID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (r *PostgresRepositorySQL) GetPostComments(ctx context.Context, postID string) ([]*models.Comment, error) {
	query := `
		SELECT comment_id, profile_id, post_id, content, created_at, updated_at
		FROM Comments
		WHERE post_id = $1
	`
	rows, err := r.db.Query(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post comments: %w", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.CommentID, &comment.ProfileID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (r *PostgresRepositorySQL) GetPostLikes(ctx context.Context, postID string) ([]*models.Profile, error) {
	query := `
		SELECT p.profile_id, p.user_id, p.username, p.bio, p.profile_picture_url, p.visibility, p.created_at, p.updated_at
		FROM Likes l
		JOIN Profile p ON l.profile_id = p.profile_id
		WHERE l.post_id = $1
	`
	rows, err := r.db.Query(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post likes: %w", err)
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

func (r *PostgresRepositorySQL) LikePost(ctx context.Context, userID, postID string) error {
	query := `
		INSERT INTO Likes (profile_id, post_id)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(ctx, query, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}
	return nil
}

func (r *PostgresRepositorySQL) UnlikePost(ctx context.Context, userID, postID string) error {
	query := `
		DELETE FROM Likes
		WHERE profile_id = $1 AND post_id = $2
	`
	_, err := r.db.Exec(ctx, query, userID, postID)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}
	return nil
}
