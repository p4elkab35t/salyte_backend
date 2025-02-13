package models

import (
	"time"

	"github.com/google/uuid"
)

type Share struct {
	ShareID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID uuid.UUID `gorm:"type:uuid"`
	PostID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
}

type Comment struct {
	CommentID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID uuid.UUID `gorm:"type:uuid"`
	PostID    uuid.UUID `gorm:"type:uuid"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Like struct {
	LikeID    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID uuid.UUID `gorm:"type:uuid"`
	PostID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
}

type Post struct {
	PostID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID   uuid.UUID `gorm:"type:uuid"`
	CommunityID uuid.UUID `gorm:"type:uuid"`
	Content     string    `gorm:"type:text"`
	MediaURL    string    `gorm:"type:varchar(255)"`
	Visibility  string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
