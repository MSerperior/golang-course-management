package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SectionController struct {
	UseCase *usecase.SectionUseCase
	Log     *logrus.Logger
}

func NewSectionController(useCase *usecase.SectionUseCase, log *logrus.Logger) *SectionController {
	return &SectionController{UseCase: useCase, Log: log}
}

func (c *SectionController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateSectionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating section")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.SectionResponse]{Data: response})
}

func (c *SectionController) Get(ctx *fiber.Ctx) error {
	request := &model.GetSectionRequest{ID: ctx.Params("sectionId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting section")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.SectionResponse]{Data: response})
}

func (c *SectionController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateSectionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("sectionId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating section")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.SectionResponse]{Data: response})
}

func (c *SectionController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteSectionRequest{ID: ctx.Params("sectionId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting section")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
