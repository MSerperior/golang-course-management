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

type CategoryUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	categoryRepository *repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryUseCase) Create(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid category request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	category := &entity.Category{
		Entity: entity.Entity{ID: &id},
		Name:   request.Name,
		Slug:   request.Slug,
	}

	if request.ParentId != nil && *request.ParentId != "" {
		parentID, err := uuid.Parse(*request.ParentId)
		if err != nil {
			c.Log.WithError(err).Warn("invalid parent category id")
			return nil, fiber.ErrBadRequest
		}
		category.ParentId = &parentID
	}

	if err := c.CategoryRepository.Create(tx, category); err != nil {
		c.Log.WithError(err).Warn("failed to create category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit category transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCase) Update(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid category request")
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find category")
		return nil, fiber.ErrNotFound
	}

	if request.Name != "" {
		category.Name = request.Name
	}
	if request.Slug != "" {
		category.Slug = request.Slug
	}
	if request.ParentId != nil && *request.ParentId != "" {
		parentID, err := uuid.Parse(*request.ParentId)
		if err != nil {
			c.Log.WithError(err).Warn("invalid parent category id")
			return nil, fiber.ErrBadRequest
		}
		category.ParentId = &parentID
	}

	if err := c.CategoryRepository.Update(tx, category); err != nil {
		c.Log.WithError(err).Warn("failed to update category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit category transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCase) Get(ctx context.Context, request *model.GetCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid category request")
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find category")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit category transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUseCase) Delete(ctx context.Context, request *model.DeleteCategoryRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid category request")
		return fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.WithError(err).Warn("failed to find category")
		return fiber.ErrNotFound
	}

	if err := c.CategoryRepository.Delete(tx, category); err != nil {
		c.Log.WithError(err).Warn("failed to delete category")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit category transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CategoryUseCase) List(ctx context.Context) ([]model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var categories []entity.Category
	if err := tx.Find(&categories).Error; err != nil {
		c.Log.WithError(err).Warn("failed to list categories")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit category transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *converter.CategoryToResponse(&category)
	}

	return responses, nil
}
