package models

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("error abriendo DB: %v", err)
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

func TestCreateAndGetBook(t *testing.T) {
	db := setupTestDB(t)
	repo := BookRepository{DB: db}

	book := Book{
		Title:  "Test",
		Author: "Autor",
		ISBN:   "123",
		Image:  "img",
	}

	if err := repo.Create(&book); err != nil {
		t.Fatalf("error creando libro: %v", err)
	}

	got, err := repo.GetByID(book.ID)
	if err != nil {
		t.Fatalf("error obteniendo libro: %v", err)
	}

	if got.Title != book.Title {
		t.Errorf("esperado %s, obtenido %s", book.Title, got.Title)
	}
}