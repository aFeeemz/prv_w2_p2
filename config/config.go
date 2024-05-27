package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost)/prv_w2_p2")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
