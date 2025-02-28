package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

type PostgresRepositorySQL struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepositorySQL(db *pgxpool.Pool) UserRepository {
	return &PostgresRepositorySQL{db: db}
}

// CheckCredentials(email, password string) (*models.User, error)

func (r *PostgresRepositorySQL) CheckCredentials(ctx context.Context, email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	row := r.db.QueryRow(ctx,
		"SELECT * FROM users WHERE email = $1 AND password_hash = $2",
		email, password)
	user := &models.User{}
	err := row.Scan(&user.User_id, &user.Email, &user.Password_hash, &user.Is_verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// 	UpdateCredentials(user *models.User) error

func (r *PostgresRepositorySQL) UpdateCredentials(ctx context.Context, user *models.User) error {

	if user.User_id == "" {
		return errors.New("user_id is required")
	}

	updates := []string{}
	args := []interface{}{}
	argPos := 1

	if user.Email != "" {
		updates = append(updates, fmt.Sprintf("email = $%d", argPos))
		args = append(args, user.Email)
		argPos++
	}
	if user.Password_hash != "" {
		updates = append(updates, fmt.Sprintf("password_hash = $%d", argPos))
		args = append(args, user.Password_hash)
		argPos++
	}

	if len(updates) == 0 {
		return errors.New("at least one field (email or password_hash) is required")
	}

	args = append(args, user.User_id)
	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id = $%d", strings.Join(updates, ", "), argPos)

	_, err := r.db.Exec(ctx, query, args...)

	return err
}

// 	CreateUser(user *models.User) error

func (r *PostgresRepositorySQL) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx,
		"INSERT INTO users (email, password_hash) VALUES ($1, $2)",
		user.Email, user.Password_hash)

	return err
}

// 	GetUserByEmail(email string) (*models.User, error)

func (r *PostgresRepositorySQL) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email are required")
	}
	row := r.db.QueryRow(ctx,
		"SELECT * FROM users WHERE email = $1",
		email)
	user := &models.User{}
	err := row.Scan(&user.User_id, &user.Email, &user.Password_hash, &user.Is_verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// 	GetUserById(id string) (*models.User, error)

func (r *PostgresRepositorySQL) GetUserById(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id are required")
	}
	row := r.db.QueryRow(ctx,
		"SELECT * FROM users WHERE user_id = $1",
		id)
	user := &models.User{}
	err := row.Scan(&user.User_id, &user.Email, &user.Password_hash, &user.Is_verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// 	GetUserByToken(token string) (*models.User, error)

func (r *PostgresRepositorySQL) GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	if token == "" {
		return nil, errors.New("token are required")
	}
	row := r.db.QueryRow(ctx,
		`SELECT users.* FROM users JOIN sessions 
		ON users.user_id = sessions.user_id 
		WHERE sessions.session_token = $1 
		AND (sessions.expires_at > NOW() 
		OR sessions.expires_at IS NULL)`,
		token)
	user := &models.User{}
	err := row.Scan(&user.User_id, &user.Email, &user.Password_hash, &user.Is_verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
