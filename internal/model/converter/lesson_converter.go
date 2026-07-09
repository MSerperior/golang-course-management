package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func LessonToResponse(lesson *entity.Lesson) *model.LessonResponse {
	var sectionID string
	if lesson.SectionId != nil {
		sectionID = lesson.SectionId.String()
	}

	var lessonID string
	if lesson.ID != nil {
		lessonID = lesson.ID.String()
	}

	return &model.LessonResponse{
		ID:              lessonID,
		SectionId:       sectionID,
		Title:           lesson.Title,
		Type:            lesson.Type,
		ContentURL:      lesson.ContentURL,
		DurationSeconds: lesson.DurationSeconds,
		SortOrder:       lesson.SortOrder,
		CreatedAt:       lesson.CreatedAt,
		UpdatedAt:       lesson.UpdatedAt,
	}
}
