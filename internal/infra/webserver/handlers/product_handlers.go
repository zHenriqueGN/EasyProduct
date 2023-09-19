package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zHenriqueGN/EasyProduct/internal/dto"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/repository"
)

type ProductHandler struct {
	ProductRepository repository.ProductInterface
}

func NewProductHandler(productRepository repository.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductRepository: productRepository}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProductInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&createProductInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(createProductInput.Name, createProductInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductRepository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
