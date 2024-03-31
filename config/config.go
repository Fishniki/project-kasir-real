package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConectDB() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3306)/kasir"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err

	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db, nil
}
