package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/handler"
	"github.com/kozhamseitova/test-task/internal/httpserver"
	"github.com/kozhamseitova/test-task/internal/repository"
	"github.com/kozhamseitova/test-task/internal/service"
	"github.com/kozhamseitova/test-task/pkg/client/postgres"
	"github.com/kozhamseitova/test-task/pkg/jwttoken"
	"github.com/kozhamseitova/test-task/pkg/logger"
)

func Run(ctx context.Context, cfg *config.Config) error {
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

	server := httpserver.New(
		hndlr,
		cfg,
	)

	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		logger.Infof(ctx, "signal received: %s", s.String())
	case err = <-server.Notify():
		logger.Errorf(ctx, "server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		logger.Errorf(ctx, "server shutdown err: %s", err)
	}


	return nil
}
