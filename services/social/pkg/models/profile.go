package models

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ProfileID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID            uuid.UUID `gorm:"type:uuid"`
	Username          string    `gorm:"type:varchar(255)"`
	Bio               *string   `gorm:"type:text"`
	ProfilePictureURL *string   `gorm:"type:varchar(255)"`
	Visibility        *string   `gorm:"type:varchar(50)"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type Follower struct {
	FollowerID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FollowerProfileID uuid.UUID `gorm:"type:uuid"`
	FollowedProfileID uuid.UUID `gorm:"type:uuid"`
	CreatedAt         time.Time
}

type Interchange struct {
	InterchangeID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID       uuid.UUID `gorm:"type:uuid"`
	FriendProfileID uuid.UUID `gorm:"type:uuid"`
	Status          string    `gorm:"type:varchar(50)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Setting struct {
	SettingID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProfileID       uuid.UUID `gorm:"type:uuid"`
	DarkModeEnabled bool
	Language        string `gorm:"type:varchar(50)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
