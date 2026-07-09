package entity

import "github.com/google/uuid"

// Lesson represents a lesson inside a section (UUID PK)
type Lesson struct {
	Entity
	SectionId       *uuid.UUID `gorm:"column:section_id;type:varchar(36)"`
	Title           string     `gorm:"column:title"`
	Type            string     `gorm:"column:type"`
	ContentURL      string     `gorm:"column:content_url;type:text"`
	DurationSeconds int        `gorm:"column:duration_seconds"`
	SortOrder       int        `gorm:"column:sort_order"`

	Section *Section `gorm:"foreignKey:section_id;references:id"`
}

func (l *Lesson) TableName() string {
	return "lessons"
}
