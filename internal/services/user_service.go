package services

import (
	"net/http"

	"CourseApp/internal/config"
	"CourseApp/internal/domain"
	"CourseApp/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	Register(user *domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
}

type userServices struct {
	userRepository repositories.UserRepository
}

func NewUserServices(userRepository repositories.UserRepository) *userServices {
	return &userServices{userRepository}
}

func (s *userServices) Register(user *domain.User) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	// Check if user is already registered to database
	_, err = repo.FindByEmail(user.Email)

	if err == nil {
		return nil, &ErrorMessages{
			Message:    "Email already registered",
			StatusCode: http.StatusBadRequest,
		}
	}

	// hash user password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to hash password",
			StatusCode: http.StatusInternalServerError,
		}
	}

	user.Password = string(hashPassword)

	// Insert user to database
	user, err = repo.Insert(user)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Failed to insert user to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return user, nil

}

func (s *userServices) Login(email, password string) (*domain.User, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message:   "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewUserRepository(conn)

	user, err := repo.FindByEmail(email)

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "User not found!",
			StatusCode: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessages{
			Message:    "Invalid password",
			StatusCode: http.StatusUnauthorized,
		}
	}

	return user, nil
}
