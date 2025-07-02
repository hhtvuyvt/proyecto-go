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

	schema := `
	CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT NOT NULL,
		image TEXT NOT NULL
	);
	INSERT INTO books (title, author, isbn, image) VALUES
	('1984', 'George Orwell', '978-0451524935', 'img.jpg');`

	if _, err := db.Exec(schema); err != nil {
		t.Fatal(err)
	}

	return BookRepository{DB: db}
}

func TestRepoGetAll(t *testing.T) {
	repo := setupTestRepo(t)
	books, err := repo.GetAll()
	if err != nil {
		t.Fatalf("error en GetAll: %v", err)
	}
	if len(books) != 1 || books[0].Title != "1984" {
		t.Errorf("esperaba 1 libro '1984', obtuve: %+v", books)
	}
}

func TestRepoCreate(t *testing.T) {
	repo := setupTestRepo(t)
	book := &Book{
		Title:  "Clean Code",
		Author: "Robert C. Martin",
		ISBN:   "1234567890",
		Image:  "img.jpg",
	}
	err := repo.Create(book)
	if err != nil {
		t.Fatalf("error al crear libro: %v", err)
	}
	if book.ID == 0 {
		t.Error("ID del libro no fue asignado")
	}

	books, _ := repo.GetAll()
	if len(books) != 2 {
		t.Errorf("esperaba 2 libros después de crear, obtuve %d", len(books))
	}
}

func TestRepoDelete(t *testing.T) {
	repo := setupTestRepo(t)

	err := repo.Delete(1)
	if err != nil {
		t.Fatalf("error eliminando libro: %v", err)
	}

	books, _ := repo.GetAll()
	if len(books) != 0 {
		t.Error("libro no fue eliminado correctamente")
	}
}

func TestRepoDeleteNonExistent(t *testing.T) {
	repo := setupTestRepo(t)

	err := repo.Delete(999) // no existe
	if err != nil {
		t.Errorf("esperaba que eliminar libro inexistente no fallara, pero falló: %v", err)
	}
}
