package handler

import (
	"errors"
	"net/http"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/service"
)

type ProductHandler struct {
	products *service.ProductService
}

func NewProductHandler(products *service.ProductService) *ProductHandler {
	return &ProductHandler{products: products}
}

type createProductRequest struct {
	Name string `json:"name"`
	Price int64 `json:"price"`
	Stock int64 `json:"stock"`
}

type productResponse struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Price int64 `json:"price"`
	Stock int64 `json:"stock"`
}

func productToResponse(product model.Product) productResponse {
	return productResponse{
		ID: product.ID,
		Name: product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request createProductRequest
	if err := decodeJSON(r, request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	product, err := p.products.AddNewProduct(r.Context(), request.ID, request.Name, request.Price, request.Stock)
	if err != nil{
		if errors.Is(err, service.ErrInvalidName) || errors.Is(err, service.ErrInvalidPrice) || errors.Is(err, service.ErrInvalidStock) {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "внутреняя ошибка")
		return
	}

	_ = writeJSON(w, http.StatusCreated, productToResponse(product))
}

