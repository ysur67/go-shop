package usecase

import (
	"shop/internal/category"
	"shop/internal/product"
	"shop/models"
)

type UseCase struct {
	repo         product.Repository
	categoryRepo category.Repository
}

func NewUseCase(repo product.Repository, catRepo category.Repository) *UseCase {
	return &UseCase{
		repo:         repo,
		categoryRepo: catRepo,
	}
}

func (us *UseCase) GetAll() (*[]models.Product, error) {
	return us.repo.GetAll()
}

func (us *UseCase) GetProductsByCategorySlug(categorySlug string) (*[]models.Product, error) {
	category, err := us.categoryRepo.GetCategoryBySlug(categorySlug)
	if err != nil {
		return nil, err
	}
	return us.repo.GetProductsByCategory(category)
}
