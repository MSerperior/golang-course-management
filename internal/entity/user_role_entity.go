package entity

import "github.com/google/uuid"

// UserRole links users and roles (bigint PK)
type UserRole struct {
	Entity
	UserId *uuid.UUID `gorm:"column:user_id;type:varchar(36)"`
	RoleId *uuid.UUID `gorm:"column:role_id;type:varchar(36)"`

	User *User `gorm:"foreignKey:user_id;references:id"`
	Role *Role `gorm:"foreignKey:role_id;references:id"`
}

func (ur *UserRole) TableName() string {
	return "user_roles"
}
