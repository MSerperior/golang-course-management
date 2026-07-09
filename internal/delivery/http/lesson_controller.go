package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LessonController struct {
	UseCase *usecase.LessonUseCase
	Log     *logrus.Logger
}

func NewLessonController(useCase *usecase.LessonUseCase, log *logrus.Logger) *LessonController {
	return &LessonController{UseCase: useCase, Log: log}
}

func (c *LessonController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateLessonRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating lesson")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.LessonResponse]{Data: response})
}

func (c *LessonController) Get(ctx *fiber.Ctx) error {
	request := &model.GetLessonRequest{ID: ctx.Params("lessonId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting lesson")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.LessonResponse]{Data: response})
}

func (c *LessonController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateLessonRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("lessonId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating lesson")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.LessonResponse]{Data: response})
}

func (c *LessonController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteLessonRequest{ID: ctx.Params("lessonId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting lesson")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
