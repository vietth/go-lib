package repository

import (
	"gorm.io/gorm"
	"github.com/vietth/go-lib/domain"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository (db *gorm.DB) domain.ProductRepository {
	return &productRepository{db}
}

func  (pr *productRepository) Fetch() (result [] domain.Product, err error) {
	var product []domain.Product
	err = pr.db.Order(" id desc ").Find(&product).Error
	return product, err
}

func (pr *productRepository) Create(p domain.Product) (result domain.Product, err error) {
	 err = pr.db.Create(&p).Error
    return p, err
}