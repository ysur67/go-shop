package usecase

import (
	"shop/internal/category"
)

type UseCase struct {
	repo category.Repository
}

func NewUseCase(repo category.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
