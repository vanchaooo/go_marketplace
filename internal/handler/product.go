package handler

import (
	"log"
	"errors"
	"net/http"
	"strconv"
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

type updateProductRequest struct {
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

func handleServiceErrors(w http.ResponseWriter, err error) {
	switch {
		case
			errors.Is(err, service.ErrInvalidName),
			errors.Is(err, service.ErrInvalidPrice),
			errors.Is(err, service.ErrInvalidStock):
			writeError(w, http.StatusBadRequest, err.Error())
	
		case
			errors.Is(err, service.ErrProductNotFound):
			writeError(w, http.StatusNotFound, err.Error())

		default:
			log.Println("internal error:", err)
			writeError(w, http.StatusInternalServerError, "внутренняя ошибка")
	}
}

func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request createProductRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	product, err := p.products.AddNewProduct(r.Context(), request.Name, request.Price, request.Stock)
	if err != nil{
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusCreated, productToResponse(product))
}

func (p *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	var request updateProductRequest
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	product, err := p.products.UpdateProduct(r.Context(), id, request.Name, request.Price, request.Stock)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, productToResponse(product))
}

func (p *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	if err := p.products.DeleteProduct(r.Context(), id); err != nil {
		handleServiceErrors(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "некорректный ID")
		return
	}

	product, err := p.products.GetProduct(r.Context(), id)
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	_ = writeJSON(w, http.StatusOK, productToResponse(product))
}

func (p *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.products.GetAllProducts(r.Context())
	if err != nil {
		handleServiceErrors(w, err)
		return
	}

	response := make([]productResponse, 0, len(products))
	for _, product := range products {
		response = append(response, productToResponse(product))
	}

	_ = writeJSON(w, http.StatusOK, response)
}