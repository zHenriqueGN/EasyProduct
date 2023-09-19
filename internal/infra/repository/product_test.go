package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zHenriqueGN/EasyProduct/internal/entity"
	"github.com/zHenriqueGN/EasyProduct/internal/infra/database"
)

var (
	testDBDriver string
	testDBConn   string
)

func init() {
	testDBDriver = "mysql"
	testDBConn = "root:secret@tcp(localhost:3306)/easyproduct?charset=utf8&parseTime=True&loc=Local"
}

func TruncateProductsTable(DB *sql.DB) error {
	stmt, err := DB.Prepare("TRUNCATE TABLE products")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func TestProductRepository_Create(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewProductRepository(DB)
	product, err := entity.NewProduct("Product 1", 199.99)
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(product)
	assert.Nil(t, err)
}

func TestProductRepository_FindByID(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewProductRepository(DB)
	product, err := entity.NewProduct("Product 1", 199.99)
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(product)
	if err != nil {
		t.Fatal(err)
	}
	productFound, err := repository.FindByID(product.ID.String())
	assert.NotNil(t, productFound)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, productFound.CreatedAt)
}

func TestProductRepository_Update(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewProductRepository(DB)
	product, err := entity.NewProduct("Product 1", 199.99)
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(product)
	if err != nil {
		t.Fatal(err)
	}
	product.Name = "Product 2"
	product.Price = 211.99
	err = repository.Update(product)
	assert.Nil(t, err)
	productFound, err := repository.FindByID(product.ID.String())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProductRepository_Delete(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewProductRepository(DB)
	product, err := entity.NewProduct("Product 1", 199.99)
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Create(product)
	if err != nil {
		t.Fatal(err)
	}
	err = repository.Delete(product.ID.String())
	assert.Nil(t, err)
	productFound, err := repository.FindByID(product.ID.String())
	assert.Nil(t, productFound)
	assert.Equal(t, ErrProductNotFound, err)
}

func TestProductRepository_FindAll(t *testing.T) {
	DB := database.ConnectToDatabase(testDBDriver, testDBConn)
	defer DB.Close()
	repository := NewProductRepository(DB)
	err := TruncateProductsTable(DB)
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i <= 15; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), (rand.Float64() * 100))
		if err != nil {
			t.Fatal(err)
		}
		err = repository.Create(product)
		if err != nil {
			t.Fatal(err)
		}
	}
	products, err := repository.FindAll(1, 10)
	assert.NotEmpty(t, products)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	products, err = repository.FindAll(2, 10)
	assert.NotEmpty(t, products)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(products))
	products, err = repository.FindAll(0, 0)
	assert.NotEmpty(t, products)
	assert.Nil(t, err)
	assert.Equal(t, 15, len(products))
}
