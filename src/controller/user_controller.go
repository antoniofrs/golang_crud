package controller

import (
	"golang_crud/src/dto"
	"golang_crud/src/plugin/logger"
	"golang_crud/src/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: *service}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var input dto.InsertUserDto

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	logger.Log.Info("controller: create user")

	user, err := uc.service.Create(c.Context(), input)
	if err != nil {
		logger.Log.Error("create user failed", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}


func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := uc.service.GetByID(c.Context(), id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}


func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetAll(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(users)
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uc.service.Delete(c.Context(), id); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
