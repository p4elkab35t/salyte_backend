package logic

import (
	"context"
	"errors"

	// "os/user"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/storage/repository"
)

type AuthLogicService struct {
	UserRepo                  repository.UserRepository
	PostgresSessionRepository repository.SessionRepository
	RedisSessionRepository    repository.SessionRepository
}

func NewAuthLogic(userRepo repository.UserRepository, sessionPostgresRepo repository.SessionRepository, sessionRedisRepo repository.SessionRepository) *AuthLogicService {
	return &AuthLogicService{
		UserRepo:                  userRepo,
		PostgresSessionRepository: sessionPostgresRepo,
		RedisSessionRepository:    sessionRedisRepo,
	}
}

// for any of theese possible to return error, if something went wrong
// any new token is generated, returned to user and stored in redis and postgres
// signup user with credentials and sign him in right away with generating new token

func (s *AuthLogicService) SignUp(ctx context.Context, email, password string) (*models.User, error) {
	user := &models.User{
		Email: email,
	}
	err := s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		User_id:       user.User_id,
		Session_token: GenerateToken(),
		Expires_at:    time.Now().Add(time.Hour * 24),
	}
	session, err = s.RedisSessionRepository.CreateSession(ctx, session)
	session, err = s.PostgresSessionRepository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// signin user with credentials with generating new token

func (s *AuthLogicService) SignIn(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.UserRepo.CheckCredentials(ctx, email, password)
	if err != nil {
		return nil, err
	}

	session := &models.Session{
		User_id:       user.User_id,
		Session_token: GenerateToken(),
		Expires_at:    time.Now().Add(time.Hour * 24),
	}
	session, err = s.RedisSessionRepository.CreateSession(ctx, session)
	session, err = s.PostgresSessionRepository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// check user token and return user, if token is less than 1 hour, then regenerate token and return along with user
// first to check in redis, if not found then check in postgres

func (s *AuthLogicService) CheckToken(ctx context.Context, token string) (*models.User, error) {
	session, err := s.RedisSessionRepository.GetSessionByToken(ctx, token)
	if err != nil {
		session, err = s.PostgresSessionRepository.GetSessionByToken(ctx, token)
		if err != nil {
			return nil, err
		}
	}

	if session.Expires_at.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	if time.Until(session.Expires_at) < time.Hour {
		session := &models.Session{
			User_id:       session.User_id,
			Session_token: GenerateToken(),
			Expires_at:    time.Now().Add(time.Hour * 24),
		}
		session, err = s.RedisSessionRepository.CreateSession(ctx, session)
		session, err = s.PostgresSessionRepository.CreateSession(ctx, session)
		if err != nil {
			return nil, err
		}
	}

	user, err := s.UserRepo.GetUserById(ctx, session.User_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// signout user and delete token from redis and set expire to 0 in postgres

func (s *AuthLogicService) SignOut(ctx context.Context, token string) error {
	err := s.RedisSessionRepository.DeleteSession(ctx, token)
	if err != nil {
		return err
	}

	session, err := s.PostgresSessionRepository.GetSessionByToken(ctx, token)
	if err != nil {
		return err
	}

	session.Expires_at = time.Now()
	err = s.PostgresSessionRepository.UpdateSession(ctx, session)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthLogicService) VerifySession(ctx context.Context, token string, userID string) (bool, error) {
	if token == "" || userID == "" {
		return false, errors.New("token and user_id are required")
	}

	result, err := s.RedisSessionRepository.VerifySession(ctx, token, userID)
	if err != nil {
		return false, err
	}

	return result, nil
}
