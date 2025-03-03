package logic

import (
	"context"
	"errors"

	// "os/user"
	// "fmt"
	"time"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/storage/repository"
)

type ReactionService struct {
	messageRepo repository.MessageRepository
}

func NewReactionService(messageRepo repository.MessageRepository) *ReactionService {
	return &ReactionService{messageRepo: messageRepo}
}

// Get all reactions for message
func (s *ReactionService) GetReactionsByMessageID(ctx context.Context, messageID, userID uuid.UUID) ([]*models.Reaction, error) {
	// security mesaures
	// get chat by chat ID, get chat members by chat ID, check if user is in chat members
	message, err := s.messageRepo.GetMessageByID(ctx, messageID)
	if err != nil {
		return nil, err
	}
	chat, err := s.messageRepo.GetChatByMessageID(ctx, message.ID)
	if err != nil {
		return nil, err
	}
	chatMembers, err := s.messageRepo.GetChatMembers(ctx, chat.ID)
	if err != nil {
		return nil, err
	}
	if len(chatMembers) == 0 {
		return nil, errors.New("chat not found")
	}
	for _, member := range chatMembers {
		if member == userID {
			return s.messageRepo.GetReactionsByMessageID(ctx, messageID)
		}
	}
	return nil, errors.New("chat not found")
}

// Apply reaction to message
func (s *ReactionService) ApplyReaction(ctx context.Context, messageID, userID uuid.UUID, reaction *models.Reaction) (*models.Reaction, error) {
	message, err := s.messageRepo.GetMessageByID(ctx, messageID)
	if err != nil {
		return nil, err
	}

	// Check if user has already reacted to this message
	reactions, err := s.messageRepo.GetReactionsByMessageID(ctx, message.ID)
	if err != nil {
		return nil, err
	}
	for _, r := range reactions {
		if r.UserID == userID {
			return nil, errors.New("user has already reacted to this message")
		}
	}

	// Security mesaures
	// Check if user is in chat members
	chat, err := s.messageRepo.GetChatByMessageID(ctx, messageID)
	if err != nil {
		return nil, err
	}
	chatMembers, err := s.messageRepo.GetChatMembers(ctx, chat.ID)
	if err != nil {
		return nil, err
	}
	if len(chatMembers) == 0 {
		return nil, errors.New("chat not found")
	}
	for _, member := range chatMembers {
		if member == userID {
			reaction.MessageID = messageID
			reaction.UserID = userID
			reaction.CreatedAt = time.Now()

			return s.messageRepo.ApplyReaction(ctx, reaction)
		}
	}

	return nil, errors.New("chat not found")

}

// Remove reaction from message
func (s *ReactionService) RemoveReaction(ctx context.Context, messageID, userID uuid.UUID) error {
	reactions, err := s.messageRepo.GetReactionsByMessageID(ctx, messageID)
	if err != nil {
		return err
	}

	// Check if user has reacted to this message
	var reactionID uuid.UUID
	for _, r := range reactions {
		if r.UserID == userID {
			reactionID = r.ID
		}
	}
	if reactionID == uuid.Nil {
		return errors.New("user has not reacted to this message")
	}

	return s.messageRepo.RemoveReaction(ctx, reactionID)
}
