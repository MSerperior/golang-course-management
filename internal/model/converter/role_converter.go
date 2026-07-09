package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func RoleToResponse(role *entity.Role) *model.RoleResponse {
	var roleID string
	if role.ID != nil {
		roleID = role.ID.String()
	}

	return &model.RoleResponse{
		ID:        roleID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}
