package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type RoleRepository struct {
	Repository[entity.Role]
	Log *logrus.Logger
}

func NewRoleRepository(log *logrus.Logger) *RoleRepository {
	return &RoleRepository{Log: log}
}
