package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func EnrollmentToResponse(enrollment *entity.Enrollment) *model.EnrollmentResponse {
	var studentID string
	if enrollment.StudentId != nil {
		studentID = enrollment.StudentId.String()
	}

	var courseID string
	if enrollment.CourseId != nil {
		courseID = enrollment.CourseId.String()
	}

	var enrollmentID string
	if enrollment.ID != nil {
		enrollmentID = enrollment.ID.String()
	}

	return &model.EnrollmentResponse{
		ID:          enrollmentID,
		StudentId:   studentID,
		CourseId:    courseID,
		Status:      enrollment.Status,
		EnrolledAt:  enrollment.EnrolledAt,
		CompletedAt: enrollment.CompletedAt,
		CreatedAt:   enrollment.CreatedAt,
		UpdatedAt:   enrollment.UpdatedAt,
	}
}
