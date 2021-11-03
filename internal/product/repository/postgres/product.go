package postgres

import (
	"gorm.io/gorm"
	"shop/internal/product"
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
	return repo.db.AutoMigrate(&product.Product{})
}

func (repo *Repository) GetAll() (*[]models.Product, error) {
	dbProduct := new([]product.Product)
	result := repo.db.Find(&dbProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return toModels(dbProduct), nil
}

func (repo *Repository) GetProductsByCategory(category *models.Category) (*[]models.Product, error) {
	dbProduct := new([]product.Product)
	result := repo.db.Where("category_id = ?", category.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return toModels(dbProduct), nil
}

func toModel(product *product.Product) *models.Product {
	return &models.Product{
		Id:    product.Id,
		Title: product.Title,
		Category: models.Category{
			Id:          strconv.Itoa(product.CategoryID),
			Title:       product.Category.Title,
			Slug:        product.Category.Slug,
			Description: product.Category.Description,
			Image: models.Image{
				Url: product.Category.ImageUrl,
			},
		},
		Description: product.Description,
		Image: models.Image{
			Url: product.ImageUrl,
		},
		Price:    product.Price,
		OldPrice: product.OldPrice,
		Amount:   product.Amount,
	}
}

func toModels(products *[]product.Product) *[]models.Product {
	out := make([]models.Product, len(*products))
	for index, instance := range *products {
		out[index] = *toModel(&instance)
	}
	return &out
}
