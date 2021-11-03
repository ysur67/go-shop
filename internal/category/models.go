package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title       string
	Slug        string
	Description string
	ImageUrl    string
}
