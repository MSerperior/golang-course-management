package entity

import "github.com/google/uuid"

// Review represents a course review by a student (UUID PK)
type Review struct {
	Entity
	CourseId  *uuid.UUID `gorm:"column:course_id;type:varchar(36)"`
	StudentId *uuid.UUID `gorm:"column:student_id;type:varchar(36)"`
	Rating    int        `gorm:"column:rating"`
	Comment   string     `gorm:"column:comment;type:text"`

	Course  *Course `gorm:"foreignKey:course_id;references:id"`
	Student *User   `gorm:"foreignKey:student_id;references:id"`
}

func (r *Review) TableName() string {
	return "reviews"
}
