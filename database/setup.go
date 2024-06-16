package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func SetupDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	createProductTable := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);`

	_, err = DB.Exec(createProductTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}
}
