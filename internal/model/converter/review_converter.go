package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func ReviewToResponse(review *entity.Review) *model.ReviewResponse {
	var courseID string
	if review.CourseId != nil {
		courseID = review.CourseId.String()
	}

	var studentID string
	if review.StudentId != nil {
		studentID = review.StudentId.String()
	}

	var reviewID string
	if review.ID != nil {
		reviewID = review.ID.String()
	}

	return &model.ReviewResponse{
		ID:        reviewID,
		CourseId:  courseID,
		StudentId: studentID,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}
