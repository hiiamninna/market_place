package repository

import (
	"database/sql"
	"fmt"
	"market_place/collections"
)

type BankAccount struct {
	db *sql.DB
}

func NewBankAccountRepository(db *sql.DB) BankAccount {
	return BankAccount{
		db: db,
	}
}

func (c *BankAccount) Create(input collections.BankAccountInput) error {

	sql :=
		`INSERT INTO
				bank_accounts (name, account_name, account_number, created_at, updated_at)
		VALUES
			 ($1, $2, $3, current_timestamp, current_timestamp);`

	_, err := c.db.Exec(sql, input.BankName, input.BankAccountName, input.BankAccountNumber)
	if err != nil {
		return fmt.Errorf("insert : %w", err)
	}

	return nil
}

func (c *BankAccount) Update(input collections.BankAccountInput) error {

	sql :=
		`UPDATE bank_accounts 
		SET name = $1, account_name = $2, account_number = $3, updated_at = current_timestamp
		WHERE id = $4 AND deleted_at is null;`
	_, err := c.db.Exec(sql, input.BankName, input.BankAccountName, input.BankAccountNumber, input.ID)
	if err != nil {
		return fmt.Errorf("update : %w", err)
	}

	return nil
}

func (c *BankAccount) Delete(id string) error {
	sql := `UPDATE bank_accounts SET deleted_at = current_timestamp WHERE id = $1;`
	_, err := c.db.Exec(sql, id)
	if err != nil {
		return fmt.Errorf("delete : %w", err)
	}

	return nil
}

func (c *BankAccount) GetByID(id string) (collections.BankAccount, error) {

	bankAccount := collections.BankAccount{}

	sql := `SELECT TEXT(id), name, account_name, account_number FROM bank_accounts WHERE id = $1 and deleted_at is null;`
	err := c.db.QueryRow(sql, id).Scan(&bankAccount.ID, &bankAccount.BankName, &bankAccount.BankAccountName, &bankAccount.BankAccountNumber)
	if err != nil {
		return bankAccount, fmt.Errorf("get by id : %w", err)
	}

	return bankAccount, nil
}

func (c *BankAccount) List() ([]collections.BankAccount, error) {
	bankAccounts := []collections.BankAccount{}

	sql := `SELECT id, name, account_name, account_number FROM bank_accounts WHERE deleted_at is null;`
	rows, err := c.db.Query(sql)
	if err != nil {
		return bankAccounts, fmt.Errorf("select list : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		b := collections.BankAccount{}

		err := rows.Scan(&b.ID, &b.BankName, &b.BankAccountName, &b.BankAccountNumber)
		if err != nil {
			return bankAccounts, fmt.Errorf("rows scan : %w", err)
		}

		bankAccounts = append(bankAccounts, b)
	}

	return bankAccounts, nil
}
