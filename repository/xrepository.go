package repository

import (
	"database/sql"
)

type Repository struct {
	USER    User
	PRODUCT Product
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		USER:    NewUserRepository(db),
		PRODUCT: NewProductRepository(db),
	}
}
