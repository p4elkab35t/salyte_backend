package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/models"
)

type PostgresRepositorySQL struct {
	db *pgxpool.Pool
}

func NewPostgresMessageRepositorySQL(db *pgxpool.Pool) MessageRepository {
	return &PostgresRepositorySQL{db: db}
}

//SendMessage(ctx context.Context, message *models.Message) (*models.Message, error)

func (r *PostgresRepositorySQL) SendMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	if message.ChatID == uuid.Nil || message.SenderID == uuid.Nil || message.Content == "" {
		return nil, errors.New("chat_id, sender_id and content are required")
	}
	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now
	row := r.db.QueryRow(ctx,
		"INSERT INTO messages (chat_id, sender_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING message_id",
		message.ChatID, message.SenderID, message.Content, message.CreatedAt, message.UpdatedAt)
	err := row.Scan(&message.ID)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// EditMessage(ctx context.Context, messageID uuid.UUID, newContent string) error

func (r *PostgresRepositorySQL) EditMessage(ctx context.Context, messageID uuid.UUID, newContent string) error {
	if messageID == uuid.Nil || newContent == "" {
		return errors.New("message_id and new_content are required")
	}
	_, err := r.db.Exec(ctx,
		"UPDATE messages SET content = $1, updated_at = $2 WHERE message_id = $3",
		newContent, time.Now(), messageID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMessage(ctx context.Context, messageID uuid.UUID) error

func (r *PostgresRepositorySQL) DeleteMessage(ctx context.Context, messageID uuid.UUID) error {
	if messageID == uuid.Nil {
		return errors.New("message_id is required")
	}
	_, err := r.db.Exec(ctx,
		"UPDATE messages SET is_deleted = true WHERE message_id = $1",
		messageID)
	if err != nil {
		return err
	}
	return nil
}

// GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]*models.Message, error)

func (r *PostgresRepositorySQL) GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]*models.Message, error) {
	if chatID == uuid.Nil {
		return nil, errors.New("chat_id is required")
	}
	if limit < 0 || offset < 0 {
		return nil, errors.New("limit and offset must be greater than 0")
	}
	if limit > 100 {
		return nil, errors.New("limit must be less than or equal to 100")
	}
	rows, err := r.db.Query(ctx,
		"SELECT message_id, chat_id, sender_id, content, created_at, updated_at, is_deleted FROM messages WHERE chat_id = $1 AND is_deleted = FALSE ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		chatID, limit, offset*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := make([]*models.Message, 0)
	for rows.Next() {
		message := &models.Message{}
		err = rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.IsDeleted)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// GetAllChatsForUser(ctx context.Context, userID uuid.UUID) ([]int64, error)

func (r *PostgresRepositorySQL) GetAllChatsForUser(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}
	rows, err := r.db.Query(ctx,
		"SELECT chat_id FROM chat_members WHERE user_id = $1",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	chatIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		var chatID uuid.UUID
		err = rows.Scan(&chatID)
		if err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}
	return chatIDs, nil
}

// CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)

func (r *PostgresRepositorySQL) CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {

	if chat.Name == "" {
		return nil, errors.New("chat name is required")
	}

	now := time.Now()
	chat.CreatedAt = now
	chat.UpdatedAt = now
	row := r.db.QueryRow(ctx,
		"INSERT INTO chats (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING chat_id",
		chat.Name, chat.CreatedAt, chat.UpdatedAt)
	err := row.Scan(&chat.ID)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// GetChatByID(ctx context.Context, chatID uuid.UUID) (*models.Chat, error)

func (r *PostgresRepositorySQL) GetChatByID(ctx context.Context, chatID uuid.UUID) (*models.Chat, error) {
	if chatID == uuid.Nil {
		return nil, errors.New("chat_id is required")
	}

	row := r.db.QueryRow(ctx,
		"SELECT chat_id, title, created_at, updated_at FROM chats WHERE chat_id = $1",
		chatID)
	chat := &models.Chat{}
	err := row.Scan(&chat.ID, &chat.Name, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// AddUserToChat(ctx context.Context, chatID, userID uuid.UUID) error

func (r *PostgresRepositorySQL) AddUserToChat(ctx context.Context, chatID, userID uuid.UUID) error {
	if chatID == uuid.Nil || userID == uuid.Nil {
		return errors.New("chat_id and user_id are required")
	}
	_, err := r.db.Exec(ctx,
		"INSERT INTO chat_members (chat_id, user_id) VALUES ($1, $2)",
		chatID, userID)
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserFromChat(ctx context.Context, chatID, userID uuid.UUID) error

func (r *PostgresRepositorySQL) RemoveUserFromChat(ctx context.Context, chatID, userID uuid.UUID) error {
	if chatID == uuid.Nil || userID == uuid.Nil {
		return errors.New("chat_id and user_id are required")
	}
	_, err := r.db.Exec(ctx,
		"DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2",
		chatID, userID)
	if err != nil {
		return err
	}
	return nil
}

// GetChatMembers(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error)

func (r *PostgresRepositorySQL) GetChatMembers(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error) {
	if chatID == uuid.Nil {
		return nil, errors.New("chat_id is required")
	}
	rows, err := r.db.Query(ctx,
		"SELECT user_id FROM chat_members WHERE chat_id = $1",
		chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	userIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		var userID uuid.UUID
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}

// GetChatByMembers(ctx context.Context, members []uuid.UUID) (*models.Chat, error)

func (r *PostgresRepositorySQL) GetChatByMembers(ctx context.Context, members []uuid.UUID) (*models.Chat, error) {
	if len(members) == 0 {
		return nil, errors.New("members are required")
	}
	query := "SELECT chat_id, title, created_at, updated_at FROM chats WHERE chat_id IN (SELECT chat_id FROM chat_members WHERE user_id IN ("
	args := make([]interface{}, 0)
	for i, member := range members {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("$%d", i+1)
		args = append(args, member)
	}
	query += fmt.Sprintf(") GROUP BY chat_id HAVING COUNT(DISTINCT user_id) = %d)", len(members))
	row := r.db.QueryRow(ctx, query, args...)
	chat := &models.Chat{}
	err := row.Scan(&chat.ID, &chat.Name, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

// DeleteAllMessagesInChat(ctx context.Context, chatID uuid.UUID) error

func (r *PostgresRepositorySQL) DeleteAllMessagesInChat(ctx context.Context, chatID uuid.UUID) error {
	if chatID == uuid.Nil {
		return errors.New("chat_id is required")
	}
	_, err := r.db.Exec(ctx,
		"UPDATE messages SET is_deleted = true WHERE chat_id = $1",
		chatID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepositorySQL) GetMessageByID(ctx context.Context, messageID uuid.UUID) (*models.Message, error) {
	if messageID == uuid.Nil {
		return nil, errors.New("message_id is required")
	}
	row := r.db.QueryRow(ctx,
		"SELECT message_id, chat_id, sender_id, content, created_at, updated_at, is_deleted FROM messages WHERE message_id = $1 AND is_deleted = FALSE",
		messageID)
	message := &models.Message{}
	err := row.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.IsDeleted)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *PostgresRepositorySQL) GetChatByMessageID(ctx context.Context, messageID uuid.UUID) (*models.Chat, error) {
	if messageID == uuid.Nil {
		return nil, errors.New("message_id is required")
	}
	row := r.db.QueryRow(ctx,
		"SELECT chat_id FROM messages WHERE message_id = $1 AND is_deleted = FALSE",
		messageID)
	var chatID uuid.UUID
	err := row.Scan(&chatID)
	if err != nil {
		return nil, err
	}
	return r.GetChatByID(ctx, chatID)
}

func (r *PostgresRepositorySQL) GetReactionsByMessageID(ctx context.Context, messageID uuid.UUID) ([]*models.Reaction, error) {
	if messageID == uuid.Nil {
		return nil, errors.New("message_id is required")
	}
	rows, err := r.db.Query(ctx,
		"SELECT reaction_id, message_id, user_id, reaction, created_at FROM reactions WHERE message_id = $1",
		messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	reactions := make([]*models.Reaction, 0)
	for rows.Next() {
		reaction := &models.Reaction{}
		err = rows.Scan(&reaction.ID, &reaction.MessageID, &reaction.UserID, &reaction.Emoji, &reaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}
	return reactions, nil
}

func (r *PostgresRepositorySQL) ApplyReaction(ctx context.Context, reaction *models.Reaction) (*models.Reaction, error) {
	if reaction.MessageID == uuid.Nil || reaction.UserID == uuid.Nil {
		return nil, errors.New("message_id and user_id are required")
	}
	_, err := r.db.Exec(ctx,
		"INSERT INTO reactions (message_id, user_id, reaction, created_at) VALUES ($1, $2, $3, $4)",
		reaction.MessageID, reaction.UserID, reaction.Emoji, time.Now())
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

func (r *PostgresRepositorySQL) RemoveReaction(ctx context.Context, reactionID uuid.UUID) error {
	if reactionID == uuid.Nil {
		return errors.New("reaction_id is required")
	}
	_, err := r.db.Exec(ctx,
		"DELETE FROM reactions WHERE reaction_id = $1",
		reactionID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepositorySQL) ReadMessage(ctx context.Context, messageID uuid.UUID, userID uuid.UUID) error {
	if messageID == uuid.Nil || userID == uuid.Nil {
		return errors.New("message_id and user_id are required")
	}
	_, err := r.db.Exec(ctx,
		"INSERT INTO read_messages (message_id, user_id) VALUES ($1, $2)",
		messageID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepositorySQL) GetUnreadMessages(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}
	rows, err := r.db.Query(ctx,
		"SELECT message_id FROM messages WHERE message_id NOT IN (SELECT message_id FROM read_messages WHERE user_id = $1 AND is_deleted = FALSE) AND is_deleted = FALSE",
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messageIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		var messageID uuid.UUID
		err = rows.Scan(&messageID)
		if err != nil {
			return nil, err
		}
		messageIDs = append(messageIDs, messageID)
	}
	return messageIDs, nil
}
