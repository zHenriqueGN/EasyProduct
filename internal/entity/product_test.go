package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 9.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 9.99, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 9.99)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -1)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProduct_Validate(t *testing.T) {
	product, err := NewProduct("Product 1", 9.99)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validade())
}
