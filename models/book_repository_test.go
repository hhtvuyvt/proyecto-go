package models

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestRepo(t *testing.T) BookRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("error abriendo base de datos:", err)
	}

	const ddl = `
	CREATE TABLE books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT NOT NULL,
		image TEXT NOT NULL
	);`
	if _, err := db.Exec(ddl); err != nil {
		t.Fatal("error creando tabla:", err)
	}

	return BookRepository{DB: db}
}

func TestCreateAndGetPaginated(t *testing.T) {
	repo := setupTestRepo(t)

	books := []Book{
		{Title: "Go", Author: "Gopher", ISBN: "123", Image: "img1.jpg"},
		{Title: "Rust", Author: "Rustacean", ISBN: "456", Image: "img2.jpg"},
		{Title: "Go in Depth", Author: "Expert", ISBN: "789", Image: "img3.jpg"},
	}

	for _, b := range books {
		if err := repo.Create(&b); err != nil {
			t.Fatal("error al crear libro:", err)
		}
	}

	// Buscar todos (sin filtro)
	found, err := repo.GetPaginated("", 10, 0)
	if err != nil {
		t.Fatal("error en GetPaginated:", err)
	}
	if len(found) != 3 {
		t.Errorf("esperaba 3 libros, obtuve %d", len(found))
	}

	// Búsqueda filtrada
	goBooks, err := repo.GetPaginated("Go", 10, 0)
	if err != nil {
		t.Fatal("error buscando libros con 'Go':", err)
	}
	if len(goBooks) != 2 {
		t.Errorf("esperaba 2 libros con 'Go', obtuve %d", len(goBooks))
	}
}

func TestDeleteBook(t *testing.T) {
	repo := setupTestRepo(t)

	book := Book{Title: "Eliminable", Author: "Autor", ISBN: "999", Image: "x.jpg"}
	if err := repo.Create(&book); err != nil {
		t.Fatal("error creando libro:", err)
	}

	if err := repo.Delete(int(book.ID)); err != nil {
		t.Fatal("error eliminando libro:", err)
	}

	books, err := repo.GetPaginated("", 10, 0)
	if err != nil {
		t.Fatal("error listando libros:", err)
	}

	for _, b := range books {
		if b.ID == book.ID {
			t.Error("el libro eliminado aún aparece en la lista")
		}
	}
}
