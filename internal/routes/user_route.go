package routes

import (
	"CourseApp/internal/controllers"
	"CourseApp/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router, userService services.UserServices) {
	userController := controllers.NewUserController(userService)

	router.Post("/register", userController.Register)
	router.Post("/login", userController.Login)
	
}