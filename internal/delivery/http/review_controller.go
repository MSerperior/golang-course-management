package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ReviewController struct {
	UseCase *usecase.ReviewUseCase
	Log     *logrus.Logger
}

func NewReviewController(useCase *usecase.ReviewUseCase, log *logrus.Logger) *ReviewController {
	return &ReviewController{UseCase: useCase, Log: log}
}

func (c *ReviewController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateReviewRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating review")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ReviewResponse]{Data: response})
}

func (c *ReviewController) Get(ctx *fiber.Ctx) error {
	request := &model.GetReviewRequest{ID: ctx.Params("reviewId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting review")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ReviewResponse]{Data: response})
}

func (c *ReviewController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateReviewRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("reviewId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating review")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.ReviewResponse]{Data: response})
}

func (c *ReviewController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteReviewRequest{ID: ctx.Params("reviewId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting review")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
