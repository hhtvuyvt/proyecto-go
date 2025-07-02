package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestRepo(t *testing.T) models.BookRepository {
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

	return models.BookRepository{DB: db}
}

func TestGetBooks(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("esperaba status 200, obtuve %d", res.Code)
	}

	if !strings.Contains(res.Body.String(), "1984") {
		t.Error("respuesta no contiene libro esperado")
	}
}

func TestCreateBook(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	body := `{"title":"Clean Code","author":"Robert C. Martin","isbn":"1234567890","image":"img.jpg"}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("esperaba status 200, obtuve %d", res.Code)
	}

	if !strings.Contains(res.Body.String(), "Clean Code") {
		t.Error("respuesta no contiene libro agregado")
	}
}

func TestCreateBookInvalidJSON(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	body := `{"title": "Fallo`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperaba 400 por JSON inválido, obtuve %d", res.Code)
	}
}

func TestCreateBookMissingFields(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	body := `{"title": "", "author": "", "isbn": "", "image": ""}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperaba 400 por campos vacíos, obtuve %d", res.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	req := httptest.NewRequest(http.MethodDelete, "/api/books/1", nil)
	res := httptest.NewRecorder()

	h.Book(res, req)

	if res.Code != http.StatusNoContent {
		t.Fatalf("esperaba 204 No Content, obtuve %d", res.Code)
	}

	// Verificamos que el libro ya no exista
	books, err := repo.GetAll()
	if err != nil {
		t.Fatal("error consultando libros:", err)
	}
	if len(books) != 0 {
		t.Error("esperaba que no haya libros después del delete")
	}
}

func TestDeleteInvalidID(t *testing.T) {
	repo := setupTestRepo(t)
	h := BookHandler{Repo: repo}

	req := httptest.NewRequest(http.MethodDelete, "/api/books/xyz", nil)
	res := httptest.NewRecorder()

	h.Book(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperaba 400 por ID inválido, obtuve %d", res.Code)
	}
}
