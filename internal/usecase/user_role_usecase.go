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

type UserRoleUseCase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	UserRoleRepository *repository.UserRoleRepository
	RoleRepository     *repository.RoleRepository
}

func NewUserRoleUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	urRepo *repository.UserRoleRepository, roleRepo *repository.RoleRepository) *UserRoleUseCase {
	return &UserRoleUseCase{DB: db, Log: logger, Validate: validate, UserRoleRepository: urRepo, RoleRepository: roleRepo}
}

func (c *UserRoleUseCase) Create(ctx context.Context, request *model.CreateUserRoleRequest) (*model.UserRoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid user role request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	ur := &entity.UserRole{Entity: entity.Entity{ID: &id}}
	if uid, err := uuid.Parse(request.UserId); err == nil {
		ur.UserId = &uid
	}
	if rid, err := uuid.Parse(request.RoleId); err == nil {
		ur.RoleId = &rid
	}

	if err := c.UserRoleRepository.AssignRole(tx, ur); err != nil {
		c.Log.WithError(err).Warn("failed to assign role")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit user role transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserRoleToResponse(ur), nil
}

func (c *UserRoleUseCase) Assign(ctx context.Context, request *model.CreateUserRoleRequest) (*model.UserRoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid user role request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	ur := &entity.UserRole{Entity: entity.Entity{ID: &id}}
	if uid, err := uuid.Parse(request.UserId); err == nil {
		ur.UserId = &uid
	}
	if rid, err := uuid.Parse(request.RoleId); err == nil {
		ur.RoleId = &rid
	}

	if err := c.UserRoleRepository.AssignRole(tx, ur); err != nil {
		c.Log.WithError(err).Warn("failed to assign role")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit assign role transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserRoleToResponse(ur), nil
}

func (c *UserRoleUseCase) ListByUser(ctx context.Context, userId string) ([]model.UserRoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	urs, err := c.UserRoleRepository.FindRolesByUserId(tx, userId)
	if err != nil {
		c.Log.WithError(err).Warn("failed to list user roles")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit list user roles transaction")
		return nil, fiber.ErrInternalServerError
	}

	resp := make([]model.UserRoleResponse, len(urs))
	for i := range urs {
		resp[i] = *converter.UserRoleToResponse(&urs[i])
	}
	return resp, nil
}

func (c *UserRoleUseCase) Delete(ctx context.Context, request *model.DeleteUserRoleRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid user role request")
		return fiber.ErrBadRequest
	}

	ur := new(entity.UserRole)
	if err := c.UserRoleRepository.FindById(tx, ur, request.ID); err != nil {
		c.Log.WithError(err).Warn("user role not found")
		return fiber.ErrNotFound
	}

	if err := c.UserRoleRepository.Delete(tx, ur); err != nil {
		c.Log.WithError(err).Warn("failed to delete user role")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit delete user role transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
