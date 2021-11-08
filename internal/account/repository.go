package account

import "shop/models"

type Repository interface {
	AutoMigrate() error
	CreateUser(user models.User) error
	GetUser(user models.User) (*models.User, error)
}
