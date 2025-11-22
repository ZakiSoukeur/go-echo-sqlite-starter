package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	pool *sql.DB
	err  error
)

func InitDB(path string) error {
	pool, err = sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return nil
}

func GetDB() *sql.DB {
	return pool
}
