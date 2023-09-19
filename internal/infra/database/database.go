package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase(driver, connStr string) *sql.DB {
	DB, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
