package main

import (
	"CourseApp/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	app.StartApp()
}
