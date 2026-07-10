package entity

import (
	"time"
)

type BlacklistedToken struct {
	Entity
	TokenString string    `gorm:"unique;not null" json:"token_string"`
	ExpiredAt   time.Time `gorm:"not null" json:"expired_at"`
}
