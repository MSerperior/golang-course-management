package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func SectionToResponse(section *entity.Section) *model.SectionResponse {
	var courseID string
	if section.CourseId != nil {
		courseID = section.CourseId.String()
	}

	var sectionID string
	if section.ID != nil {
		sectionID = section.ID.String()
	}

	return &model.SectionResponse{
		ID:        sectionID,
		CourseId:  courseID,
		Title:     section.Title,
		SortOrder: section.SortOrder,
		CreatedAt: section.CreatedAt,
		UpdatedAt: section.UpdatedAt,
	}
}
