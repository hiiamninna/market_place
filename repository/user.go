package repository

import (
	"database/sql"
	"fmt"
	"market_place/collections"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	return User{
		db: db,
	}
}

func (c *User) Create(input collections.InputUserRegister) (interface{}, error) {

	sql :=
		`INSERT INTO 
			users (id, name, username, password, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, current_timestamp, current_timestamp);`
	result, err := c.db.Exec(sql, input.ID, input.Name, input.Username, input.Password)

	if err != nil {
		return nil, fmt.Errorf("insert : %w", err)
	}

	return result, nil
}

func (c *User) GetByID(id string) (collections.User, error) {

	user := collections.User{}

	sql := `SELECT id, name, username, password FROM users WHERE id = $1 and deleted_at is null;`
	err := c.db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, fmt.Errorf("get by id : %w", err)
	}

	return user, nil
}

func (c *User) GetByUsername(username string) (collections.User, error) {

	user := collections.User{}

	sql := `SELECT id, name, username, password FROM users WHERE UPPER(username) = UPPER($1) and deleted_at is null;`
	err := c.db.QueryRow(sql, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, fmt.Errorf("get by username : %w", err)
	}

	return user, nil
}
