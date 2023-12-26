package app

import (
	"os"

	"CourseApp/internal/config"
	"CourseApp/internal/repositories"
	"CourseApp/internal/routes"
	"CourseApp/internal/services"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	// init database
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	// init  user repository
	userRepositories := repositories.NewUserRepository(db)

	// init user  service
	userServices := services.NewUserServices(userRepositories)

	// init course repo
	courseRepositories := repositories.NewCourseRepository(db)

	// init course service
	courseService := services.NewCourseServices(courseRepositories)

	// init route
	apiEndpoint := app.Group("/api")
	routes.SetupUserRoute(apiEndpoint, userServices)
	routes.SetupCourseRoute(apiEndpoint, courseService)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
