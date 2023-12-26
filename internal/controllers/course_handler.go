package controllers

import "CourseApp/internal/services"

type CourseController struct {
	courseService services.CourseServices
}

func NewCourseController(courseService services.CourseServices) *CourseController {
	return &CourseController{courseService}
}
