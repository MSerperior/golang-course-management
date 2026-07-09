package entity

import "github.com/google/uuid"

// Course represents a course (UUID PK)
type Course struct {
	Entity
	InstructorId *uuid.UUID `gorm:"column:instructor_id;type:varchar(36)"`
	CategoryId   *uuid.UUID `gorm:"column:category_id;type:varchar(36)"`
	Title        string     `gorm:"column:title"`
	Slug         string     `gorm:"column:slug;unique"`
	Description  string     `gorm:"column:description;type:text"`
	Price        float64    `gorm:"column:price"`
	Status       string     `gorm:"column:status"`

	Instructor   *User         `gorm:"foreignKey:instructor_id;references:id"`
	Category     *Category     `gorm:"foreignKey:category_id;references:id"`
	Sections     []Section     `gorm:"foreignKey:course_id;references:id"`
	Enrollments  []Enrollment  `gorm:"foreignKey:course_id;references:id"`
	Transactions []Transaction `gorm:"foreignKey:course_id;references:id"`
	Reviews      []Review      `gorm:"foreignKey:course_id;references:id"`
}

func (c *Course) TableName() string {
	return "courses"
}
