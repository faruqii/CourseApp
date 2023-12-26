package dto

import "time"

type CourseRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TutorName   string    `json:"tutor_name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CourseResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TutorName   string    `json:"tutor_name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CourseUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TutorName   string `json:"tutor_name"`
	Price       int    `json:"price"`
	UpdateAt    string `json:"updated_at"`
}

type CoursePrice struct {
	Price int `json:"price"`
}

type CoursePriceResponse struct {
	Price int `json:"price"`
}