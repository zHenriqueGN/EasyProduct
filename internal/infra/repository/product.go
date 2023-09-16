package repository

import (
	"database/sql"
	"errors"

	"github.com/zHenriqueGN/EasyProduct/internal/entity"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) *ProductRepository {
	return &ProductRepository{DB}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	stmt, err := p.DB.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) FindByID(id string) (*entity.Product, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, price, created_at FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		var product entity.Product
		err = row.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &product, nil
	}
	return nil, ErrProductNotFound
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var (
		products     []entity.Product
		findAllQuery string
		queryParams  []any
	)
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		queryParams = append(queryParams, limit, (page-1)*limit, sort)
		findAllQuery = "SELECT id, name, price, created_at FROM products LIMIT ? OFFSET ? ORDER BY created_at ?"
	} else {
		findAllQuery = "SELECT id, name, price, created_at FROM products ORDER BY created_at ?"
	}
	stmt, err := p.DB.Prepare(findAllQuery)
	if err != nil {
		return products, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(queryParams...)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	stmt, err := p.DB.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID.String())
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	stmt, err := p.DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
