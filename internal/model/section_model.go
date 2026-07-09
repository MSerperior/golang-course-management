package model

import "time"

// Section

type SectionResponse struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id,omitempty"`
	Title     string    `json:"title"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSectionRequest struct {
	CourseId  string `json:"course_id" validate:"required,max=100,uuid"`
	Title     string `json:"title" validate:"required,max=255"`
	SortOrder int    `json:"sort_order,omitempty"`
}

type UpdateSectionRequest struct {
	ID        string `json:"-" validate:"required,max=100,uuid"`
	CourseId  string `json:"course_id,omitempty" validate:"max=100,uuid"`
	Title     string `json:"title,omitempty" validate:"max=255"`
	SortOrder int    `json:"sort_order,omitempty"`
}

type GetSectionRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteSectionRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
