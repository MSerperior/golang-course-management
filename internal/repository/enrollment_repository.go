package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	Repository[entity.Enrollment]
	Log *logrus.Logger
}

func NewEnrollmentRepository(log *logrus.Logger) *EnrollmentRepository {
	return &EnrollmentRepository{Log: log}
}

func (r *EnrollmentRepository) FindByStudentId(db *gorm.DB, studentId string) ([]entity.Enrollment, error) {
	var enrollments []entity.Enrollment
	if err := db.Where("student_id = ?", studentId).Find(&enrollments).Error; err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (r *EnrollmentRepository) CountByCourseId(db *gorm.DB, courseId string) (int64, error) {
	var total int64
	if err := db.Model(&entity.Enrollment{}).Where("course_id = ?", courseId).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
