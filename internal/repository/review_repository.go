package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type ReviewRepository struct {
	Repository[entity.Review]
	Log *logrus.Logger
}

func NewReviewRepository(log *logrus.Logger) *ReviewRepository {
	return &ReviewRepository{Log: log}
}
