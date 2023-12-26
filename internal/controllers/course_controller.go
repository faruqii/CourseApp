package controllers

import (
	"CourseApp/internal/domain"
	"CourseApp/internal/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (c *CourseController) CreateCourse(ctx *fiber.Ctx) error {
	var addRequest dto.CourseRequest

	if err := ctx.BodyParser(&addRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	course := &domain.Course{
		Name:        addRequest.Name,
		Description: addRequest.Description,
		TutorName:   addRequest.TutorName,
		Price:       addRequest.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	course, err := c.courseService.CreateCourse(course)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.CourseResponse{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		TutorName:   course.TutorName,
		Price:       course.Price,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Course successfully created",
		"data":    response,
	})
}

func (c *CourseController) GetAllCourses(ctx *fiber.Ctx) error {
	courses, err := c.courseService.GetAllCourses()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var response []dto.CourseResponse

	for _, course := range courses {
		response = append(response, dto.CourseResponse{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			TutorName:   course.TutorName,
			Price:       course.Price,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *CourseController) GetCourseById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	course, err := c.courseService.GetCourseById(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.CourseResponse{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		TutorName:   course.TutorName,
		Price:       course.Price,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (c *CourseController) UpdateCourse(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var updateRequest dto.CourseUpdateRequest

	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	course := &domain.Course{
		Name:        updateRequest.Name,
		Description: updateRequest.Description,
		TutorName:   updateRequest.TutorName,
		Price:       updateRequest.Price,
		UpdatedAt:   time.Now(),
	}

	course, err := c.courseService.UpdateCourse(id, course)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.CourseResponse{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		TutorName:   course.TutorName,
		Price:       course.Price,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Course successfully updated",
		"data":    response,
	})
}

func (c *CourseController) DeleteCourse(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.courseService.DeleteCourse(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *CourseController) CountPrice(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	var req dto.CoursePrice

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	// get course price
	price, err := c.courseService.GetCoursePrice(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// count price
	response := dto.CoursePriceResponse{
		Price: price * req.Price,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})

}
