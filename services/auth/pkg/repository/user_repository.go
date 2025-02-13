package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

type UserRepository interface {
	// Get user by email
	CheckCredentials(ctx context.Context, email, password string) (*models.User, error)
	// Update user credentials
	UpdateCredentials(ctx context.Context, user *models.User) error
	// Create new user
	CreateUser(ctx context.Context, user *models.User) error
	// Get user by email
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// Get user by id
	GetUserById(ctx context.Context, id string) (*models.User, error)
	// Get user by token
	GetUserByToken(ctx context.Context, token string) (*models.User, error)
}
