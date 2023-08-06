package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InitRouter(router fiber.Router) {
	router.Use(h.generateTraceId)

	auth := router.Group("auth")
	auth.Post("/register", h.createUser)
	auth.Post("/login", h.login)

	api := router.Group("api/v1")

	user := api.Group("/users", h.userIdentity)

	user.Get("/", h.getUsers)
	user.Get("/:id", h.getUserById)
	user.Put("/", h.updateUser)
	user.Delete("/:id", h.deleteUser)
}
