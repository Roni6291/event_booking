package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbName string, maxOpenConns int, maxIdleConns int) *sql.DB {
	DB, err := sql.Open("sqlite3", dbName)

	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxOpenConns(maxOpenConns)
	DB.SetMaxIdleConns(maxIdleConns)

	createTables(DB)
	return DB
}

func createTables(db *sql.DB) {

	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := db.Exec(createUsersTableQuery)
	if err != nil {
		panic("Couldn't create users table")
	}
	createEventTableQuery := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			time DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	_, err = db.Exec(createEventTableQuery)
	if err != nil {
		panic("Couldn't create events table")
	}
}
