package product

import "shop/models"

type Repository interface {
	AutoMigrate() error
	GetAll() (*[]models.Product, error)
	GetProduct(id string) (*models.Product, error)
	GetProductsByCategory(category *models.Category) (*[]models.Product, error)
}
