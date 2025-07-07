// Paquete handlers gestiona las solicitudes HTTP relacionadas con libros.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
	"github.com/hhtvuyvt/proyecto-go/utils"
)

// BookHandler representa un controlador HTTP para recursos de libros.
type BookHandler struct {
	Repo models.BookRepository
}

// Books maneja solicitudes a /api/books.
// Soporta:
//   - GET: lista libros con búsqueda y paginación
//   - POST: crea un nuevo libro validando sus campos
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
			http.Error(w, "error insertando libro", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)

	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

// Book maneja DELETE /api/books/{id}.
// Permite eliminar libros por ID.
func (h BookHandler) Book(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, _ := strconv.Atoi(idStr)

	res, err := h.Repo.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		http.Error(w, "error eliminando libro", http.StatusInternalServerError)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		http.Error(w, "libro no encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
