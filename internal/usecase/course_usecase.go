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

type CourseUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	CourseRepository *repository.CourseRepository
}

func NewCourseUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	courseRepository *repository.CourseRepository) *CourseUseCase {
	return &CourseUseCase{
		DB:               db,
		Log:              logger,
		Validate:         validate,
		CourseRepository: courseRepository,
	}
}

func (c *CourseUseCase) Create(ctx context.Context, request *model.CreateCourseRequest) (*model.CourseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid course request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	course := &entity.Course{
		Entity:      entity.Entity{ID: &id},
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Price:       request.Price,
		Status:      request.Status,
	}

	if request.InstructorId != "" {
		if instructorID, err := uuid.Parse(request.InstructorId); err == nil {
			course.InstructorId = &instructorID
		}
	}
	if request.CategoryId != "" {
		if categoryID, err := uuid.Parse(request.CategoryId); err == nil {
			course.CategoryId = &categoryID
		}
	}

	if err := c.CourseRepository.Create(tx, course); err != nil {
		c.Log.WithError(err).Warn("failed to create course")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit course transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CourseToResponse(course), nil
}

func (c *CourseUseCase) Get(ctx context.Context, request *model.GetCourseRequest) (*model.CourseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid course request")
		return nil, fiber.ErrBadRequest
	}

	course := new(entity.Course)
	if err := c.CourseRepository.FindById(tx, course, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find course")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit course transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CourseToResponse(course), nil
}

func (c *CourseUseCase) Update(ctx context.Context, request *model.UpdateCourseRequest) (*model.CourseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid course request")
		return nil, fiber.ErrBadRequest
	}

	course := new(entity.Course)
	if err := c.CourseRepository.FindById(tx, course, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find course")
		return nil, fiber.ErrNotFound
	}

	if request.Title != "" {
		course.Title = request.Title
	}
	if request.Slug != "" {
		course.Slug = request.Slug
	}
	if request.Description != "" {
		course.Description = request.Description
	}
	if request.Price != 0 {
		course.Price = request.Price
	}
	if request.Status != "" {
		course.Status = request.Status
	}
	if request.InstructorId != "" {
		if instructorID, err := uuid.Parse(request.InstructorId); err == nil {
			course.InstructorId = &instructorID
		}
	}
	if request.CategoryId != "" {
		if categoryID, err := uuid.Parse(request.CategoryId); err == nil {
			course.CategoryId = &categoryID
		}
	}

	if err := c.CourseRepository.Update(tx, course); err != nil {
		c.Log.WithError(err).Warn("failed to update course")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit course transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CourseToResponse(course), nil
}

func (c *CourseUseCase) Delete(ctx context.Context, request *model.DeleteCourseRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid course request")
		return fiber.ErrBadRequest
	}

	course := new(entity.Course)
	if err := c.CourseRepository.FindById(tx, course, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find course")
		return fiber.ErrNotFound
	}

	if err := c.CourseRepository.Delete(tx, course); err != nil {
		c.Log.WithError(err).Warn("failed to delete course")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit course transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CourseUseCase) List(ctx context.Context, request *model.ListCourseRequest) ([]model.CourseResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if request != nil {
		if err := c.Validate.Struct(request); err != nil {
			c.Log.WithError(err).Warn("invalid list course request")
			return nil, 0, fiber.ErrBadRequest
		}
	}

	var courses []entity.Course
	var total int64

	// If instructor filter provided, use repository helper
	if request != nil && request.InstructorId != "" {
		page := request.Page
		size := request.Size
		if page == 0 {
			page = 1
		}
		if size == 0 {
			size = 10
		}
		cs, t, err := c.CourseRepository.FindAllByInstructorId(tx, request.InstructorId, page, size)
		if err != nil {
			c.Log.WithError(err).Warn("failed to list courses by instructor")
			return nil, 0, fiber.ErrInternalServerError
		}
		responses := make([]model.CourseResponse, len(cs))
		for i := range cs {
			responses[i] = *converter.CourseToResponse(&cs[i])
		}
		return responses, t, nil
	}

	q := tx.Model(&entity.Course{})
	if request != nil && request.Title != "" {
		q = q.Where("title ILIKE ?", "%"+request.Title+"%")
	}
	if request != nil && request.CategoryId != "" {
		if categoryID, err := uuid.Parse(request.CategoryId); err == nil {
			q = q.Where("category_id = ?", categoryID)
		}
	}
	if request != nil && request.Status != "" {
		q = q.Where("status = ?", request.Status)
	}

	if request != nil && request.Page > 0 && request.Size > 0 {
		q = q.Offset((request.Page - 1) * request.Size).Limit(request.Size)
	} else {
		q = q.Limit(100)
	}

	if err := q.Find(&courses).Error; err != nil {
		c.Log.WithError(err).Warn("failed to list courses")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Model(&entity.Course{}).Count(&total).Error; err != nil {
		c.Log.WithError(err).Warn("failed to count courses")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit course transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CourseResponse, len(courses))
	for i := range courses {
		responses[i] = *converter.CourseToResponse(&courses[i])
	}

	return responses, total, nil
}
