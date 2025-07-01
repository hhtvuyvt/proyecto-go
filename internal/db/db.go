package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Open devuelve una conexión y garantiza que exista la tabla books.
func Open() *sql.DB {
	// Crea la carpeta data/ si no existe.
	if err := os.MkdirAll("data", 0o755); err != nil {
		log.Fatal(err)
	}

	dsn := filepath.Join("data", "books.db")
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}

	const ddl = `
	CREATE TABLE IF NOT EXISTS books(
		id     INTEGER PRIMARY KEY AUTOINCREMENT,
		title  TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn   TEXT NOT NULL UNIQUE,
		image  TEXT NOT NULL
	);`
	if _, err = db.Exec(ddl); err != nil {
		log.Fatal(err)
	}
	return db
}
