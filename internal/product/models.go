package product

import (
	"gorm.io/gorm"
	categoryModels "shop/internal/category"
)

type Product struct {
	gorm.Model
	Title       string
	CategoryID  int
	Category    categoryModels.Category
	Description string
	ImageUrl    string
	Price       float64
	OldPrice    float64
	Amount      int
}
