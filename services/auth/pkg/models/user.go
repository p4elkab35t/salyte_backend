package models

import (
	"time"
)

// users table struct

type User struct {
	User_id       string    `db:"user_id"`
	Email         string    `db:"email"`
	Password_hash string    `db:"password_hash"`
	Is_verified   bool      `db:"is_verified"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
