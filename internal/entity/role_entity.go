package entity

// Role represents a role with a bigint primary key
type Role struct {
	Entity
	Name string `gorm:"column:name;unique"`

	UserRoles []UserRole `gorm:"foreignKey:role_id;references:id"`
}

func (r *Role) TableName() string {
	return "roles"
}
