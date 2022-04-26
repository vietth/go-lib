package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/vietth/go-lib/domain"
)

type ProductHandle struct {
	productUsecase domain.ProductUsecase
}

type ProductsResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Data    []domain.Product
}

type ProductResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Data    domain.Product
}

func ListProductHandle(w http.ResponseWriter, r *http.Request, pu domain.ProductUsecase) {
	handler := &ProductHandle{pu}
	handler.FetchProduct(w, r)
}

func CreateProductHandle(w http.ResponseWriter, r *http.Request, pu domain.ProductUsecase) {
	handler := &ProductHandle{pu}
	if r.Method == "POST" {
		handler.NewProduct(w, r)
	}
}

func (ph *ProductHandle) FetchProduct(w http.ResponseWriter, r *http.Request) {
	product, _ := ph.productUsecase.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var responseData = ProductsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    product,
	}
	json.NewEncoder(w).Encode(responseData)

}

func (ph *ProductHandle) NewProduct(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.Atoi(r.FormValue("price"))
	w.Header().Set("Content-Type", "application/json")
	var p = domain.Product{
		Name:  r.FormValue("name"),
		Code:  r.FormValue("code"),
		Price: int64(price),
	}
	product, err := ph.productUsecase.New(p)

	if err != nil {
		fmt.Printf("Create product error : %v \n", err)
	}

	w.WriteHeader(http.StatusOK)
	var responseData = ProductResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    product,
	}
	json.NewEncoder(w).Encode(responseData)

}
