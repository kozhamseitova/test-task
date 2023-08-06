package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kozhamseitova/test-task/pkg/logger"
)

const usersTable = "users"

type Manager struct {
	pool *pgxpool.Pool
	logger logger.Logger
}

func NewRepository(pool *pgxpool.Pool, logger logger.Logger) *Manager {
	return &Manager{
		pool: pool,
		logger: logger,
	}
}
