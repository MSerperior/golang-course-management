package model

import "time"

// Category

type CategoryResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	ParentId  string    `json:"parent_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Slug     string `json:"slug" validate:"required,max=100"`
	ParentId string `json:"parent_id,omitempty" validate:"max=100,uuid"`
}

type UpdateCategoryRequest struct {
	ID       string `json:"-" validate:"required,max=100,uuid"`
	Name     string `json:"name,omitempty" validate:"max=100"`
	Slug     string `json:"slug,omitempty" validate:"max=100"`
	ParentId string `json:"parent_id,omitempty" validate:"max=100,uuid"`
}

type GetCategoryRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
