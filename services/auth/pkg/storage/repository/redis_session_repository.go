package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
)

type RedisSessionRepository struct {
	redis *redis.Client
}

func NewRedisSessionRepository(redis *redis.Client) *RedisSessionRepository {
	return &RedisSessionRepository{redis: redis}
}

// Create a new session and store it in Redis
func (r *RedisSessionRepository) CreateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	if session.Session_token == "" || session.User_id == "" {
		return nil, errors.New("session_token and user_id are required")
	}

	// Convert session struct to JSON
	sessionData, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	// Store session in Redis with expiration
	duration := time.Until(session.Expires_at)
	err = r.redis.Set(ctx, session.Session_token, sessionData, duration).Err()
	if err != nil {
		return nil, err
	}

	// Maintain a set of sessions per user
	err = r.redis.SAdd(ctx, "user_sessions:"+session.User_id, session.Session_token).Err()
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Verify session by token and user_id
func (r *RedisSessionRepository) VerifySession(ctx context.Context, token string, userID string) (bool, error) {
	if token == "" || userID == "" {
		return false, errors.New("token and user_id are required")
	}

	session, err := r.GetSessionByToken(ctx, token)
	if err != nil {
		return false, err
	}

	if session.User_id != userID {
		return false, nil
	}

	return true, nil
}

// Get session by token
func (r *RedisSessionRepository) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	if token == "" {
		return nil, errors.New("token is required")
	}

	sessionData, err := r.redis.Get(ctx, token).Result()
	if err == redis.Nil {
		return nil, errors.New("session not found")
	} else if err != nil {
		return nil, err
	}

	var session models.Session
	err = json.Unmarshal([]byte(sessionData), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// Get all active sessions for a user
func (r *RedisSessionRepository) GetAllActiveSessionsByUserID(ctx context.Context, userID string) ([]*models.Session, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}

	// Retrieve all session tokens for the user
	sessionTokens, err := r.redis.SMembers(ctx, "user_sessions:"+userID).Result()
	if err != nil {
		return nil, err
	}

	var sessions []*models.Session
	for _, token := range sessionTokens {
		session, err := r.GetSessionByToken(ctx, token)
		if err == nil {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

// Delete session by token
func (r *RedisSessionRepository) DeleteSession(ctx context.Context, token string) error {
	if token == "" {
		return errors.New("token is required")
	}

	// Retrieve session to get the user_id
	session, err := r.GetSessionByToken(ctx, token)
	if err != nil {
		return err
	}

	// Remove session from Redis
	err = r.redis.Del(ctx, token).Err()
	if err != nil {
		return err
	}

	// Remove session from user's session list
	err = r.redis.SRem(ctx, "user_sessions:"+session.User_id, token).Err()
	if err != nil {
		return err
	}

	return nil
}

// Delete all active sessions for a user (e.g., on logout from all devices)
func (r *RedisSessionRepository) DeleteAllSessionsByUserID(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("user_id is required")
	}

	// Get all session tokens for user
	sessionTokens, err := r.redis.SMembers(ctx, "user_sessions:"+userID).Result()
	if err != nil {
		return err
	}

	// Remove all sessions
	for _, token := range sessionTokens {
		r.redis.Del(ctx, token)
	}

	// Remove user session index
	r.redis.Del(ctx, "user_sessions:"+userID)

	return nil
}

// Update session expiration
func (r *RedisSessionRepository) UpdateSession(ctx context.Context, session *models.Session) error {
	if session.Session_token == "" || session.User_id == "" {
		return errors.New("session_token and user_id are required")
	}

	// Convert session struct to JSON
	sessionData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	// Store session in Redis with expiration
	duration := time.Until(session.Expires_at)
	err = r.redis.Set(ctx, session.Session_token, sessionData, duration).Err()
	if err != nil {
		return err
	}

	return nil
}
