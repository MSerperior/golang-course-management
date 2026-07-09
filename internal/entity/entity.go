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
	CreatedBy string     `gorm:"column:created_by;default:null"`
	UpdatedBy string     `gorm:"column:updated_by;default:null"`
}
