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
