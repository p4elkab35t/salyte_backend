package grpc_handlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
	models "github.com/p4elkab35t/salyte_backend/services/social/pkg/models"
	proto "github.com/p4elkab35t/salyte_backend/services/social/pkg/server/proto"
)

type ProfileHandler struct {
	profileLogic *logic.ProfileService
	proto.UnimplementedSocialServiceServer
}

func NewProfileHandler(profileLogic *logic.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileLogic: profileLogic}
}

func (h *ProfileHandler) CreateProfile(ctx context.Context, req *proto.CreateUserProfileRequest) (*proto.CreateUserProfileResponse, error) {
	UserIdParsedUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return &proto.CreateUserProfileResponse{Status: 1}, err
	}

	profile := models.Profile{
		UserID:   UserIdParsedUUID,
		Username: req.Email,
	}

	newProfile, err := h.profileLogic.CreateProfile(ctx, &profile)
	if err != nil {
		return &proto.CreateUserProfileResponse{Status: 1}, err
	}

	return &proto.CreateUserProfileResponse{ProfileId: newProfile.UserID.String(), Status: 0}, nil
}
