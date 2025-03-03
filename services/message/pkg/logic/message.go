package logic

import (
	"context"
	"errors"

	// "os/user"
	// "fmt"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/storage/repository"
)

type MessageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo}
}

// Create new message
func (s *MessageService) SendMessage(ctx context.Context, message *models.Message, userID uuid.UUID) (*models.Message, error) {
	if message.Content == "" {
		return nil, errors.New("content is required")
	}

	if message.ChatID == uuid.Nil {
		return nil, errors.New("chat_id is required")
	}

	if message.SenderID == uuid.Nil {
		return nil, errors.New("sender_id is required")
	}

	//security mesaures
	// get chat by chat ID, get chat members by chat ID, check if user is in chat members
	chat, err := s.messageRepo.GetChatByID(ctx, message.ChatID)
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
			return s.messageRepo.SendMessage(ctx, message)
		}
	}
	return nil, errors.New("chat not found")
}

func (s *MessageService) GetMessageByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*models.Message, error) {
	message, err := s.messageRepo.GetMessageByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// security mesaures
	// get chat by message ID, get chat members by chat ID, check if user is in chat members
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
			return message, nil
		}
	}
	return nil, errors.New("message not found")

}

func (s *MessageService) EditMessage(ctx context.Context, messageID uuid.UUID, newContent string, userID uuid.UUID) error {
	if newContent == "" {
		return errors.New("new content is empty")
	}

	if messageID == uuid.Nil {
		return errors.New("message_id is required")
	}

	message, err := s.GetMessageByID(ctx, messageID, userID)
	if err != nil {
		return errors.New("message not found")
	}

	// security mesaures
	// check if user is the sender of the message
	if message.SenderID != userID {
		return errors.New("you can't delete this message")
	}

	return s.messageRepo.EditMessage(ctx, messageID, newContent)
}

func (s *MessageService) DeleteMessage(ctx context.Context, messageID uuid.UUID, userID uuid.UUID) error {
	if messageID == uuid.Nil {
		return errors.New("message_id is required")
	}

	message, err := s.GetMessageByID(ctx, messageID, userID)
	if err != nil {
		return errors.New("message not found")
	}

	// security mesaures
	// check if user is the sender of the message
	if message.SenderID != userID {
		return errors.New("you can't delete this message")
	}

	return s.messageRepo.DeleteMessage(ctx, messageID)
}

func (s *MessageService) GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, limit, offset int, userID uuid.UUID) ([]*models.Message, error) {
	// security mesaures
	// get chat by chat ID, get chat members by chat ID, check if user is in chat members
	chat, err := s.messageRepo.GetChatByID(ctx, chatID)
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
			return s.messageRepo.GetMessagesByChatID(ctx, chatID, limit, offset)
		}
	}
	return nil, errors.New("chat not found")
}

func (s *MessageService) GetNewMessages(ctx context.Context, chatID uuid.UUID, since int64, userID uuid.UUID) ([]*models.Message, error) {
	// security mesaures
	// get chat by chat ID, get chat members by chat ID, check if user is in chat members
	chat, err := s.messageRepo.GetChatByID(ctx, chatID)
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
			return s.messageRepo.GetNewMessages(ctx, chatID, since)
		}
	}
	return nil, errors.New("chat not found")
}

func (s *MessageService) DeleteAllMessagesInChat(ctx context.Context, chatID uuid.UUID, userID uuid.UUID) error {
	// security mesaures
	// get chat by chat ID, get chat members by chat ID, check if user is in chat members
	chat, err := s.messageRepo.GetChatByID(ctx, chatID)
	if err != nil {
		return err
	}
	chatMembers, err := s.messageRepo.GetChatMembers(ctx, chat.ID)
	if err != nil {
		return err
	}
	if len(chatMembers) == 0 {
		return errors.New("chat not found")
	}
	for _, member := range chatMembers {
		if member == userID {
			return s.messageRepo.DeleteAllMessagesInChat(ctx, chatID)
		}
	}
	return errors.New("chat not found")
}
