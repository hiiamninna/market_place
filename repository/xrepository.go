package repository

import (
	"database/sql"
)

type Repository struct {
	USER         User
	PRODUCT      Product
	BANK_ACCOUNT BankAccount
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		USER:         NewUserRepository(db),
		PRODUCT:      NewProductRepository(db),
		BANK_ACCOUNT: NewBankAccountRepository(db),
	}
}
