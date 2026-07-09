package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoleController struct {
	UseCase *usecase.RoleUseCase
	Log     *logrus.Logger
}

func NewRoleController(useCase *usecase.RoleUseCase, log *logrus.Logger) *RoleController {
	return &RoleController{UseCase: useCase, Log: log}
}

func (c *RoleController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating role")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

func (c *RoleController) List(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	responses, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Warn("error listing roles")
		return err
	}
	// simple paging metadata
	paging := &model.PageMetadata{Page: 1, Size: len(responses), TotalItem: int64(len(responses)), TotalPage: 1}
	return ctx.JSON(model.WebResponse[[]model.RoleResponse]{Data: responses, Paging: paging})
}

func (c *RoleController) Get(ctx *fiber.Ctx) error {
	request := &model.GetRoleRequest{ID: ctx.Params("roleId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting role")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

func (c *RoleController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("roleId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating role")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

func (c *RoleController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteRoleRequest{ID: ctx.Params("roleId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting role")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
