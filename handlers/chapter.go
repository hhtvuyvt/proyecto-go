package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type ChapterHandler struct {
	ChapterRepo models.ChapterRepositoryInterface
}

// CreateChapterHandler procesa la creación de un nuevo capítulo
func (h *ChapterHandler) CreateChapterHandler(w http.ResponseWriter, r *http.Request) {
	var c models.Chapter
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Datos de entrada inválidos", http.StatusBadRequest)
		return
	}

	bookIDStr := r.PathValue("book_id")
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de libro inválido", http.StatusBadRequest)
		return
	}

	c.BookID = bookID

	if err := h.ChapterRepo.Create(&c); err != nil {
		http.Error(w, "Error al guardar el capítulo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(c)
}

// GetChapterHandler maneja la petición para obtener un capítulo específico
func (h *ChapterHandler) GetChapterHandler(w http.ResponseWriter, r *http.Request) {
	chapterIDStr := r.PathValue("id")
	chapterID, err := strconv.ParseInt(chapterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de capítulo inválido", http.StatusBadRequest)
		return
	}

	chapter, err := h.ChapterRepo.GetByID(chapterID)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	if chapter == nil {
		http.Error(w, "Capítulo no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(chapter)
}

// ListChaptersHandler maneja la petición para listar todos los capítulos de un libro
func (h *ChapterHandler) ListChaptersHandler(w http.ResponseWriter, r *http.Request) {
	bookIDStr := r.PathValue("book_id")
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de libro inválido", http.StatusBadRequest)
		return
	}

	chapters, err := h.ChapterRepo.GetByBookID(bookID)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(chapters)
}
