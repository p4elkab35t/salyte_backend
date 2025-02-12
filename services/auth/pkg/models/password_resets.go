package models

import (
	"time"
)

// PasswordReset struct represents password_resets table in database

type PasswordReset struct {
	Reset_id    string    `db:"reset_id"`
	User_id     string    `db:"user_id"`
	Reset_token string    `db:"reset_token"`
	Expires_at  time.Time `db:"expires_at"`
	CreatedAt   time.Time `db:"created_at"`
}
