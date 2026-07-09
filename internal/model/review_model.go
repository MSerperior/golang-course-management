package model

import "time"

// Review

type ReviewResponse struct {
	ID        string    `json:"id"`
	CourseId  string    `json:"course_id,omitempty"`
	StudentId string    `json:"student_id,omitempty"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateReviewRequest struct {
	CourseId  string `json:"course_id" validate:"required,max=100,uuid"`
	StudentId string `json:"student_id" validate:"required,max=100,uuid"`
	Rating    int    `json:"rating" validate:"required,min=1,max=5"`
	Comment   string `json:"comment,omitempty" validate:"max=2000"`
}

type UpdateReviewRequest struct {
	ID      string `json:"-" validate:"required,max=100,uuid"`
	Rating  int    `json:"rating,omitempty" validate:"min=1,max=5"`
	Comment string `json:"comment,omitempty" validate:"max=2000"`
}

type GetReviewRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteReviewRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
