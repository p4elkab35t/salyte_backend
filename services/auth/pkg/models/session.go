package models

import (
	"time"
)

// Session represents a user session

type Session struct {
	Session_id    string    `db:"session_id"`
	User_id       string    `db:"user_id"`
	Session_token string    `db:"session_token"`
	Expires_at    time.Time `db:"expires_at"`
	CreatedAt     time.Time `db:"created_at"`
}
