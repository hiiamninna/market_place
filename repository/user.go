package repository

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	return User{
		db: db,
	}
}

func (c *User) Register(name, username, password string) (int, string, interface{}, error) {

	sql :=
		`INSERT INTO 
			users (name, username, password, created_at, updated_at) 
		VALUES 
			(?, ?, ?, current_timestamp, current_timestamp) ; `
	_, err := c.db.Query(sql, name, username, password)

	if err != nil {
		return http.StatusInternalServerError, "failed register", nil, fmt.Errorf("insert : %w", err)
	}

	return http.StatusOK, "success register", nil, nil
}
