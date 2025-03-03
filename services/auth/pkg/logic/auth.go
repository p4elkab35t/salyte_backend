package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	// "os/user"
	// "fmt"
	"time"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/storage/repository"
)

type AuthLogicService struct {
	UserRepo                  repository.UserRepository
	PostgresSessionRepository repository.SessionRepository
	Logger                    *SecurityLogLogicService
	//RedisSessionRepository    repository.SessionRepository
}

func NewAuthLogic(userRepo repository.UserRepository, sessionPostgresRepo repository.SessionRepository, logger *SecurityLogLogicService) *AuthLogicService { //, sessionRedisRepo repository.SessionRepository) *AuthLogicService {
	return &AuthLogicService{
		UserRepo:                  userRepo,
		PostgresSessionRepository: sessionPostgresRepo,
		Logger:                    logger,
		//RedisSessionRepository:    sessionRedisRepo,
	}
}

// for any of theese possible to return error, if something went wrong
// any new token is generated, returned to user and stored in redis and postgres
// signup user with credentials and sign him in right away with generating new token

func (s *AuthLogicService) SignUp(ctx context.Context, email, password string) (*models.Session, error) {
	hashedPassword, err := HashPassword([]byte(password))
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:         email,
		Password_hash: string(hashedPassword),
	}
	err = s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// TODO: Implement user creation hook to social service

	// fmt.Println("user created")
	// fmt.Println(user.User_id)
	session, err := s.SignIn(ctx, email, password)
	if err != nil {
		//s.Logger.CreateSecurityLog(ctx, user.User_id, "Initial user sign in fail")
		return nil, err
	}

	profileData := map[string]string{
		"userID": session.User_id,
		"email":  user.Email, // Ensure JSON field names match
	}

	jsonData, err := json.Marshal(profileData)
	if err != nil {
		return nil, err
	}

	hookUrl := "http://localhost:8081/social/profile"
	req, err := http.NewRequest("POST", hookUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Profile service returned error: %s", string(bodyBytes))
		return nil, errors.New("failed to create profile")
	}

	s.Logger.CreateSecurityLog(ctx, session.User_id, "User created")

	// fmt.Println("session created")
	// fmt.Println(session)

	return session, nil
}

// signin user with credentials with generating new token

func (s *AuthLogicService) SignIn(ctx context.Context, email, password string) (*models.Session, error) {

	// fmt.Println("signin")

	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// fmt.Println("user found")
	// fmt.Println(user.User_id)

	// MAIN AUTH FUNC CHECK PASSWORD HERE
	// IMPORTANT

	err = ComparePasswords([]byte(user.Password_hash), []byte(password))
	if err != nil {
		s.Logger.CreateSecurityLog(ctx, user.User_id, "User sign in attempt fail (bad password)")
		return nil, err
	}

	// fmt.Println("password correct")

	//session, err = s.RedisSessionRepository.CreateSession(ctx, session)
	session, err := s.createSession(ctx, user.User_id)
	if err != nil {
		s.Logger.CreateSecurityLog(ctx, user.User_id, "User sign in attempt fail (session creation fail)")
		return nil, err
	}

	s.Logger.CreateSecurityLog(ctx, user.User_id, "User sign in success")
	// fmt.Println("session created")
	// fmt.Println(session)

	return session, nil
}

// check user token and return user, if token is less than 1 hour, then regenerate token and return along with user
// first to check in redis, if not found then check in postgres

func (s *AuthLogicService) CheckToken(ctx context.Context, token string) (*models.Session, error) {
	//session, err := s.RedisSessionRepository.GetSessionByToken(ctx, token)
	//if err != nil {
	session, err := s.PostgresSessionRepository.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	//}

	if session.Expires_at.Before(time.Now()) {
		s.Logger.CreateSecurityLog(ctx, session.User_id, "User token accessed (expired)")
		return nil, errors.New("token expired")
	}

	if time.Until(session.Expires_at) < time.Hour {
		session, err = s.createSession(ctx, session.User_id)

		if err != nil {
			return nil, err
		}

		s.Logger.CreateSecurityLog(ctx, session.User_id, "User token accessed (renewed)")

	}

	user, err := s.UserRepo.GetUserById(ctx, session.User_id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	s.Logger.CreateSecurityLog(ctx, user.User_id, "User token accessed (valid)")

	return session, nil
}

// signout user and delete token from redis and set expire to 0 in postgres

func (s *AuthLogicService) SignOut(ctx context.Context, token string) error {
	//err := s.RedisSessionRepository.DeleteSession(ctx, token)
	//if err != nil {
	//	return err
	//}

	session, err := s.PostgresSessionRepository.GetSessionByToken(ctx, token)
	if err != nil {
		return err
	}

	session.Expires_at = time.Now()
	err = s.PostgresSessionRepository.UpdateSession(ctx, session)
	if err != nil {
		s.Logger.CreateSecurityLog(ctx, session.User_id, "User sign out (fail)")
		return err
	}

	s.Logger.CreateSecurityLog(ctx, session.User_id, "User sign out (success)")

	return nil
}

func (s *AuthLogicService) VerifySession(ctx context.Context, token string, userID string) (bool, error) {
	if token == "" || userID == "" {
		return false, errors.New("token and user_id are required")
	}

	//result, err := s.RedisSessionRepository.VerifySession(ctx, token, userID)
	result, err := s.PostgresSessionRepository.VerifySession(ctx, token, userID)
	if err != nil {
		s.Logger.CreateSecurityLog(ctx, userID, "User session verification fail")
		return false, err
	}

	return result, nil
}

func (s *AuthLogicService) createSession(ctx context.Context, user_id string) (*models.Session, error) {
	session := &models.Session{
		User_id:       user_id,
		Session_token: GenerateToken(),
		Expires_at:    time.Now().Add(time.Hour * 24),
	}

	session, err := s.PostgresSessionRepository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
