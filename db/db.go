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
	createEventTableQuery := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			time DATETIME NOT NULL,
			user_id TEXT
		)
	`
	_, err := db.Exec(createEventTableQuery)
	if err != nil {
		panic(err.Error())
	}
}
