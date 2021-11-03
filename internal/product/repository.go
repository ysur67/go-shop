package product

import (
	"errors"
	"shop/models"
)

var DoesNotExist = errors.New("product not found")

type Repository interface {
	AutoMigrate() error
	GetAll() (*[]models.Product, error)
	GetProduct(id string) (*models.Product, error)
	GetProductsByCategory(category *models.Category) (*[]models.Product, error)
}
