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

	sql := fmt.Sprintf(
		`INSERT INTO 
			users (name, username, password, created_at, updated_at) 
		VALUES 
			('%s', '%s', '%s', current_timestamp, current_timestamp) ; `, name, username, password)
	_, err := c.db.Query(sql)

	if err != nil {
		return http.StatusInternalServerError, "failed register", nil, fmt.Errorf("insert : %w", err)
	}

	return http.StatusOK, "success register", nil, nil
}
