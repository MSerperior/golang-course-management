package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LessonRepository struct {
	Repository[entity.Lesson]
	Log *logrus.Logger
}

func NewLessonRepository(log *logrus.Logger) *LessonRepository {
	return &LessonRepository{Log: log}
}

func (r *LessonRepository) FindBySectionId(db *gorm.DB, sectionId string) ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	if err := db.Where("section_id = ?", sectionId).Order("sort_order asc").Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

// UpdateSortOrder updates the sort_order of multiple lessons within the provided transaction.
// Caller should provide a transaction (tx := db.Begin()) to ensure atomicity.
func (r *LessonRepository) UpdateSortOrder(tx *gorm.DB, lessons []entity.Lesson) error {
	for _, l := range lessons {
		if l.ID == nil {
			continue
		}
		if err := tx.Model(&entity.Lesson{}).
			Where("id = ?", l.ID).
			Update("sort_order", l.SortOrder).Error; err != nil {
			return err
		}
	}
	return nil
}
