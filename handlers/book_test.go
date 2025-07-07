// book_test contiene pruebas unitarias para los handlers HTTP relacionados con libros.
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

// setupTestHandler crea un entorno de prueba con una base de datos SQLite en memoria.
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

// TestCreateAndListBooks prueba que se puedan crear y luego listar libros correctamente.
func TestCreateAndListBooks(t *testing.T) {
	h := setupTestHandler()

	body := `{"title":"Go","author":"Gopher","isbn":"123","image":"cover.jpg"}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	w := httptest.NewRecorder()
	h.Books(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperado 200, obtuve %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/books", nil)
	w = httptest.NewRecorder()
	h.Books(w, req)

	var books []models.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Fatal("respuesta no es JSON válido")
	}
	if len(books) != 1 || books[0].Title != "Go" {
		t.Error("el libro no se insertó o recuperó correctamente")
	}
}

// TestPaginationAndSearch verifica que la búsqueda y paginación funcionen correctamente.
func TestPaginationAndSearch(t *testing.T) {
	h := setupTestHandler()

	titles := []string{"Go", "Rust", "Python", "Go in Depth"}
	for _, title := range titles {
		b := models.Book{Title: title, Author: "Autor", ISBN: "000", Image: "img.jpg"}
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
		t.Error("no se encontraron resultados para búsqueda 'Go'")
	}
}
