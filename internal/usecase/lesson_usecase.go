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

type LessonUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	LessonRepository  *repository.LessonRepository
	SectionRepository *repository.SectionRepository
}

func NewLessonUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	lessonRepository *repository.LessonRepository, sectionRepository *repository.SectionRepository) *LessonUseCase {
	return &LessonUseCase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		LessonRepository:  lessonRepository,
		SectionRepository: sectionRepository,
	}
}

func (c *LessonUseCase) Create(ctx context.Context, request *model.CreateLessonRequest) (*model.LessonResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid lesson request")
		return nil, fiber.ErrBadRequest
	}

	// ensure section exists
	section := new(entity.Section)
	if err := c.SectionRepository.FindById(tx, section, request.SectionId); err != nil {
		c.Log.WithError(err).Warn("section not found")
		return nil, fiber.ErrNotFound
	}

	id := uuid.New()
	lesson := &entity.Lesson{
		Entity:          entity.Entity{ID: &id},
		SectionId:       section.ID,
		Title:           request.Title,
		Type:            request.Type,
		ContentURL:      request.ContentURL,
		DurationSeconds: request.DurationSeconds,
		SortOrder:       request.SortOrder,
	}

	if err := c.LessonRepository.Create(tx, lesson); err != nil {
		c.Log.WithError(err).Warn("failed to create lesson")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit lesson transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LessonToResponse(lesson), nil
}

func (c *LessonUseCase) Get(ctx context.Context, request *model.GetLessonRequest) (*model.LessonResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid lesson request")
		return nil, fiber.ErrBadRequest
	}

	lesson := new(entity.Lesson)
	if err := c.LessonRepository.FindById(tx, lesson, request.ID); err != nil {
		c.Log.WithError(err).Warn("lesson not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit lesson transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LessonToResponse(lesson), nil
}

func (c *LessonUseCase) Update(ctx context.Context, request *model.UpdateLessonRequest) (*model.LessonResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid lesson request")
		return nil, fiber.ErrBadRequest
	}

	lesson := new(entity.Lesson)
	if err := c.LessonRepository.FindById(tx, lesson, request.ID); err != nil {
		c.Log.WithError(err).Warn("lesson not found")
		return nil, fiber.ErrNotFound
	}

	if request.SectionId != "" {
		if sectionID, err := uuid.Parse(request.SectionId); err == nil {
			lesson.SectionId = &sectionID
		}
	}
	if request.Title != "" {
		lesson.Title = request.Title
	}
	if request.Type != "" {
		lesson.Type = request.Type
	}
	if request.ContentURL != "" {
		lesson.ContentURL = request.ContentURL
	}
	if request.DurationSeconds != 0 {
		lesson.DurationSeconds = request.DurationSeconds
	}
	if request.SortOrder != 0 {
		lesson.SortOrder = request.SortOrder
	}

	if err := c.LessonRepository.Update(tx, lesson); err != nil {
		c.Log.WithError(err).Warn("failed to update lesson")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit lesson transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.LessonToResponse(lesson), nil
}

func (c *LessonUseCase) Delete(ctx context.Context, request *model.DeleteLessonRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid lesson request")
		return fiber.ErrBadRequest
	}

	lesson := new(entity.Lesson)
	if err := c.LessonRepository.FindById(tx, lesson, request.ID); err != nil {
		c.Log.WithError(err).Warn("lesson not found")
		return fiber.ErrNotFound
	}

	if err := c.LessonRepository.Delete(tx, lesson); err != nil {
		c.Log.WithError(err).Warn("failed to delete lesson")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit lesson transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
