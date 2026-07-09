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

type RoleUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	RoleRepository *repository.RoleRepository
}

func NewRoleUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	roleRepository *repository.RoleRepository) *RoleUseCase {
	return &RoleUseCase{DB: db, Log: logger, Validate: validate, RoleRepository: roleRepository}
}

func (c *RoleUseCase) Create(ctx context.Context, request *model.CreateRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid role request")
		return nil, fiber.ErrBadRequest
	}

	id := uuid.New()
	role := &entity.Role{Entity: entity.Entity{ID: &id}, Name: request.Name}
	if err := c.RoleRepository.Create(tx, role); err != nil {
		c.Log.WithError(err).Warn("failed to create role")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit role transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCase) Get(ctx context.Context, request *model.GetRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid role request")
		return nil, fiber.ErrBadRequest
	}
	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Warn("role not found")
		return nil, fiber.ErrNotFound
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit role transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCase) Update(ctx context.Context, request *model.UpdateRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid role request")
		return nil, fiber.ErrBadRequest
	}
	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Warn("role not found")
		return nil, fiber.ErrNotFound
	}
	if request.Name != "" {
		role.Name = request.Name
	}
	if err := c.RoleRepository.Update(tx, role); err != nil {
		c.Log.WithError(err).Warn("failed to update role")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit role transaction")
		return nil, fiber.ErrInternalServerError
	}
	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCase) Delete(ctx context.Context, request *model.DeleteRoleRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warn("invalid role request")
		return fiber.ErrBadRequest
	}
	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Warn("role not found")
		return fiber.ErrNotFound
	}
	if err := c.RoleRepository.Delete(tx, role); err != nil {
		c.Log.WithError(err).Warn("failed to delete role")
		return fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit role transaction")
		return fiber.ErrInternalServerError
	}
	return nil
}

func (c *RoleUseCase) List(ctx context.Context) ([]model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	var roles []entity.Role
	if err := tx.Find(&roles).Error; err != nil {
		c.Log.WithError(err).Warn("failed to list roles")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Warn("failed to commit role transaction")
		return nil, fiber.ErrInternalServerError
	}
	responses := make([]model.RoleResponse, len(roles))
	for i := range roles {
		responses[i] = *converter.RoleToResponse(&roles[i])
	}
	return responses, nil
}
