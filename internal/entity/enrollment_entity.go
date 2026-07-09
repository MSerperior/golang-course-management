package entity

import (
	"time"

	"github.com/google/uuid"
)

// Enrollment represents a student's enrollment in a course (UUID PK)
type Enrollment struct {
	Entity
	StudentId   *uuid.UUID `gorm:"column:student_id;type:varchar(36)"`
	CourseId    *uuid.UUID `gorm:"column:course_id;type:varchar(36)"`
	Status      string     `gorm:"column:status"`
	EnrolledAt  time.Time  `gorm:"column:enrolled_at"`
	CompletedAt time.Time  `gorm:"column:completed_at"`

	Student *User   `gorm:"foreignKey:student_id;references:id"`
	Course  *Course `gorm:"foreignKey:course_id;references:id"`
}

func (e *Enrollment) TableName() string {
	return "enrollments"
}
