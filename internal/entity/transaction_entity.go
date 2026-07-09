package entity

import "github.com/google/uuid"

// Transaction represents a purchase/payment (UUID PK)
type Transaction struct {
	Entity
	UserId        *uuid.UUID `gorm:"column:user_id;type:varchar(36)"`
	CourseId      *uuid.UUID `gorm:"column:course_id;type:varchar(36)"`
	Amount        float64    `gorm:"column:amount"`
	Currency      string     `gorm:"column:currency"`
	Provider      string     `gorm:"column:provider"`
	PaymentStatus string     `gorm:"column:payment_status"`
	ExternalTxID  string     `gorm:"column:external_tx_id;unique"`

	User   *User   `gorm:"foreignKey:user_id;references:id"`
	Course *Course `gorm:"foreignKey:course_id;references:id"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
