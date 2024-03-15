package repository

import (
	"database/sql"
	"fmt"
	"market_place/collections"
	"strings"

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

	sql := `SELECT TEXT(id), name, price, image_url, stock, condition, is_purchaseable, user_id, tags FROM products WHERE id = $1 and deleted_at is null;`
	err := c.db.QueryRow(sql, id).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.Stock, &p.Condition, &p.IsPurchaseable, &p.UserID, pq.Array(&p.Tags))
	if err != nil {
		return p, fmt.Errorf("get by id : %w", err)
	}

	return p, nil
}

func (c *Product) List(input collections.ProductPageInput) ([]collections.ProductList, error) {
	products := []collections.ProductList{}
	var filter, order string
	var values []interface{}

	valuesCount := 1

	isSort := false

	sql := `SELECT products.id, products.name, products.price, products.image_url, 
	products.stock, products.condition, products.is_purchaseable, products.tags, sum(payment.quantity)
	FROM products
	INNER JOIN payments on products.id = payments.products_id
	WHERE products.deleted_at is null [filter] [order] [limit];`

	if input.UserOnly {
		filter += " AND products.user_id = $" + fmt.Sprint(valuesCount)
		values = append(values, input.UserID)
		valuesCount += 1
	}

	if len(input.Tags) > 0 {
		filter += " AND products.tags = ANY($" + fmt.Sprint(valuesCount) + ")"
		values = append(values, input.Tags)
		valuesCount += 1
	}

	if input.Condition == "new" || input.Condition == "second" {
		filter += " AND products.condition = $" + fmt.Sprint(valuesCount)
		values = append(values, input.Condition)
		valuesCount += 1
	}

	if !input.ShowEmptyStock {
		filter += " AND products.stock > 0"
	}

	if input.MaxPrice > 0 {
		filter += " AND products.price < $" + fmt.Sprint(valuesCount)
		values = append(values, input.MaxPrice)
		valuesCount += 1
	}

	if input.MinPrice > 0 {
		filter += " AND products.price > $" + fmt.Sprint(valuesCount)
		values = append(values, input.MinPrice)
		valuesCount += 1
	}

	if input.Search != "" {
		filter += " AND products.name ~* $" + fmt.Sprint(valuesCount)
		values = append(values, input.Search)
		valuesCount += 1
	}

	switch input.SortBy {
	case "price":
		order += " ORDER BY products.price "
		isSort = true
	case "date":
		order += " ORDER BY products.created_at "
		isSort = true
	}

	if isSort && input.OrderBy != "" {
		order += " ORDER BY " + input.OrderBy
	}

	sql = strings.ReplaceAll(sql, "[filter]", filter)
	sql = strings.ReplaceAll(sql, "[order]", order)
	sql = strings.ReplaceAll(sql, "[limit]", fmt.Sprintf("limit %d offset %d", input.Limit, input.Offset))

	rows, err := c.db.Query(sql, values)
	if err != nil {
		return products, fmt.Errorf("select list : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := collections.ProductList{}

		err := rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.ImageURL, &p.Stock, &p.Condition, &p.IsPurchaseable, pq.Array(&p.Tags), &p.PurchaseCount)
		if err != nil {
			return products, fmt.Errorf("rows scan : %w", err)
		}

		products = append(products, p)
	}

	return products, nil
}

func (c *Product) CountList(input collections.ProductPageInput) (int, error) {
	var filter string
	var values []interface{}

	valuesCount := 1

	sql := `SELECT count(*)
	FROM products
	INNER JOIN payments on products.id = payments.products_id
	WHERE products.deleted_at is null [filter] `

	if input.UserOnly {
		filter += " AND products.user_id = $" + fmt.Sprint(valuesCount)
		values = append(values, input.UserID)
		valuesCount += 1
	}

	if len(input.Tags) > 0 {
		filter += " AND products.tags = ANY($" + fmt.Sprint(valuesCount) + ")"
		values = append(values, input.Tags)
		valuesCount += 1
	}

	if input.Condition == "new" || input.Condition == "second" {
		filter += " AND products.condition = $" + fmt.Sprint(valuesCount)
		values = append(values, input.Condition)
		valuesCount += 1
	}

	if !input.ShowEmptyStock {
		filter += " AND products.stock > 0"
	}

	if input.MaxPrice > 0 {
		filter += " AND products.price < $" + fmt.Sprint(valuesCount)
		values = append(values, input.MaxPrice)
		valuesCount += 1
	}

	if input.MinPrice > 0 {
		filter += " AND products.price > $" + fmt.Sprint(valuesCount)
		values = append(values, input.MinPrice)
		valuesCount += 1
	}

	if input.Search != "" {
		filter += " AND products.name ~* $" + fmt.Sprint(valuesCount)
		values = append(values, input.Search)
		valuesCount += 1
	}

	sql = strings.ReplaceAll(sql, "[filter]", filter)

	rows, err := c.db.Query(sql, values)
	if err != nil {
		return 0, fmt.Errorf("select list : %w", err)
	}
	defer rows.Close()

	var totalRow int
	for rows.Next() {
		err := rows.Scan(&totalRow)
		if err != nil {
			return 0, fmt.Errorf("rows scan : %w", err)
		}
	}

	return totalRow, nil
}

func (c *Product) UpdateStock(id string, stock int) error {

	sql := `UPDATE products SET stock = $1 WHERE id = $2;`
	_, err := c.db.Exec(sql, stock, id)
	if err != nil {
		return fmt.Errorf("update stock : %w", err)
	}

	return nil
}
