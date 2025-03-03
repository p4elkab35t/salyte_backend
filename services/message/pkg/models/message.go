package models

import (
	"time"

	"github.com/google/uuid"
)

// Message represents a chat message in the database.
type Message struct {
	ID        uuid.UUID `db:"id"`
	ChatID    uuid.UUID `db:"chat_id"`
	SenderID  uuid.UUID `db:"sender_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsDeleted bool      `db:"is_deleted"`
}
