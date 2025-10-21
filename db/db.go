package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "keymatch.db")
	if err != nil {
		panic("Could not conect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createKeymatchTable := `
	CREATE TABLE IF NOT EXISTS keymatch (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		keyword TEXT NOT NULL,
		domain TEXT NOT NULL,
		synonym_id INTEGER,
		UNIQUE(keyword, domain),
		FOREIGN KEY (synonym_id) REFERENCES synonyms(id)
	)
	`

	_, err := DB.Exec(createKeymatchTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createSynonymTable := `
	CREATE TABLE IF NOT EXISTS synonym (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		domain TEXT NOT NULL,
		url TEXT NOT NULL,
		caption TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createSynonymTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
