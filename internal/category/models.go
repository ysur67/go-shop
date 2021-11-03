package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id          int64
	Title       string
	Slug        string
	Description string
	ImageUrl    string
}
