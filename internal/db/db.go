package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // driver
)

// Open crea (o abre) `books.db` y garantiza la tabla books.
func Open() *sql.DB {
	database, err := sql.Open("sqlite3", "./data/books.db")
	if err != nil {
		log.Fatal(err)
	}

	const ddl = `CREATE TABLE IF NOT EXISTS books(
		id     INTEGER PRIMARY KEY AUTOINCREMENT,
		title  TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn   TEXT NOT NULL,
		image  TEXT NOT NULL
	);`
	if _, err = database.Exec(ddl); err != nil {
		log.Fatal(err)
	}
	return database
}
