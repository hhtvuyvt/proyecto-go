package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// ⚙️ Crea una base de datos en memoria para tests
func setupTestDB(t *testing.T) *sql.DB {
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
	return db
}

// ✅ Test GET /api/books
func TestGetBooks(t *testing.T) {
	db := setupTestDB(t)
	h := BookHandler{DB: db}

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

// ✅ Test POST /api/books (libro válido)
func TestCreateBook(t *testing.T) {
	db := setupTestDB(t)
	h := BookHandler{DB: db}

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

// ❌ Test POST con JSON inválido
func TestCreateBookInvalidJSON(t *testing.T) {
	db := setupTestDB(t)
	h := BookHandler{DB: db}

	body := `{"title": "Fallo` // JSON mal formado
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperaba 400 por JSON inválido, obtuve %d", res.Code)
	}
}

// ❌ Test POST con campos vacíos
func TestCreateBookMissingFields(t *testing.T) {
	db := setupTestDB(t)
	h := BookHandler{DB: db}

	body := `{"title": "", "author": "", "isbn": "", "image": ""}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	h.Books(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("esperaba 400 por campos vacíos, obtuve %d", res.Code)
	}
}
