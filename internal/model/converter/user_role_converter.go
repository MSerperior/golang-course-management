package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func UserRoleToResponse(userRole *entity.UserRole) *model.UserRoleResponse {
	var userID string
	if userRole.UserId != nil {
		userID = userRole.UserId.String()
	}

	var roleID string
	if userRole.RoleId != nil {
		roleID = userRole.RoleId.String()
	}

	var userRoleID string
	if userRole.ID != nil {
		userRoleID = userRole.ID.String()
	}

	return &model.UserRoleResponse{
		ID:        userRoleID,
		UserId:    userID,
		RoleId:    roleID,
		CreatedAt: userRole.CreatedAt,
		UpdatedAt: userRole.UpdatedAt,
	}
}
