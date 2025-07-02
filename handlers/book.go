package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type BookHandler struct{ DB *sql.DB }

// GET /api/books  → lista
// POST /api/books → crea
func (h BookHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, _ := h.DB.Query("SELECT id,title,author,isbn,image FROM books")
		defer rows.Close()

		var out []models.Book
		for rows.Next() {
			var b models.Book
			rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Image)
			out = append(out, b)
		}
		json.NewEncoder(w).Encode(out)

	case http.MethodPost:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "json inválido", 400)
			return
		}
		res, _ := h.DB.Exec(
			"INSERT INTO books(title,author,isbn,image) VALUES (?,?,?,?)",
			b.Title, b.Author, b.ISBN, b.Image)
		b.ID, _ = res.LastInsertId()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

// DELETE /api/books/{id}
func (h BookHandler) Book(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, _ := strconv.Atoi(idStr)
	h.DB.Exec("DELETE FROM books WHERE id = ?", id)
	w.WriteHeader(http.StatusNoContent)
}
