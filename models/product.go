package models

type Product struct {
	Id          string
	Title       string
	Category    Category
	Description string
	Image       Image
	Price       float64
	OldPrice    float64
	Amount      int
}
