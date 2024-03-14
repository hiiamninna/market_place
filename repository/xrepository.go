package repository

import (
	"database/sql"
)

type Repository struct {
	USER         User
	PRODUCT      Product
	BANK_ACCOUNT BankAccount
	PAYMENT      Payment
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		USER:         NewUserRepository(db),
		PRODUCT:      NewProductRepository(db),
		BANK_ACCOUNT: NewBankAccountRepository(db),
		PAYMENT:      NewPaymentRepository(db),
	}
}
