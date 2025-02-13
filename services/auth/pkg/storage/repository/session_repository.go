package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

type SessionRepository interface {
	// Verify session by token and user_email or user_id
	VerifySession(ctx context.Context, token string, user_id string) (bool, error)
	// Delete session by token
	DeleteSession(ctx context.Context, token string) error
	// Create new session
	CreateSession(ctx context.Context, session *models.Session) (*models.Session, error)
	// Get All session for user by user_id
	GetAllActiveSessionByUserID(ctx context.Context, user_id string) ([]*models.Session, error)
	// Get session by token
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)
	// Update session by token
	UpdateSession(ctx context.Context, session *models.Session) error
	// Delete all sessions for user by user_id
	DeleteAllSessionsByUserID(ctx context.Context, user_id string) error
}
