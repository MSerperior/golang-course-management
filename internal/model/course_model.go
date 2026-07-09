package model

import "time"

// Course

type CourseResponse struct {
	ID           string    `json:"id"`
	InstructorId string    `json:"instructor_id,omitempty"`
	CategoryId   string    `json:"category_id,omitempty"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateCourseRequest struct {
	InstructorId string  `json:"instructor_id,omitempty" validate:"max=100,uuid"`
	CategoryId   string  `json:"category_id,omitempty" validate:"max=100,uuid"`
	Title        string  `json:"title" validate:"required,max=255"`
	Slug         string  `json:"slug" validate:"required,max=255"`
	Description  string  `json:"description,omitempty" validate:"max=2000"`
	Price        float64 `json:"price,omitempty"`
	Status       string  `json:"status,omitempty" validate:"max=50"`
}

type UpdateCourseRequest struct {
	ID           string  `json:"-" validate:"required,max=100,uuid"`
	InstructorId string  `json:"instructor_id,omitempty" validate:"max=100,uuid"`
	CategoryId   string  `json:"category_id,omitempty" validate:"max=100,uuid"`
	Title        string  `json:"title,omitempty" validate:"max=255"`
	Slug         string  `json:"slug,omitempty" validate:"max=255"`
	Description  string  `json:"description,omitempty" validate:"max=2000"`
	Price        float64 `json:"price,omitempty"`
	Status       string  `json:"status,omitempty" validate:"max=50"`
}

type GetCourseRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteCourseRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type ListCourseRequest struct {
	Title        string `json:"title,omitempty" validate:"max=255"`
	CategoryId   string `json:"category_id,omitempty" validate:"max=100,uuid"`
	InstructorId string `json:"instructor_id,omitempty" validate:"max=100,uuid"`
	Status       string `json:"status,omitempty" validate:"max=50"`
	Page         int    `json:"page,omitempty" validate:"min=1"`
	Size         int    `json:"size,omitempty" validate:"min=1,max=100"`
}
