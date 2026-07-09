package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type UserRoleRepository struct {
	Repository[entity.UserRole]
	Log *logrus.Logger
}

func NewUserRoleRepository(log *logrus.Logger) *UserRoleRepository {
	return &UserRoleRepository{Log: log}
}
