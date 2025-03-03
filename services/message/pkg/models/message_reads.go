package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageRead struct {
	MessageID uuid.UUID `db:"type:uuid;not null;primaryKey;references:messages(id);onDelete:CASCADE"`
	UserID    uuid.UUID `db:"type:uuid;not null;primaryKey"`
	ReadAt    time.Time `db:"type:timestamptz;default:now()"`
}
