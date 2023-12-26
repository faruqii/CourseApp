package repositories

import (
	"CourseApp/internal/domain"

	"gorm.io/gorm"
)

type CourseRepository interface {
	Insert(course *domain.Course) (*domain.Course, error)
	FindAll() ([]domain.Course, error)
	FindById(id string) (*domain.Course, error)
	Update(id string, course *domain.Course) (*domain.Course, error)
	Delete(id string) error
	GetCoursePrice(id string) (int, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) Insert(course *domain.Course) (*domain.Course, error) {
	err := r.db.Create(&course).Error

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *courseRepository) FindAll() ([]domain.Course, error) {
	var courses []domain.Course

	err := r.db.Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepository) FindById(id string) (*domain.Course, error) {
	var course domain.Course

	err := r.db.Where("id = ?", id).First(&course).Error

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *courseRepository) Update(id string, course *domain.Course) (*domain.Course, error) {
	err := r.db.Model(&course).Where("id = ?", id).Updates(course).Error

	if err != nil {
		return nil, err
	}

	return course, nil
}


func (r *courseRepository) Delete(id string) error {
	var course domain.Course

	err := r.db.Where("id = ?", id).Delete(&course).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *courseRepository) GetCoursePrice(id string) (int, error) {
	var course domain.Course

	err := r.db.Where("id = ?", id).First(&course).Error

	if err != nil {
		return 0, err
	}

	return course.Price, nil
}


