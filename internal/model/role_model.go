package model

import "time"

// Role

type RoleResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type UpdateRoleRequest struct {
	ID   string `json:"-" validate:"required,max=100,uuid"`
	Name string `json:"name,omitempty" validate:"max=100"`
}

type GetRoleRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteRoleRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
