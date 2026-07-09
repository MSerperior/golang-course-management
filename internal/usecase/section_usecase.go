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

type SectionUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	SectionRepository *repository.SectionRepository
}

func NewSectionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	sectionRepository *repository.SectionRepository) *SectionUseCase {
	return &SectionUseCase{DB: db, Log: logger, Validate: validate, SectionRepository: sectionRepository}
}

func (c *SectionUseCase) Create(ctx context.Context, request *model.CreateSectionRequest) (*model.SectionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid section request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	section := &entity.Section{Entity: entity.Entity{ID: &id}, Title: request.Title, SortOrder: request.SortOrder}
	if cid, err := uuid.Parse(request.CourseId); err == nil {
		section.CourseId = &cid
	}

	if err := c.SectionRepository.Create(tx, section); err != nil {
		c.Log.WithError(err).Warn("failed to create section")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit section transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SectionToResponse(section), nil
}

func (c *SectionUseCase) Get(ctx context.Context, request *model.GetSectionRequest) (*model.SectionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid section request")
		return nil, fiber.ErrBadRequest
	}

	section := new(entity.Section)
	if err := c.SectionRepository.FindById(tx, section, request.ID); err != nil {
		c.Log.WithError(err).Warn("section not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit section transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SectionToResponse(section), nil
}

func (c *SectionUseCase) Update(ctx context.Context, request *model.UpdateSectionRequest) (*model.SectionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid section request")
		return nil, fiber.ErrBadRequest
	}

	section := new(entity.Section)
	if err := c.SectionRepository.FindById(tx, section, request.ID); err != nil {
		c.Log.WithError(err).Warn("section not found")
		return nil, fiber.ErrNotFound
	}

	if request.Title != "" {
		section.Title = request.Title
	}
	if request.SortOrder != 0 {
		section.SortOrder = request.SortOrder
	}
	if request.CourseId != "" {
		if cid, err := uuid.Parse(request.CourseId); err == nil {
			section.CourseId = &cid
		}
	}

	if err := c.SectionRepository.Update(tx, section); err != nil {
		c.Log.WithError(err).Warn("failed to update section")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit section transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SectionToResponse(section), nil
}

func (c *SectionUseCase) Delete(ctx context.Context, request *model.DeleteSectionRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid section request")
		return fiber.ErrBadRequest
	}

	section := new(entity.Section)
	if err := c.SectionRepository.FindById(tx, section, request.ID); err != nil {
		c.Log.WithError(err).Warn("section not found")
		return fiber.ErrNotFound
	}

	if err := c.SectionRepository.Delete(tx, section); err != nil {
		c.Log.WithError(err).Warn("failed to delete section")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit section transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
