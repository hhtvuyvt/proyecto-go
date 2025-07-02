package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type BookHandler struct {
	Repo models.BookRepository
}

// GET /api/books  → lista
// POST /api/books → crea
func (h BookHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		books, err := h.Repo.GetAll()
		if err != nil {
			http.Error(w, "error consultando libros", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "json inválido", http.StatusBadRequest)
			return
		}

		if b.Title == "" || b.Author == "" || b.ISBN == "" || b.Image == "" {
			http.Error(w, "faltan campos obligatorios", http.StatusBadRequest)
			return
		}

		if err := h.Repo.Create(&b); err != nil {
			http.Error(w, "error insertando libro", http.StatusInternalServerError)
			return
		}

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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "error eliminando libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
