package category

import (
	"context"
	"shop/models"
)

type UseCase interface {
	GetAll(ctx context.Context) (*[]models.Category, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
}
