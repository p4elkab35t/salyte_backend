package models

import (
	"time"

	"github.com/google/uuid"
)

// Chat represents a chat session or group conversation.
type Chat struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
