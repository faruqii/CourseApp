package routes

import (
	"CourseApp/internal/controllers"
	"CourseApp/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupCourseRoute(router fiber.Router, courseService services.CourseServices) {
	courseController := controllers.NewCourseController(courseService)

	course := router.Group("/course")

	course.Post("/create", courseController.CreateCourse)
	course.Get("/get", courseController.GetAllCourses)
	course.Get("/get/:id", courseController.GetCourseById)
	course.Put("/update/:id", courseController.UpdateCourse)
	course.Delete("/delete/:id", courseController.DeleteCourse)
	course.Post("/price/:id", courseController.CountPrice)

}
