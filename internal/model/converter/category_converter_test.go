package converter

import (
	"testing"
	"time"

	"golang-clean-architecture/internal/entity"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCategoryToResponse(t *testing.T) {
	id := uuid.New()
	parentID := uuid.New()
	category := &entity.Category{
		Entity: entity.Entity{
			ID:        &id,
			CreatedAt: time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC),
			UpdatedAt: time.Date(2024, 1, 2, 3, 4, 6, 0, time.UTC),
		},
		Name:     "Science",
		Slug:     "science",
		ParentId: &parentID,
	}

	response := CategoryToResponse(category)

	require.Equal(t, id.String(), response.ID)
	require.Equal(t, "Science", response.Name)
	require.Equal(t, "science", response.Slug)
	require.Equal(t, parentID.String(), response.ParentId)
	require.Equal(t, category.CreatedAt, response.CreatedAt)
	require.Equal(t, category.UpdatedAt, response.UpdatedAt)
}
