package plugin

import (
	"golang_crud/src/controller"
	"golang_crud/src/plugin/logger"
	"golang_crud/src/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RegisterRoutes(userService *service.UserService) {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	api := app.Group("/api")
	user := api.Group("/users")

	// User controller
	userController := controller.NewUserController(userService)
	user.Get("/v1", userController.GetUsers)
	user.Post("/v1", userController.CreateUser)
	user.Get("/:id/v1", userController.GetUserByID)
	user.Delete("/:id/v1", userController.DeleteUser)

	if err := app.Listen(":3000"); err != nil {
		logger.Log.Fatal("fiber failed to start", zap.Error(err))
	}
}
