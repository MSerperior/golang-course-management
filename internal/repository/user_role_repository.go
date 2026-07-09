package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	Repository[entity.UserRole]
	Log *logrus.Logger
}

func NewUserRoleRepository(log *logrus.Logger) *UserRoleRepository {
	return &UserRoleRepository{Log: log}
}

func (r *UserRoleRepository) FindRolesByUserId(db *gorm.DB, userId string) ([]entity.UserRole, error) {
	var urs []entity.UserRole
	if err := db.Where("user_id = ?", userId).Find(&urs).Error; err != nil {
		return nil, err
	}
	return urs, nil
}

func (r *UserRoleRepository) AssignRole(tx *gorm.DB, ur *entity.UserRole) error {
	return tx.Create(ur).Error
}
