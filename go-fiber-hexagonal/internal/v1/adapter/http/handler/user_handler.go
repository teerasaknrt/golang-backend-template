package handler

import (
	"go-fiber-hexagonal/internal/v1/port"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase port.UserUseCase
}

func NewUserHandler(uc port.UserUseCase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.useCase.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
