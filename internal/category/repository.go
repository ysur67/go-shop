package category

import (
	"context"
	"shop/models"
)

type Repository interface {
	AutoMigrate() error
	GetAll(ctx context.Context) (*[]models.Category, error)
	GetCategoryById(ctx context.Context, id string) (*models.Category, error)
}
