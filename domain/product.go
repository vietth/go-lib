package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	Name 	string 	`json:"name" validate:"required"`
	Code    string  `json:"code" validate:"required"`
	Price   int64   `json:"price" validate:"required"`
	gorm.Model
}

type ProductUsecase interface {
	GetAll() (res []Product, err error)
	New(p Product) (res Product, err error)
}

type ProductRepository interface {
	Fetch() (res []Product, err error)
	Create(p Product) (res Product, err error)
}
