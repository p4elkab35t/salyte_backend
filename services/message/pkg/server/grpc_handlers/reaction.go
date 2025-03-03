package grpc_handler

import (
	"context"

	// "fmt"
	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
	proto "github.com/p4elkab35t/salyte_backend/services/message/pkg/server/proto"
)

type ReactionHandler struct {
	reactionService *logic.ReactionService
	proto.UnimplementedMessagingServiceServer
}

func NewReactionHandler(reactionService *logic.ReactionService) *ReactionHandler {
	return &ReactionHandler{reactionService: reactionService}
}

func (h *ReactionHandler) AddReaction(ctx context.Context, req *proto.AddReactionRequest) (*proto.AddReactionResponse, error) {
	reaction := &models.Reaction{
		MessageID: uuid.MustParse(req.MessageId),
		UserID:    uuid.MustParse(req.UserId),
		Emoji:     req.Reaction,
	}

	_, err := h.reactionService.ApplyReaction(ctx, reaction.MessageID, reaction.UserID, reaction)
	if err != nil {
		return &proto.AddReactionResponse{Success: false}, err
	}

	return &proto.AddReactionResponse{Success: true}, nil
}

func (h *ReactionHandler) RemoveReaction(ctx context.Context, req *proto.RemoveReactionRequest) (*proto.RemoveReactionResponse, error) {
	reaction := &models.Reaction{
		MessageID: uuid.MustParse(req.MessageId),
		UserID:    uuid.MustParse(req.UserId),
		Emoji:     req.Reaction,
	}

	err := h.reactionService.RemoveReaction(ctx, reaction.MessageID, reaction.UserID)
	if err != nil {
		return &proto.RemoveReactionResponse{Success: false}, err
	}

	return &proto.RemoveReactionResponse{Success: true}, nil
}
