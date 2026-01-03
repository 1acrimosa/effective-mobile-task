package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Price  string `json:"price" validate:"required"`
	Vendor string `json:"vendor" validate:"required"`
}
