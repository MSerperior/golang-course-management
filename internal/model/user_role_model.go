package model

import "time"

// UserRole

type UserRoleResponse struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id,omitempty"`
	RoleId    string    `json:"role_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRoleRequest struct {
	UserId string `json:"user_id" validate:"required,max=100,uuid"`
	RoleId string `json:"role_id" validate:"required,max=100,uuid"`
}

type UpdateUserRoleRequest struct {
	ID     string `json:"-" validate:"required,max=100,uuid"`
	UserId string `json:"user_id,omitempty" validate:"max=100,uuid"`
	RoleId string `json:"role_id,omitempty" validate:"max=100,uuid"`
}

type GetUserRoleRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteUserRoleRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
