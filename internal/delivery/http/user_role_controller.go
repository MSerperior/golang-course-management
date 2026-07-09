package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserRoleController struct {
	UseCase *usecase.UserRoleUseCase
	Log     *logrus.Logger
}

func NewUserRoleController(useCase *usecase.UserRoleUseCase, log *logrus.Logger) *UserRoleController {
	return &UserRoleController{UseCase: useCase, Log: log}
}

func (c *UserRoleController) Create(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	request := new(model.CreateUserRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error assigning role")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.UserRoleResponse]{Data: response})
}

func (c *UserRoleController) ListByUser(ctx *fiber.Ctx) error {
	_ = middleware.GetUser(ctx)
	userId := ctx.Params("userId")
	responses, err := c.UseCase.ListByUser(ctx.UserContext(), userId)
	if err != nil {
		c.Log.WithError(err).Warn("error listing user roles")
		return err
	}
	return ctx.JSON(model.WebResponse[[]model.UserRoleResponse]{Data: responses})
}

func (c *UserRoleController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteUserRoleRequest{ID: ctx.Params("userRoleId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting user role")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
