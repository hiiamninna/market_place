package repository

import (
	"database/sql"
	"fmt"
	"market_place/collections"
)

type Product struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) Product {
	return Product{
		db: db,
	}
}

func (c *Product) Create(input collections.ProductInput) (interface{}, error) {

	sql :=
		`INSERT INTO 
			products (id, name, price, image_url, stock, condition, is_purchaseable, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, current_timestamp, current_timestamp);`
	result, err := c.db.Exec(sql, input.ID, input.Name, input.Price, input.ImageUrl, input.Stock, input.Condition, input.IsPurchaseable)

	if err != nil {
		return nil, fmt.Errorf("insert : %w", err)
	}

	return result, nil
}

func (c *Product) Update(input collections.ProductInput) (interface{}, error) {
	sql :=
		`UPDATE products 
		SET name = $1, price = $2, condition = $3, is_purchaseable = $4, updated_at = current_timestamp
		WHERE id = $5 AND deleted_at is null;`
	result, err := c.db.Exec(sql, input.Name, input.Price, input.Condition, input.IsPurchaseable, input.ID)
	if err != nil {
		return nil, fmt.Errorf("update : %w", err)
	}

	return result, nil
}

func (c *Product) Delete(id string) error {
	sql := `UPDATE products SET deleted_at = current_timestamp WHERE id = $1;`
	_, err := c.db.Exec(sql, id)
	if err != nil {
		return fmt.Errorf("delete : %w", err)
	}

	return nil
}

func (c *Product) GetByID(id string) (collections.Product, error) {

	p := collections.Product{}

	sql := `SELECT id, name, price, image_url, stock, condition, is_purchaseable FROM products WHERE id = $1 and deleted_at is null;`
	err := c.db.QueryRow(sql, id).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.Stock, &p.Condition, &p.IsPurchaseable)
	if err != nil {
		return p, fmt.Errorf("get by id : %w", err)
	}

	return p, nil
}

func (c *Product) List() ([]collections.Product, error) {
	products := []collections.Product{}

	sql := `SELECT id, name, price, image_url, stock, condition, is_purchaseable FROM products WHERE deleted_at is null;`
	rows, err := c.db.Query(sql)
	if err != nil {
		return products, fmt.Errorf("select list : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := collections.Product{}

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.Stock, &p.Condition, &p.IsPurchaseable)
		if err != nil {
			return products, fmt.Errorf("rows scan : %w", err)
		}

		products = append(products, p)
	}

	return products, nil
}

func (c *Product) UpdateStock(id string, stock int) error {

	sql := `UPDATE products SET stock = $1 WHERE id = $2;`
	_, err := c.db.Exec(sql, stock, id)
	if err != nil {
		return fmt.Errorf("update stock : %w", err)
	}

	return nil
}
