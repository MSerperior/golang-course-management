package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SectionRepository struct {
	Repository[entity.Section]
	Log *logrus.Logger
}

func NewSectionRepository(log *logrus.Logger) *SectionRepository {
	return &SectionRepository{Log: log}
}

func (r *SectionRepository) FindAllByCourseId(db *gorm.DB, courseId string) ([]entity.Section, error) {
	var sections []entity.Section
	if err := db.Where("course_id = ?", courseId).Order("sort_order asc").Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
}
