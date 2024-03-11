package repository

import (
	"database/sql"
)

type Repository struct {
	USER User
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		USER: NewUserRepository(db),
	}
}
