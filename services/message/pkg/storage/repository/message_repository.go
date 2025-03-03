package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
)

type MessageRepository interface {
	SendMessage(ctx context.Context, message *models.Message) (*models.Message, error)
	EditMessage(ctx context.Context, messageID uuid.UUID, newContent string) error
	DeleteMessage(ctx context.Context, messageID uuid.UUID) error
	GetMessageByID(ctx context.Context, messageID uuid.UUID) (*models.Message, error)
	GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]*models.Message, error)
	GetAllChatsForUser(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
	CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	GetChatByID(ctx context.Context, chatID uuid.UUID) (*models.Chat, error)
	AddUserToChat(ctx context.Context, chatID, userID uuid.UUID) error
	RemoveUserFromChat(ctx context.Context, chatID, userID uuid.UUID) error
	GetChatMembers(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error)
	GetChatByMembers(ctx context.Context, members []uuid.UUID) (*models.Chat, error)
	DeleteAllMessagesInChat(ctx context.Context, chatID uuid.UUID) error
	GetChatByMessageID(ctx context.Context, messageID uuid.UUID) (*models.Chat, error)
	GetReactionsByMessageID(ctx context.Context, messageID uuid.UUID) ([]*models.Reaction, error)
	ApplyReaction(ctx context.Context, reaction *models.Reaction) (*models.Reaction, error)
	RemoveReaction(ctx context.Context, reactionID uuid.UUID) error
	ReadMessage(ctx context.Context, messageID uuid.UUID, userID uuid.UUID) error
	GetUnreadMessages(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
}
