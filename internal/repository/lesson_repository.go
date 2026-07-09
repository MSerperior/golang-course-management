package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type LessonRepository struct {
	Repository[entity.Lesson]
	Log *logrus.Logger
}

func NewLessonRepository(log *logrus.Logger) *LessonRepository {
	return &LessonRepository{Log: log}
}
