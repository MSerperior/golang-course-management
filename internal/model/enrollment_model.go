package model

import "time"

// Enrollment

type EnrollmentResponse struct {
	ID          string    `json:"id"`
	StudentId   string    `json:"student_id,omitempty"`
	CourseId    string    `json:"course_id,omitempty"`
	Status      string    `json:"status"`
	EnrolledAt  time.Time `json:"enrolled_at"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateEnrollmentRequest struct {
	StudentId   string    `json:"student_id" validate:"required,max=100,uuid"`
	CourseId    string    `json:"course_id" validate:"required,max=100,uuid"`
	Status      string    `json:"status,omitempty" validate:"max=50"`
	EnrolledAt  time.Time `json:"enrolled_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

type UpdateEnrollmentRequest struct {
	ID          string    `json:"-" validate:"required,max=100,uuid"`
	StudentId   string    `json:"student_id,omitempty" validate:"max=100,uuid"`
	CourseId    string    `json:"course_id,omitempty" validate:"max=100,uuid"`
	Status      string    `json:"status,omitempty" validate:"max=50"`
	EnrolledAt  time.Time `json:"enrolled_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

type GetEnrollmentRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteEnrollmentRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
