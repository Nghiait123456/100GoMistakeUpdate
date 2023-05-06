package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	fmt.Println("Init DB")
}

func NewDB() {
	fmt.Println("newDB")
}

func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
