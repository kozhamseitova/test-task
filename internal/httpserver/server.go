package httpserver

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/handler"
)

type Server struct {
	app             *fiber.App
	cfg *config.Config
	shutdownTimeout time.Duration
	notify          chan error
}

func New(handler *handler.Handler, cfg *config.Config) *Server {
	app := fiber.New()

	handler.InitRouter(app)

	s := &Server{
		app: app,
		cfg: cfg,
		notify: make(chan error, 1),
	}
	
	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.app.Listen(s.cfg.HTTP.Port)
		close(s.notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)

	defer cancel()

	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
