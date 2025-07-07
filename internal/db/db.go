// Paquete db se encarga de abrir y preparar la base de datos SQLite.
package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Driver para SQLite
)

// Open abre la base de datos SQLite y crea la tabla books si no existe.
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
