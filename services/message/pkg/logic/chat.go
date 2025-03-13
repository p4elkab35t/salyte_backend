package logic

import (
	"context"

	// "os/user"
	// "fmt"
	"errors"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/storage/repository"
)

type ChatService struct {
	messageRepo repository.MessageRepository
}

func NewChatService(messageRepo repository.MessageRepository) *ChatService {
	return &ChatService{messageRepo: messageRepo}
}

// Create new chat
func (s *ChatService) CreateChat(ctx context.Context, users []*models.ChatMember, chat *models.Chat) (*models.Chat, error) {
	chat, err := s.messageRepo.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		err := s.messageRepo.AddUserToChat(ctx, chat.ID, user.UserID)
		if err != nil {
			return nil, err
		}
	}
	return chat, nil
}

func (s *ChatService) GetChatByID(ctx context.Context, id uuid.UUID) (*models.Chat, error) {
	return s.messageRepo.GetChatByID(ctx, id)
}

func (s *ChatService) GetChatsByUserID(ctx context.Context, userID uuid.UUID) ([]*models.Chat, error) {
	chatIDs, err := s.messageRepo.GetAllChatsForUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	chats := make([]*models.Chat, 0)
	for _, chatID := range chatIDs {
		chat, err := s.messageRepo.GetChatByID(ctx, chatID)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil

}

func (s *ChatService) AddUserToChat(ctx context.Context, chatID, userID, addedUserID uuid.UUID) error {
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
			return s.messageRepo.AddUserToChat(ctx, chatID, addedUserID)
		}
	}
	return errors.New("chat not found")
}

func (s *ChatService) RemoveUserFromChat(ctx context.Context, chatID, userID, removedUserID uuid.UUID) error {
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
			return s.messageRepo.RemoveUserFromChat(ctx, chatID, removedUserID)
		}
	}
	return errors.New("chat not found")
}

func (s *ChatService) GetChatMembers(ctx context.Context, userID, chatID uuid.UUID) ([]uuid.UUID, error) {
	//security mesaures
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
			return chatMembers, nil
		}
	}
	return nil, errors.New("chat not found")
}

func (s *ChatService) GetChatByTwoUsers(ctx context.Context, userID, secondUserID uuid.UUID) (*models.Chat, error) {
	chatMembers := []uuid.UUID{userID, secondUserID}
	chat, err := s.messageRepo.GetChatByMembers(ctx, chatMembers)
	if chat == nil {
		users := []*models.ChatMember{
			{UserID: userID},
			{UserID: secondUserID},
		}
		chatName := "Chat for " + userID.String() + " and " + secondUserID.String()
		chat = &models.Chat{Name: chatName}
		chat, err = s.CreateChat(ctx, users, chat)
		if err != nil {
			return nil, err
		}

		return s.messageRepo.GetChatByID(ctx, chat.ID)
	}
	if err != nil {
		return nil, err
	}
	return s.messageRepo.GetChatByID(ctx, chat.ID)
}
