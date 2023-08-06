package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/test-task/api"
	"github.com/kozhamseitova/test-task/internal/entity"
	"github.com/kozhamseitova/test-task/pkg/utils"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	var req entity.CreateUserRequest

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}

	id, err := h.service.CreateUser(c.Context(), &req)
	if err != nil {
		if errors.Is(err, utils.ErrUserAlreadyExists) {
			return c.Status(http.StatusConflict).JSON(&api.Error{
				Code:    http.StatusConflict,
				Message: err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(&api.Ok{
		Code:    http.StatusCreated,
		Message: "success",
		Data: fiber.Map{
			"id": id,
		},
	})
}

func (h *Handler) login(c *fiber.Ctx) error {
	var req api.LoginRequest

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}

	token, err := h.service.Login(c.Context(), req.Username, req.Password)

	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(&api.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"token": token,
		},
	})
}

func (h *Handler) getUsers(c *fiber.Ctx) error {
	var filter entity.UserFilter
	c.QueryParser(&filter)
	users, err := h.service.GetAllUsers(c.Context(), filter)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"users": users,
		},
	})
}

func (h *Handler) getUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid param",
		})
	}

	user, err := h.service.GetUsersById(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
		Data: fiber.Map{
			"user": user,
		},
	})
}

func (h *Handler) updateUser(c *fiber.Ctx) error {
	var req entity.CreateUserRequest

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid body param",
		})
	}

	err = h.service.UpdateUser(c.Context(), &req)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
	})

}

func (h *Handler) deleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid param",
		})
	}

	err = h.service.DeleteUser(c.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(&api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&api.Ok{
		Code:    http.StatusOK,
		Message: "success",
	})
}
