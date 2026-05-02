package models

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("error DB: %v", err)
	}

	ddl := `
	CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		isbn TEXT,
		image TEXT
	);`

	if _, err := db.Exec(ddl); err != nil {
		t.Fatalf("error creando tabla: %v", err)
	}

	return db
}

func TestCreateBook(t *testing.T) {
	db := setupTestDB(t)
	repo := BookRepository{DB: db}

	book := Book{Title: "Test"}

	if err := repo.Create(&book); err != nil {
		t.Fatalf("error creando libro: %v", err)
	}

	if book.ID == 0 {
		t.Fatal("ID no asignado")
	}
}