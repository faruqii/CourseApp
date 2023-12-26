package services

import (
	"CourseApp/internal/config"
	"CourseApp/internal/domain"
	"CourseApp/internal/repositories"
	"net/http"
)

type CourseServices interface {
	CreateCourse(course *domain.Course) (*domain.Course, error)
	GetAllCourses() ([]domain.Course, error)
	GetCourseById(id string) (*domain.Course, error)
	UpdateCourse(id string, course *domain.Course) (*domain.Course, error)
	DeleteCourse(id string) error
	GetCoursePrice(id string) (int, error)
}

type courseServices struct {
	courseRepository repositories.CourseRepository
}

func NewCourseServices(courseRepository repositories.CourseRepository) *courseServices {
	return &courseServices{courseRepository}
}

func (s *courseServices) CreateCourse(course *domain.Course) (*domain.Course, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Insert course to database
	course, err = repo.Insert(course)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to insert course to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return course, nil
}

func (s *courseServices) GetAllCourses() ([]domain.Course, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Get all courses from database
	courses, err := repo.FindAll()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to get courses from database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return courses, nil
}

func (s *courseServices) GetCourseById(id string) (*domain.Course, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Get course by id from database
	course, err := repo.FindById(id)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to get course from database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return course, nil
}

func (s *courseServices) UpdateCourse(id string, course *domain.Course) (*domain.Course, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Update course by id from database
	course, err = repo.Update(id, course)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to update course from database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return course, nil
}

func (s *courseServices) DeleteCourse(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Delete course by id from database
	err = repo.Delete(id)

	if err != nil {
		return &ErrorMessages{
			Message:    "Failed to delete course from database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *courseServices) GetCoursePrice(id string) (int, error) {
	conn, err := config.Connect()

	if err != nil {
		return 0, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewCourseRepository(conn)

	// Get course by id from database
	course, err := repo.FindById(id)

	if err != nil {
		return 0, &ErrorMessages{
			Message:    "Failed to get course from database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return course.Price, nil
}
