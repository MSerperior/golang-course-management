package usecase

import (
	"context"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/model/converter"
	"golang-clean-architecture/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	transactionRepository *repository.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{DB: db, Log: logger, Validate: validate, TransactionRepository: transactionRepository}
}

func (c *TransactionUseCase) Create(ctx context.Context, request *model.CreateTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid transaction request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	txEnt := &entity.Transaction{Entity: entity.Entity{ID: &id}, Amount: request.Amount, Currency: request.Currency, Provider: request.Provider, PaymentStatus: request.PaymentStatus, ExternalTxID: request.ExternalTxID}
	if uid, err := uuid.Parse(request.UserId); err == nil {
		txEnt.UserId = &uid
	}
	if cid, err := uuid.Parse(request.CourseId); err == nil {
		txEnt.CourseId = &cid
	}

	if err := c.TransactionRepository.Create(tx, txEnt); err != nil {
		c.Log.WithError(err).Warn("failed to create transaction")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.TransactionToResponse(txEnt), nil
}

func (c *TransactionUseCase) Get(ctx context.Context, request *model.GetTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid transaction request")
		return nil, fiber.ErrBadRequest
	}
	t := new(entity.Transaction)
	if err := c.TransactionRepository.FindById(tx, t, request.ID); err != nil {
		c.Log.WithError(err).Warn("transaction not found")
		return nil, fiber.ErrNotFound
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.TransactionToResponse(t), nil
}

func (c *TransactionUseCase) Update(ctx context.Context, request *model.UpdateTransactionRequest) (*model.TransactionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid transaction request")
		return nil, fiber.ErrBadRequest
	}
	t := new(entity.Transaction)
	if err := c.TransactionRepository.FindById(tx, t, request.ID); err != nil {
		c.Log.WithError(err).Warn("transaction not found")
		return nil, fiber.ErrNotFound
	}
	if request.Amount != 0 {
		t.Amount = request.Amount
	}
	if request.Currency != "" {
		t.Currency = request.Currency
	}
	if request.Provider != "" {
		t.Provider = request.Provider
	}
	if request.PaymentStatus != "" {
		t.PaymentStatus = request.PaymentStatus
	}
	if request.ExternalTxID != "" {
		t.ExternalTxID = request.ExternalTxID
	}
	if request.UserId != "" {
		if uid, err := uuid.Parse(request.UserId); err == nil {
			t.UserId = &uid
		}
	}
	if request.CourseId != "" {
		if cid, err := uuid.Parse(request.CourseId); err == nil {
			t.CourseId = &cid
		}
	}
	if err := c.TransactionRepository.Update(tx, t); err != nil {
		c.Log.WithError(err).Warn("failed to update transaction")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.TransactionToResponse(t), nil
}

func (c *TransactionUseCase) Delete(ctx context.Context, request *model.DeleteTransactionRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid transaction request")
		return fiber.ErrBadRequest
	}
	t := new(entity.Transaction)
	if err := c.TransactionRepository.FindById(tx, t, request.ID); err != nil {
		c.Log.WithError(err).Warn("transaction not found")
		return fiber.ErrNotFound
	}
	if err := c.TransactionRepository.Delete(tx, t); err != nil {
		c.Log.WithError(err).Warn("failed to delete transaction")
		return fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit transaction")
		return fiber.ErrInternalServerError
	}
	return nil
}
