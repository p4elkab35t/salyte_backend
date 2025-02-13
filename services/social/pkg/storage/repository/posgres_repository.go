package repository

import (
	// "context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepositorySQL struct {
	db *pgxpool.Pool
}
