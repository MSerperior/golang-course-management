package entity

import "github.com/google/uuid"

// Category represents course categories (bigint PK)
type Category struct {
	Entity
	Name     string     `gorm:"column:name;unique"`
	Slug     string     `gorm:"column:slug;unique"`
	ParentId *uuid.UUID `gorm:"column:parent_id;type:varchar(36)"`

	Parent   *Category  `gorm:"foreignKey:parent_id;references:id"`
	Children []Category `gorm:"foreignKey:parent_id;references:id"`
	Courses  []Course   `gorm:"foreignKey:category_id;references:id"`
}

func (c *Category) TableName() string {
	return "categories"
}
