package usecase

import (
	"context"
	"shop/internal/category"
	"shop/models"
)

type UseCase struct {
	repo category.Repository
}

func NewUseCase(repo category.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (us *UseCase) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	return us.repo.GetCategoryById(ctx, id)
}
