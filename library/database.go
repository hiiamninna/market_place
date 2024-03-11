package library

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

func NewDatabaseConnection(dbCfg Database) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn(dbCfg))
	if err != nil {
		return db, fmt.Errorf("open con : %w", err)
	}

	err = db.Ping()
	if err != nil {
		return db, fmt.Errorf("db ping : %w", err)
	}

	return db, nil
}

func dsn(dbCfg Database) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
}
