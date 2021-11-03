package category

import (
	"shop/models"
)

type Repository interface {
	AutoMigrate() error
	GetAll() (*[]models.Category, error)
	GetCategoryById(id string) (*models.Category, error)
}
