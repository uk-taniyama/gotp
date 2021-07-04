package services

import (
	"database/sql"
	"time"
)

var (
	db *sql.DB = nil
)

func UseDB(f func(*sql.DB) error) error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	err = f(db)
	defer db.Close()
	return err
}
