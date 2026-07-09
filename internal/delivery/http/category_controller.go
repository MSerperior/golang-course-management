package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	UseCase *usecase.CategoryUseCase
	Log     *logrus.Logger
}

func NewCategoryController(useCase *usecase.CategoryUseCase, log *logrus.Logger) *CategoryController {
	return &CategoryController{UseCase: useCase, Log: log}
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("failed to parse category request")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("failed to create category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	responses, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Warn("failed to list categories")
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CategoryResponse]{Data: responses})
}

func (c *CategoryController) Get(ctx *fiber.Ctx) error {
	request := &model.GetCategoryRequest{ID: ctx.Params("categoryId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("failed to get category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateCategoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("failed to parse category update request")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("categoryId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("failed to update category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteCategoryRequest{ID: ctx.Params("categoryId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("failed to delete category")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
