package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	_ = godotenv.Load() // Lee variables de entorno de .env

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/books.db" // Valor por defecto
	}

	database, err := sql.Open("sqlite3", dbPath)
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
