package handler

import (
	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/service"
	"github.com/kozhamseitova/test-task/pkg/logger"
)

type Handler struct {
	service service.Service
	cfg     *config.Config
	logger logger.Logger
}

func NewHandler(service service.Service, cfg *config.Config, logger logger.Logger) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
		logger: logger,
	}
}
