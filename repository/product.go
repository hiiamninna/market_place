package repository

import (
	"database/sql"
	"fmt"
	"market_place/collections"

	"github.com/lib/pq"
)

type Product struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) Product {
	return Product{
		db: db,
	}
}

func (c *Product) Create(input collections.ProductInput) error {

	sql :=
		`INSERT INTO 
			products (name, price, image_url, stock, condition, is_purchaseable, user_id, tags, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, current_timestamp, current_timestamp);`
	_, err := c.db.Exec(sql, input.Name, input.Price, input.ImageUrl, input.Stock, input.Condition, input.IsPurchaseable, input.UserID, pq.Array(input.Tags))

	if err != nil {
		return fmt.Errorf("insert : %w", err)
	}

	return nil
}

func (c *Product) Update(input collections.ProductInput) error {
	sql :=
		`UPDATE products 
		SET name = $1, price = $2, condition = $3, is_purchaseable = $4, tags = $5, updated_at = current_timestamp
		WHERE id = $6 AND deleted_at is null;`
	_, err := c.db.Exec(sql, input.Name, input.Price, input.Condition, input.IsPurchaseable, pq.Array(input.Tags), input.ID)
	if err != nil {
		return fmt.Errorf("update : %w", err)
	}

	return nil
}

func (c *Product) Delete(id string) error {
	sql := `UPDATE products SET deleted_at = current_timestamp WHERE id = $1;`
	_, err := c.db.Exec(sql, id)
	if err != nil {
		return fmt.Errorf("delete : %w", err)
	}

	return nil
}

func (c *Product) GetOwnByID(id, userID string) (collections.Product, error) {

	p := collections.Product{}

	sql := `SELECT id, name, price, image_url, stock, condition, is_purchaseable FROM products WHERE id = $1 AND user_id = $2 AND deleted_at is null;`
	err := c.db.QueryRow(sql, id, userID).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.Stock, &p.Condition, &p.IsPurchaseable)
	if err != nil {
		return p, fmt.Errorf("get own by id : %w", err)
	}

	return p, nil
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
