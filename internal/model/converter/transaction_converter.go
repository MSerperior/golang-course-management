package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	var userID string
	if transaction.UserId != nil {
		userID = transaction.UserId.String()
	}

	var courseID string
	if transaction.CourseId != nil {
		courseID = transaction.CourseId.String()
	}

	var transactionID string
	if transaction.ID != nil {
		transactionID = transaction.ID.String()
	}

	return &model.TransactionResponse{
		ID:            transactionID,
		UserId:        userID,
		CourseId:      courseID,
		Amount:        transaction.Amount,
		Currency:      transaction.Currency,
		Provider:      transaction.Provider,
		PaymentStatus: transaction.PaymentStatus,
		ExternalTxID:  transaction.ExternalTxID,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}
