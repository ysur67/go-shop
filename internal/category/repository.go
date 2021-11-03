package category

import (
	"context"
	"shop/models"
)

type Repository interface {
	AutoMigrate() error
	GetCategoryById(ctx context.Context, id string) (*models.Category, error)
}
