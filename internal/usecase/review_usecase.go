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

type ReviewUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	ReviewRepository *repository.ReviewRepository
}

func NewReviewUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	reviewRepository *repository.ReviewRepository) *ReviewUseCase {
	return &ReviewUseCase{DB: db, Log: logger, Validate: validate, ReviewRepository: reviewRepository}
}

func (c *ReviewUseCase) Create(ctx context.Context, request *model.CreateReviewRequest) (*model.ReviewResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid review request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	review := &entity.Review{Entity: entity.Entity{ID: &id}, Rating: request.Rating, Comment: request.Comment}
	if cid, err := uuid.Parse(request.CourseId); err == nil {
		review.CourseId = &cid
	}
	if sid, err := uuid.Parse(request.StudentId); err == nil {
		review.StudentId = &sid
	}

	if err := c.ReviewRepository.Create(tx, review); err != nil {
		c.Log.WithError(err).Warn("failed to create review")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit review transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ReviewToResponse(review), nil
}

func (c *ReviewUseCase) Get(ctx context.Context, request *model.GetReviewRequest) (*model.ReviewResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid review request")
		return nil, fiber.ErrBadRequest
	}

	review := new(entity.Review)
	if err := c.ReviewRepository.FindById(tx, review, request.ID); err != nil {
		c.Log.WithError(err).Warn("review not found")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit review transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ReviewToResponse(review), nil
}

func (c *ReviewUseCase) Update(ctx context.Context, request *model.UpdateReviewRequest) (*model.ReviewResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid review request")
		return nil, fiber.ErrBadRequest
	}

	review := new(entity.Review)
	if err := c.ReviewRepository.FindById(tx, review, request.ID); err != nil {
		c.Log.WithError(err).Warn("review not found")
		return nil, fiber.ErrNotFound
	}

	if request.Rating != 0 {
		review.Rating = request.Rating
	}
	if request.Comment != "" {
		review.Comment = request.Comment
	}

	if err := c.ReviewRepository.Update(tx, review); err != nil {
		c.Log.WithError(err).Warn("failed to update review")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit review transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ReviewToResponse(review), nil
}

func (c *ReviewUseCase) Delete(ctx context.Context, request *model.DeleteReviewRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid review request")
		return fiber.ErrBadRequest
	}

	review := new(entity.Review)
	if err := c.ReviewRepository.FindById(tx, review, request.ID); err != nil {
		c.Log.WithError(err).Warn("review not found")
		return fiber.ErrNotFound
	}

	if err := c.ReviewRepository.Delete(tx, review); err != nil {
		c.Log.WithError(err).Warn("failed to delete review")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit review transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
