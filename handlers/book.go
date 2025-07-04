package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hhtvuyvt/proyecto-go/models"
	"github.com/hhtvuyvt/proyecto-go/utils"
)

type BookHandler struct {
	Repo models.BookRepository
}

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

func (h BookHandler) Book(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/api/books/")
	id, _ := strconv.Atoi(idStr)

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "error eliminando libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
