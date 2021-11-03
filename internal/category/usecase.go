package category

import (
	"shop/models"
)

type UseCase interface {
	GetAll() (*[]models.Category, error)
	GetCategory(id string) (*models.Category, error)
}
