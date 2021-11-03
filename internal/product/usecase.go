package product

import "shop/models"

type UseCase interface {
	GetAll() (*[]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	GetProductsByCategorySlug(categorySlug string) (*[]models.Product, error)
}
