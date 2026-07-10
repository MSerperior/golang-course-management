package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestCreateCategoryRequestAllowsNilParentId(t *testing.T) {
	validate := validator.New()

	req := CreateCategoryRequest{
		Name: "Science",
		Slug: "science",
	}

	err := validate.Struct(req)
	require.NoError(t, err)
}
