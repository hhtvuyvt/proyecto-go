package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// BookHandler maneja las peticiones relacionadas con libros
type BookHandler struct {
	Repo models.BookRepositoryInterface
}

// Books maneja GET (listar) y POST (crear)
func (h BookHandler) Books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		books, err := h.Repo.GetAll()
		if err != nil {
			http.Error(w, "error obteniendo libros", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(books); err != nil {
			http.Error(w, "error al listar libros", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "datos inválidos", http.StatusBadRequest)
			return
		}

		if err := h.Repo.Create(&b); err != nil {
			http.Error(w, "error creando libro", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(b); err != nil {
			http.Error(w, "error al responder", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

// Book maneja GET, PUT y DELETE por ID
func (h BookHandler) Book(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		book, err := h.Repo.GetByID(id)
		if err != nil {
			http.Error(w, "libro no encontrado", http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(book); err != nil {
			http.Error(w, "error al obtener libro", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		var b models.Book
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "datos inválidos", http.StatusBadRequest)
			return
		}
		b.ID = id

		if err := h.Repo.Update(&b); err != nil {
			http.Error(w, "error actualizando libro", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		if err := h.Repo.Delete(id); err != nil {
			http.Error(w, "error eliminando libro", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}