package repository

import (
	"context"
	"errors"

	// "fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

func NewPostgresSecurityLogRepositorySQL(db *pgxpool.Pool) SecurityLogRepository {
	return &PostgresRepositorySQL{db: db}
}

func (r *PostgresRepositorySQL) CreateSecurityLog(ctx context.Context, securityLog *models.SecurityLog) (*models.SecurityLog, error) {
	if securityLog.User_id == "" || securityLog.Action == "" {
		return nil, errors.New("user_id and action are required")
	}
	_, err := r.db.Exec(ctx, "INSERT INTO security_logs (user_id, action, timestamp) VALUES ($1, $2, $3)", securityLog.User_id, securityLog.Action, securityLog.Timestamp)
	if err != nil {
		return nil, err
	}
	return securityLog, nil
}

func (r *PostgresRepositorySQL) GetAllSecurityLogsByUserID(ctx context.Context, user_id string) ([]*models.SecurityLog, error) {
	if user_id == "" {
		return nil, errors.New("user_id is required")
	}

	// fmt.Println("user_id: ", user_id)

	rows, err := r.db.Query(ctx, "SELECT * FROM security_logs WHERE user_id = $1", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// fmt.Println("rows: ", rows)

	securityLogs := make([]*models.SecurityLog, 0)
	for rows.Next() {
		securityLog := &models.SecurityLog{}
		err := rows.Scan(&securityLog.Log_id, &securityLog.User_id, &securityLog.Action, &securityLog.Ip_address, &securityLog.Timestamp)
		if err != nil {
			// fmt.Println("error: ", err)
			return nil, err
		}
		securityLogs = append(securityLogs, securityLog)
		// fmt.Println("securityLogs: ", securityLogs)
	}
	return securityLogs, nil
}

func (r *PostgresRepositorySQL) GetSecurityLogByID(ctx context.Context, id string) (*models.SecurityLog, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	securityLog := &models.SecurityLog{}
	row := r.db.QueryRow(ctx, "SELECT * FROM security_logs WHERE log_id = $1", id)
	err := row.Scan(&securityLog.Log_id, &securityLog.User_id, &securityLog.Action, &securityLog.Ip_address, &securityLog.Timestamp)
	if err != nil {
		return nil, err
	}
	return securityLog, nil
}
