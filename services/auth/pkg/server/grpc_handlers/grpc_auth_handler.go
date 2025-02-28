package grpc_handler

import (
	"context"

	// "fmt"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
	proto "github.com/p4elkab35t/salyte_backend/services/auth/pkg/server/proto"
)

type AuthHandler struct {
	authLogic *logic.AuthLogicService
	proto.UnimplementedAuthServiceServer
}

func NewAuthHandler(authLogic *logic.AuthLogicService) *AuthHandler {
	return &AuthHandler{authLogic: authLogic}
}

func (h *AuthHandler) SignInCredentials(ctx context.Context, req *proto.SignInCredentialsRequest) (*proto.SignInResponse, error) {

	// fmt.Println("Sign in credentials")
	session, err := h.authLogic.SignIn(ctx, req.Email, req.Password)
	if err != nil {
		return &proto.SignInResponse{Status: 1}, err
	}
	if session == nil {
		return &proto.SignInResponse{Status: 1}, nil
	}

	// fmt.Println("User signed in")

	return &proto.SignInResponse{Token: session.Session_token, UserId: session.User_id, Status: 0}, nil
}

func (h *AuthHandler) SignInToken(ctx context.Context, req *proto.SignInTokenRequest) (*proto.SignInResponse, error) {
	session, err := h.authLogic.CheckToken(ctx, req.Token)
	if err != nil {
		return &proto.SignInResponse{Status: 1}, err
	}
	if session == nil {
		return &proto.SignInResponse{Status: 1}, nil
	}
	return &proto.SignInResponse{Token: session.Session_token, UserId: session.User_id, Status: 0}, nil
}

func (h *AuthHandler) SignUp(ctx context.Context, req *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	session, err := h.authLogic.SignUp(ctx, req.Email, req.Password)

	if err != nil {
		return &proto.SignUpResponse{Status: 1}, err
	}

	if session == nil {
		return &proto.SignUpResponse{Status: 1}, nil
	}

	// fmt.Println("User signed up")

	return &proto.SignUpResponse{UserId: session.User_id, Token: session.Session_token, Status: 0}, nil
}

func (h *AuthHandler) VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest) (*proto.VerifyTokenResponse, error) {
	valid, err := h.authLogic.VerifySession(ctx, req.Token, req.UserId)
	if err != nil {
		return &proto.VerifyTokenResponse{Status: 1}, err
	}
	return &proto.VerifyTokenResponse{IsValid: valid, Status: 0}, nil
}

func (h *AuthHandler) SignOut(ctx context.Context, req *proto.SignOutRequest) (*proto.SignOutResponse, error) {
	err := h.authLogic.SignOut(ctx, req.Token)
	if err != nil {
		return &proto.SignOutResponse{Status: 1}, err
	}
	return &proto.SignOutResponse{Status: 0}, nil
}
