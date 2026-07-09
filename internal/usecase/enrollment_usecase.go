package usecase

import (
	"context"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/model/converter"
	"golang-clean-architecture/internal/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EnrollmentUseCase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	EnrollmentRepository *repository.EnrollmentRepository
}

func NewEnrollmentUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	enrollmentRepository *repository.EnrollmentRepository) *EnrollmentUseCase {
	return &EnrollmentUseCase{DB: db, Log: logger, Validate: validate, EnrollmentRepository: enrollmentRepository}
}

func (c *EnrollmentUseCase) Create(ctx context.Context, request *model.CreateEnrollmentRequest) (*model.EnrollmentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid enrollment request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	enrollment := &entity.Enrollment{
		Entity:      entity.Entity{ID: &id},
		Status:      request.Status,
		EnrolledAt:  request.EnrolledAt,
		CompletedAt: request.CompletedAt,
	}
	if sid, err := uuid.Parse(request.StudentId); err == nil {
		enrollment.StudentId = &sid
	}
	if cid, err := uuid.Parse(request.CourseId); err == nil {
		enrollment.CourseId = &cid
	}

	if enrollment.EnrolledAt.IsZero() {
		enrollment.EnrolledAt = time.Now()
	}

	if err := c.EnrollmentRepository.Create(tx, enrollment); err != nil {
		c.Log.WithError(err).Warn("failed to create enrollment")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit enrollment transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.EnrollmentToResponse(enrollment), nil
}

func (c *EnrollmentUseCase) Get(ctx context.Context, request *model.GetEnrollmentRequest) (*model.EnrollmentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid enrollment request")
		return nil, fiber.ErrBadRequest
	}

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepository.FindById(tx, enrollment, request.ID); err != nil {
		c.Log.WithError(err).Warn("enrollment not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit enrollment transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.EnrollmentToResponse(enrollment), nil
}

func (c *EnrollmentUseCase) Update(ctx context.Context, request *model.UpdateEnrollmentRequest) (*model.EnrollmentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid enrollment request")
		return nil, fiber.ErrBadRequest
	}

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepository.FindById(tx, enrollment, request.ID); err != nil {
		c.Log.WithError(err).Warn("enrollment not found")
		return nil, fiber.ErrNotFound
	}

	if request.Status != "" {
		enrollment.Status = request.Status
	}
	if !request.EnrolledAt.IsZero() {
		enrollment.EnrolledAt = request.EnrolledAt
	}
	if !request.CompletedAt.IsZero() {
		enrollment.CompletedAt = request.CompletedAt
	}
	if request.StudentId != "" {
		if sid, err := uuid.Parse(request.StudentId); err == nil {
			enrollment.StudentId = &sid
		}
	}
	if request.CourseId != "" {
		if cid, err := uuid.Parse(request.CourseId); err == nil {
			enrollment.CourseId = &cid
		}
	}

	if err := c.EnrollmentRepository.Update(tx, enrollment); err != nil {
		c.Log.WithError(err).Warn("failed to update enrollment")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit enrollment transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.EnrollmentToResponse(enrollment), nil
}

func (c *EnrollmentUseCase) Delete(ctx context.Context, request *model.DeleteEnrollmentRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid enrollment request")
		return fiber.ErrBadRequest
	}

	enrollment := new(entity.Enrollment)
	if err := c.EnrollmentRepository.FindById(tx, enrollment, request.ID); err != nil {
		c.Log.WithError(err).Warn("enrollment not found")
		return fiber.ErrNotFound
	}

	if err := c.EnrollmentRepository.Delete(tx, enrollment); err != nil {
		c.Log.WithError(err).Warn("failed to delete enrollment")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit enrollment transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
