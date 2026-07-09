package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CourseController struct {
	UseCase *usecase.CourseUseCase
	Log     *logrus.Logger
}

func NewCourseController(useCase *usecase.CourseUseCase, log *logrus.Logger) *CourseController {
	return &CourseController{UseCase: useCase, Log: log}
}

func (c *CourseController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateCourseRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}

	if request.InstructorId == "" {
		request.InstructorId = auth.ID
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error creating course")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}

func (c *CourseController) List(ctx *fiber.Ctx) error {
	req := &model.ListCourseRequest{
		Title:        ctx.Query("title", ""),
		CategoryId:   ctx.Query("category_id", ""),
		InstructorId: ctx.Query("instructor_id", ""),
		Status:       ctx.Query("status", ""),
		Page:         ctx.QueryInt("page", 1),
		Size:         ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.List(ctx.UserContext(), req)
	if err != nil {
		c.Log.WithError(err).Warn("error listing courses")
		return err
	}

	paging := &model.PageMetadata{
		Page:      req.Page,
		Size:      req.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(req.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.CourseResponse]{Data: responses, Paging: paging})
}

func (c *CourseController) Get(ctx *fiber.Ctx) error {
	request := &model.GetCourseRequest{ID: ctx.Params("courseId")}
	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error getting course")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}

func (c *CourseController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateCourseRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Warn("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("courseId")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warn("error updating course")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.CourseResponse]{Data: response})
}

func (c *CourseController) Delete(ctx *fiber.Ctx) error {
	request := &model.DeleteCourseRequest{ID: ctx.Params("courseId")}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Warn("error deleting course")
		return err
	}
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
