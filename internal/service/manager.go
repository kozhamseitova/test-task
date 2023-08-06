package service

import (
	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/repository"
	"github.com/kozhamseitova/test-task/pkg/jwttoken"
	"github.com/kozhamseitova/test-task/pkg/logger"
)

type Manager struct {
	repository repository.Repository
	jwttoken   *jwttoken.JWTToken
	config     *config.Config
	logger     logger.Logger
}

func NewService(repository repository.Repository, jwttoken *jwttoken.JWTToken, config *config.Config, logger logger.Logger) *Manager {
	return &Manager{
		repository: repository,
		jwttoken:   jwttoken,
		config:     config,
		logger:     logger,
	}
}
