package models

import (
	"time"

	"github.com/google/uuid"
)

type CommunityMembers struct {
	MemberID    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CommunityID uuid.UUID `gorm:"type:uuid"`
	ProfileID   uuid.UUID `gorm:"type:uuid"`
	Role        string    `gorm:"type:varchar(50)"`
	JoinedAt    time.Time
}

type Community struct {
	CommunityID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name              string    `gorm:"type:varchar(255)"`
	Description       string    `gorm:"type:text"`
	ProfilePictureURL string    `gorm:"type:varchar(255)"`
	Visibility        string    `gorm:"type:varchar(50)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
