package controllers

import (
	"net/http"

	"CourseApp/internal/domain"
	"CourseApp/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := &domain.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	user, err := c.userService.Register(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return response
	response := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User successfully Registered",
		"data":    response,
	})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := c.userService.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return response
	response := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User successfully login",
		"data":    response,
	})
}
