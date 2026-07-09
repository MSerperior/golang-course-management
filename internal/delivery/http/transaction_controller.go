package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	UseCase *usecase.TransactionUseCase
	Log     *logrus.Logger
}

func NewTransactionController(useCase *usecase.TransactionUseCase, log *logrus.Logger) *TransactionController {
	return &TransactionController{UseCase: useCase, Log: log}
}

func (c *TransactionController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating transaction")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) Get(ctx *fiber.Ctx) error {
	request := &model.GetTransactionRequest{ID: ctx.Params("transactionId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting transaction")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("transactionId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating transaction")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteTransactionRequest{ID: ctx.Params("transactionId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting transaction")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
