package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type EnrollmentRepository struct {
	Repository[entity.Enrollment]
	Log *logrus.Logger
}

func NewEnrollmentRepository(log *logrus.Logger) *EnrollmentRepository {
	return &EnrollmentRepository{Log: log}
}
