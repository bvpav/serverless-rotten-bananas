package database

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	MovieDto
	AverageRating float64 `json:"averageRating"`
}

type Review struct {
	gorm.Model
	ReviewDto
	Movie Movie `json:"-"`
}
