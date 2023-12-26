package controllers

import "CourseApp/internal/services"

type UserController struct {
	userService services.UserServices
}

func NewUserController(userService services.UserServices) *UserController {
	return &UserController{userService}
}
