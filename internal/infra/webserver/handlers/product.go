package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	product, err := entity.NewProduct(createProductInput.Name, createProductInput.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	err = h.ProductRepository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductRepository.FindAll(pageInt, limitInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: errors.New("empty product id").Error()})
		return
	}
	product, err := h.ProductRepository.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
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
		json.NewEncoder(w).Encode(Error{Message: errors.New("empty product id").Error()})
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	product.ID, err = entityPkg.ParseID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	_, err = h.ProductRepository.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	err = h.ProductRepository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: errors.New("empty product id").Error()})
		return
	}
	_, err := h.ProductRepository.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	err = h.ProductRepository.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
