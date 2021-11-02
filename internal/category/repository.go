package category

import (
	"context"
	"shop/models"
)

type Repository interface {
	GetCategoryById(ctx context.Context, id string) (models.Category, error)
}
