package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func CourseToResponse(course *entity.Course) *model.CourseResponse {
	var instructorID string
	if course.InstructorId != nil {
		instructorID = course.InstructorId.String()
	}

	var categoryID string
	if course.CategoryId != nil {
		categoryID = course.CategoryId.String()
	}

	var courseID string
	if course.ID != nil {
		courseID = course.ID.String()
	}

	return &model.CourseResponse{
		ID:           courseID,
		InstructorId: instructorID,
		CategoryId:   categoryID,
		Title:        course.Title,
		Slug:         course.Slug,
		Description:  course.Description,
		Price:        course.Price,
		Status:       course.Status,
		CreatedAt:    course.CreatedAt,
		UpdatedAt:    course.UpdatedAt,
	}
}
