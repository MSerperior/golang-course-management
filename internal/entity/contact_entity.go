package entity

import "github.com/google/uuid"

type Contact struct {
	Entity
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	Email     string     `gorm:"column:email"`
	Phone     string     `gorm:"column:phone"`
	UserId    *uuid.UUID `gorm:"column:user_id;type:varchar(36)"`
	User      User       `gorm:"foreignKey:user_id;references:id"`
	Addresses []Address  `gorm:"foreignKey:contact_id;references:id"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
