package usecase

import (
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

func (us *UseCase) GetAll() (*[]models.Category, error) {
	return us.repo.GetAll()
}

func (us *UseCase) GetCategory(id string) (*models.Category, error) {
	return us.repo.GetCategoryById(id)
}
