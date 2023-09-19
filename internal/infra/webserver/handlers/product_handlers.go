package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zHenriqueGN/EasyProduct/internal/dto"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/repository"
	entityPkg "github.com/zHenriqueGN/EasyProduct/pkg/entity"
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

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductRepository.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductRepository.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductRepository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
