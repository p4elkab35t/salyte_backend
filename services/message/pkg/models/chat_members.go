package models

import (
	"time"

	"github.com/google/uuid"
)

// ChatMember represents a user that is part of a chat.
type ChatMember struct {
	ID       uuid.UUID `db:"id"`
	ChatID   uuid.UUID `db:"chat_id"`
	UserID   uuid.UUID `db:"user_id"`
	JoinedAt time.Time `db:"joined_at"`
}
