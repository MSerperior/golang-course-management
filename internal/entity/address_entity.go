package entity

import "github.com/google/uuid"

type Address struct {
	Entity
	ContactId  *uuid.UUID `gorm:"column:contact_id;type:varchar(36)"`
	Street     string     `gorm:"column:street"`
	City       string     `gorm:"column:city"`
	Province   string     `gorm:"column:province"`
	PostalCode string     `gorm:"column:postal_code"`
	Country    string     `gorm:"column:country"`
}

func (a *Address) TableName() string {
	return "addresses"
}
