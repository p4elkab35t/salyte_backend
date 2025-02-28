package repository

import (
	"context"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

type SecurityLogRepository interface {
	// Create new security log
	CreateSecurityLog(ctx context.Context, securityLog *models.SecurityLog) (*models.SecurityLog, error)
	// Get All security logs for user by user_id
	GetAllSecurityLogsByUserID(ctx context.Context, user_id string) ([]*models.SecurityLog, error)
	// Get security log by id
	GetSecurityLogByID(ctx context.Context, id string) (*models.SecurityLog, error)
}
