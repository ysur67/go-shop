package postgres

import (
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

func (repo *Repository) AutoMigrate() error {
	return repo.db.AutoMigrate(&categoryModels.Category{})
}

func (repo *Repository) GetAll() (*[]models.Category, error) {
	dbCategory := new([]categoryModels.Category)
	result := repo.db.Find(dbCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	return toModels(dbCategory), nil
}

func (repo *Repository) GetCategoryById(id string) (*models.Category, error) {
	dbCategory := new(categoryModels.Category)
	result := repo.db.First(&dbCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toModel(dbCategory), nil
}

func (repo *Repository) GetCategoryBySlug(slug string) (*models.Category, error) {
	dbCategory := new(categoryModels.Category)
	result := repo.db.Where("slug = ?", slug).Find(&dbCategory)
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

func toModels(categories *[]categoryModels.Category) *[]models.Category {
	out := make([]models.Category, len(*categories))
	for index, cat := range *categories {
		out[index] = *toModel(&cat)
	}
	return &out
}
