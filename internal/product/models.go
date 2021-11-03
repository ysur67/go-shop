package product

import categoryModels "shop/internal/category"

type Product struct {
	Id          string
	Title       string
	CategoryID  int
	Category    categoryModels.Category
	Description string
	ImageUrl    string
	Price       float64
	OldPrice    float64
	Amount      int
}
