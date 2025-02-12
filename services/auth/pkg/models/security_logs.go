package models

import (
	"time"
)

// SecurityLog represents a security log.

type SecurityLog struct {
	Log_id     string    `db:"log_id"`
	User_id    string    `db:"user_id"`
	Action     string    `db:"action"`
	Ip_address string    `db:"ip_address"`
	Timestamp  time.Time `db:"timestamp"`
}
