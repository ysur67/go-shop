package postgres

import (
	"context"
	"gorm.io/gorm"
	categoryModels "shop/internal/category"
	"shop/models"
	"strconv"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetCategoryById(ctx context.Context, id string) (*models.Category, error) {
	dbCategory := new(categoryModels.Category)
	result := repo.db.First(&dbCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toModel(dbCategory), nil
}

func toModel(category *categoryModels.Category) *models.Category {
	return &models.Category{
		Id:          strconv.Itoa(int(category.Model.ID)),
		Title:       category.Title,
		Slug:        category.Slug,
		Description: category.Description,
		Image: models.Image{
			Url: category.ImageUrl,
		},
	}
}
