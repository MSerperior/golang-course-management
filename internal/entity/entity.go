package entity

import (
	"time"

	"github.com/google/uuid"
)

// Entity is a base struct that other entities can embed
type Entity struct {
	ID        *uuid.UUID `gorm:"column:id;primaryKey;type:varchar(36)"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy *uuid.UUID `gorm:"column:created_by;default:null;type:varchar(36)"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by;default:null;type:varchar(36)"`

	Creator *User `gorm:"foreignKey:created_by;references:id"`
	Updater *User `gorm:"foreignKey:updated_by;references:id"`
}
