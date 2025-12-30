package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
	"github.com/hhtvuyvt/proyecto-go/utils"
)

// BookHandler maneja las rutas HTTP relacionadas con libros.
type BookHandler struct {
	Repo BookRepository
}

// BookRepository define las operaciones necesarias para manejar libros.
// Permite usar mocks en tests y distintas implementaciones.
type BookRepository interface {
	Create(*models.Book) error
	Update(int, *models.Book) error
	Delete(int) error
	GetPaginated(string, int, int) ([]models.Book, error)
	Count(string) (int, error)
}

// Books maneja /api/books
// - GET: listar libros
// - POST: crear libro
func (h BookHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		search := r.URL.Query().Get("search")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		books, err := h.Repo.GetPaginated(search, limit, offset)
		if err != nil {
			http.Error(w, "error obteniendo libros", http.StatusInternalServerError)
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

		utils.SanitizeBook(&b)

		if b.Title == "" || b.Author == "" || b.ISBN == "" || b.Image == "" {
			http.Error(w, "faltan campos obligatorios", http.StatusBadRequest)
			return
		}

		if err := h.Repo.Create(&b); err != nil {
			http.Error(w, "error creando libro", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

// Book maneja /api/books/{id}
// - DELETE: eliminar libro
// - PUT: editar libro
func (h BookHandler) Book(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodDelete:
		if err := h.Repo.Delete(id); err != nil {
			http.Error(w, "error eliminando libro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	case http.MethodPut:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "json inválido", http.StatusBadRequest)
			return
		}

		utils.SanitizeBook(&b)

		if b.Title == "" || b.Author == "" || b.ISBN == "" || b.Image == "" {
			http.Error(w, "faltan campos obligatorios", http.StatusBadRequest)
			return
		}

		if err := h.Repo.Update(id, &b); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "libro no encontrado", http.StatusNotFound)
				return
			}
			http.Error(w, "error actualizando libro", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}
