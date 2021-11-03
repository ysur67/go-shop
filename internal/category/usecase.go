package category

import (
	"context"
	"shop/models"
)

type UseCase interface {
	GetCategory(ctx context.Context, id string) (*models.Category, error)
}
