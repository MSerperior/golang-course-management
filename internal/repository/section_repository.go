package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type SectionRepository struct {
	Repository[entity.Section]
	Log *logrus.Logger
}

func NewSectionRepository(log *logrus.Logger) *SectionRepository {
	return &SectionRepository{Log: log}
}
