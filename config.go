package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}
