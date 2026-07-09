package model

import "time"

// Transaction

type TransactionResponse struct {
	ID            string    `json:"id"`
	UserId        string    `json:"user_id,omitempty"`
	CourseId      string    `json:"course_id,omitempty"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Provider      string    `json:"provider"`
	PaymentStatus string    `json:"payment_status"`
	ExternalTxID  string    `json:"external_tx_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateTransactionRequest struct {
	UserId        string  `json:"user_id" validate:"required,max=100,uuid"`
	CourseId      string  `json:"course_id" validate:"required,max=100,uuid"`
	Amount        float64 `json:"amount" validate:"required"`
	Currency      string  `json:"currency" validate:"required,max=10"`
	Provider      string  `json:"provider" validate:"required,max=100"`
	PaymentStatus string  `json:"payment_status,omitempty" validate:"max=100"`
	ExternalTxID  string  `json:"external_tx_id,omitempty" validate:"max=255"`
}

type UpdateTransactionRequest struct {
	ID            string  `json:"-" validate:"required,max=100,uuid"`
	UserId        string  `json:"user_id,omitempty" validate:"max=100,uuid"`
	CourseId      string  `json:"course_id,omitempty" validate:"max=100,uuid"`
	Amount        float64 `json:"amount,omitempty"`
	Currency      string  `json:"currency,omitempty" validate:"max=10"`
	Provider      string  `json:"provider,omitempty" validate:"max=100"`
	PaymentStatus string  `json:"payment_status,omitempty" validate:"max=100"`
	ExternalTxID  string  `json:"external_tx_id,omitempty" validate:"max=255"`
}

type GetTransactionRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}

type DeleteTransactionRequest struct {
	ID string `json:"id" validate:"required,max=100,uuid"`
}
