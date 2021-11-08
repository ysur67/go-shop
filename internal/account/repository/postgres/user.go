package postgres

import (
	"gorm.io/gorm"
	repoModels "shop/internal/account"
	"shop/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) AutoMigrate() error {
	return repo.db.AutoMigrate(&repoModels.User{})
}

func (repo *Repository) CreateUser(user models.User) error {
	return repo.db.Create(toUser(&user)).Error
}

func (repo *Repository) GetUser(user models.User) (*models.User, error) {
	dbUser := toUser(&user)
	result := repo.db.Where("username = ?", user.Username).Where("password = ?", user.Password).Find(dbUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func toUser(model *models.User) *repoModels.User {
	return &repoModels.User{
		Username: model.Username,
		Password: model.Password,
	}
}
