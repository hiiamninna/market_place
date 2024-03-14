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

func (c *User) Create(input collections.InputUserRegister) (int, error) {

	var id int
	sql :=
		`INSERT INTO 
			users (name, username, password, created_at, updated_at) 
		VALUES 
			($1, $2, $3, current_timestamp, current_timestamp)
		RETURNING id;`
	rows, err := c.db.Query(sql, input.Name, input.Username, input.Password)

	for rows.Next() {
		rows.Scan(&id)
	}
	defer rows.Close()

	if err != nil {
		return id, fmt.Errorf("insert : %w", err)
	}

	return id, nil
}

func (c *User) GetByID(id string) (collections.User, error) {

	user := collections.User{}

	sql := `SELECT TEXT(id), name, username, password FROM users WHERE id = $1 and deleted_at is null;`
	err := c.db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, fmt.Errorf("get by id : %w", err)
	}

	return user, nil
}

func (c *User) GetByUsername(username string) (collections.User, error) {

	user := collections.User{}

	sql := `SELECT TEXT(id), name, username, password FROM users WHERE UPPER(username) = UPPER($1) and deleted_at is null;`
	err := c.db.QueryRow(sql, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		return user, fmt.Errorf("get by username : %w", err)
	}

	return user, nil
}
