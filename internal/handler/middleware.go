package handler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/test-task/api"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userCtx"
)

func (h *Handler) userIdentity(c *fiber.Ctx) error {
	header := c.Get(authorizationHeader)
	h.logger.Infof("here")
	if header == "" {
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "not authorized",
		})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		})
	}

	_, err := h.service.VerifyToken(c.Context(), headerParts[1])
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&api.Error{
			Code:    http.StatusUnauthorized,
			Message: "invalid token",
		})
	}

	return c.Next()
}
