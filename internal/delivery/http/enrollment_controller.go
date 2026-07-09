package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type EnrollmentController struct {
	UseCase *usecase.EnrollmentUseCase
	Log     *logrus.Logger
}

func NewEnrollmentController(useCase *usecase.EnrollmentUseCase, log *logrus.Logger) *EnrollmentController {
	return &EnrollmentController{UseCase: useCase, Log: log}
}

func (c *EnrollmentController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateEnrollmentRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating enrollment")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EnrollmentResponse]{Data: response})
}

func (c *EnrollmentController) Get(ctx *fiber.Ctx) error {
	request := &model.GetEnrollmentRequest{ID: ctx.Params("enrollmentId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting enrollment")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EnrollmentResponse]{Data: response})
}

func (c *EnrollmentController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateEnrollmentRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("enrollmentId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating enrollment")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.EnrollmentResponse]{Data: response})
}

func (c *EnrollmentController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteEnrollmentRequest{ID: ctx.Params("enrollmentId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting enrollment")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
