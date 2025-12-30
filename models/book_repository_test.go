package models

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestRepo(t *testing.T) BookRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ddl := `
	CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		author TEXT,
		isbn TEXT,
		image TEXT
	);`
	db.Exec(ddl)

	return BookRepository{DB: db}
}

func TestCreateAndUpdateBook(t *testing.T) {
	repo := setupTestRepo(t)

	book := Book{
		Title:  "Original",
		Author: "Autor",
		ISBN:   "123",
		Image:  "img",
	}

	if err := repo.Create(&book); err != nil {
		t.Fatal(err)
	}

	book.Title = "Editado"
	if err := repo.Update(int(book.ID), &book); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteBook(t *testing.T) {
	repo := setupTestRepo(t)

	book := Book{
		Title:  "Eliminar",
		Author: "Autor",
		ISBN:   "123",
		Image:  "img",
	}
	repo.Create(&book)

	if err := repo.Delete(int(book.ID)); err != nil {
		t.Fatal(err)
	}
}
