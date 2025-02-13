package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

func NewPostgresSessionRepositorySQL(db *pgxpool.Pool) SessionRepository {
	return &PostgresRepositorySQL{db: db}
}

func (r *PostgresRepositorySQL) VerifySession(ctx context.Context, token string, user_id string) (bool, error) {
	if token == "" || user_id == "" {
		return false, errors.New("token and user_id are required")
	}
	session, err := r.GetSessionByToken(ctx, token)
	if err != nil {
		return false, err
	}
	if session.User_id != user_id {
		return false, nil
	}
	return true, nil
}

func (r *PostgresRepositorySQL) DeleteSession(ctx context.Context, token string) error {
	if token == "" {
		return errors.New("token is required")
	}
	_, err := r.db.Exec(ctx, "DELETE FROM sessions WHERE session_token = $1", token)

	return err
}

func (r *PostgresRepositorySQL) CreateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	if session.Session_token == "" || session.User_id == "" {
		return nil, errors.New("token and user_id are required")
	}
	_, err := r.db.Exec(ctx, "INSERT INTO sessions (session_token, user_id, expires_at) VALUES ($1, $2, $3)", session.Session_token, session.User_id, session.Expires_at)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *PostgresRepositorySQL) GetAllActiveSessionByUserID(ctx context.Context, user_id string) ([]*models.Session, error) {
	if user_id == "" {
		return nil, errors.New("user_id is required")
	}
	rows, err := r.db.Query(ctx, "SELECT * FROM sessions WHERE user_id = $1 AND expires_at > NOW()", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sessions := make([]*models.Session, 0)
	for rows.Next() {
		session := &models.Session{}
		err := rows.Scan(&session.Session_token, &session.User_id, &session.Expires_at)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (r *PostgresRepositorySQL) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	if token == "" {
		return nil, errors.New("token is required")
	}
	row := r.db.QueryRow(ctx, "SELECT * FROM sessions WHERE session_token = $1 AND expires_at > NOW()", token)
	session := &models.Session{}
	err := row.Scan(&session.Session_id, &session.User_id, &session.Session_token, &session.Expires_at, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *PostgresRepositorySQL) UpdateSession(ctx context.Context, session *models.Session) error {
	if session.Session_token == "" || session.User_id == "" {
		return errors.New("token and user_id are required")
	}
	_, err := r.db.Exec(ctx, "UPDATE sessions SET expires_at = $1 WHERE session_token = $2", session.Expires_at, session.Session_token)
	return err
}

func (r *PostgresRepositorySQL) DeleteAllSessionsByUserID(ctx context.Context, user_id string) error {
	if user_id == "" {
		return errors.New("user_id is required")
	}
	_, err := r.db.Exec(ctx, "DELETE FROM sessions WHERE user_id = $1", user_id)
	return err
}
