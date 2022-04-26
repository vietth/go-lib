package usecase

import (
	"github.com/vietth/go-lib/domain"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
}

func NewProductUsecase(pr domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{pr}
}

func (pu *productUsecase) GetAll() (productList [] domain.Product, err error) {
	productList, err = pu.productRepo.Fetch()
   return
}

func (pu *productUsecase) New(p domain.Product) (product domain.Product, err error) {
   product, err = pu.productRepo.Create(p)
   return
}