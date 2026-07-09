package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseRepository struct {
	Repository[entity.Course]
	Log *logrus.Logger
}

func NewCourseRepository(log *logrus.Logger) *CourseRepository {
	return &CourseRepository{Log: log}
}

func (r *CourseRepository) FindBySlug(db *gorm.DB, course *entity.Course, slug string) error {
	return db.Where("slug = ?", slug).Take(course).Error
}

func (r *CourseRepository) FindAllByInstructorId(db *gorm.DB, instructorId string, page, size int) ([]entity.Course, int64, error) {
	var courses []entity.Course
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	if err := db.Where("instructor_id = ?", instructorId).
		Offset((page - 1) * size).Limit(size).
		Find(&courses).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.Model(&entity.Course{}).
		Where("instructor_id = ?", instructorId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}
