package model

import "time"

// Lesson

type LessonResponse struct {
	ID              string    `json:"id"`
	SectionId       string    `json:"section_id,omitempty"`
	Title           string    `json:"title"`
	Type            string    `json:"type"`
	ContentURL      string    `json:"content_url"`
	DurationSeconds int       `json:"duration_seconds"`
	SortOrder       int       `json:"sort_order"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateLessonRequest struct {
	SectionId       string `json:"section_id" validate:"required,max=100,uuid"`
	Title           string `json:"title" validate:"required,max=255"`
	Type            string `json:"type,omitempty" validate:"max=50"`
	ContentURL      string `json:"content_url,omitempty" validate:"max=1000"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
	SortOrder       int    `json:"sort_order,omitempty"`
}

type UpdateLessonRequest struct {
	ID              string `json:"-" validate:"required,max=100,uuid"`
	SectionId       string `json:"section_id,omitempty" validate:"max=100,uuid"`
	Title           string `json:"title,omitempty" validate:"max=255"`
	Type            string `json:"type,omitempty" validate:"max=50"`
	ContentURL      string `json:"content_url,omitempty" validate:"max=1000"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
	SortOrder       int    `json:"sort_order,omitempty"`
}

type GetLessonRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteLessonRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
