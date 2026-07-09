package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func CategoryToResponse(category *entity.Category) *model.CategoryResponse {
	var parentID string
	if category.ParentId != nil {
		parentID = category.ParentId.String()
	}

	var categoryID string
	if category.ID != nil {
		categoryID = category.ID.String()
	}

	return &model.CategoryResponse{
		ID:        categoryID,
		Name:      category.Name,
		Slug:      category.Slug,
		ParentId:  parentID,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
