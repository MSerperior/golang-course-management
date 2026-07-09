package entity

import "github.com/google/uuid"

// Section represents a section inside a course (UUID PK)
type Section struct {
	Entity
	CourseId  *uuid.UUID `gorm:"column:course_id;type:varchar(36)"`
	Title     string     `gorm:"column:title"`
	SortOrder int        `gorm:"column:sort_order"`

	Course  *Course  `gorm:"foreignKey:course_id;references:id"`
	Lessons []Lesson `gorm:"foreignKey:section_id;references:id"`
}

func (s *Section) TableName() string {
	return "sections"
}
