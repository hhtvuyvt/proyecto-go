package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestHandler() BookHandler {
	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec(`CREATE TABLE books(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT NOT NULL,
		image TEXT NOT NULL
	);`)
	repo := models.BookRepository{DB: db}
	return BookHandler{Repo: repo}
}

func TestCreateAndListBooks(t *testing.T) {
	h := setupTestHandler()

	// Crear libro
	body := `{"title":"Go","author":"Gopher","isbn":"123","image":"cover.jpg"}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	w := httptest.NewRecorder()
	h.Books(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200, obtuve %d", w.Code)
	}

	// Listar libros
	req = httptest.NewRequest(http.MethodGet, "/api/books", nil)
	w = httptest.NewRecorder()
	h.Books(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200 al listar, obtuve %d", w.Code)
	}

	var books []models.Book
	json.NewDecoder(w.Body).Decode(&books)
	if len(books) != 1 || books[0].Title != "Go" {
		t.Error("libro no encontrado o mal insertado")
	}
}

func TestPaginationAndSearch(t *testing.T) {
	h := setupTestHandler()

	// Insertar libros
	for _, title := range []string{"Go", "Rust", "Python", "Go in Depth"} {
		b := models.Book{Title: title, Author: "Author", ISBN: "123", Image: "img.jpg"}
		h.Repo.Create(&b)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/books?search=Go&page=1&limit=2", nil)
	w := httptest.NewRecorder()
	h.Books(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200 en búsqueda, obtuve %d", w.Code)
	}

	body, _ := io.ReadAll(w.Body)
	if !bytes.Contains(body, []byte("Go")) {
		t.Error("no se encontraron resultados para búsqueda")
	}
}
