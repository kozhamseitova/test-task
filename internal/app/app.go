package app

import (
	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/handler"
	"github.com/kozhamseitova/test-task/internal/repository"
	"github.com/kozhamseitova/test-task/internal/service"
	"github.com/kozhamseitova/test-task/pkg/client/postgres"
	"github.com/kozhamseitova/test-task/pkg/jwttoken"
	"github.com/kozhamseitova/test-task/pkg/logger"
)

func Run(cfg *config.Config) error {
	dbConn, err := postgres.New(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)

	if err != nil {
		return err
	}

	token := jwttoken.New(cfg.Token.SecretKey)

	logger, err := logger.NewLogger(cfg.App.Production)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(dbConn.Pool, logger)
	srvc := service.NewService(repo, token, cfg, logger)
	hndlr := handler.NewHandler(srvc, cfg, logger)

	hndlr.InitRouter()
	return nil
}
