package repository

import (
	"context"
	"database/sql"
	"fmt"
	"market_place/collections"
)

type Payment struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) Payment {
	return Payment{
		db: db,
	}
}

func (c Payment) Create(input collections.PaymentInput) error {

	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction : %w", err)
	}

	sql := `INSERT INTO 
				payments (user_id, bank_account_id, product_id, quantity, image_url, total_payment, created_at, updated_at)
			VALUES
				($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp);`
	_, err = tx.ExecContext(ctx, sql, input.UserID, input.BankAccountID, input.ProductID, input.Quantity, input.PaymentProof, input.TotalPayment)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert : %w", err)
	}

	sql = `UPDATE products SET stock = stock - $1 WHERE id = $2;`
	_, err = tx.ExecContext(ctx, sql, input.Quantity, input.ProductID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update stock : %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit : %w", err)
	}

	return nil
}
