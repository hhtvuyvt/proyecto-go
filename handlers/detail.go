package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// DetailHandler maneja las peticiones HTTP relacionadas con Details (antes Chapters).
type DetailHandler struct {
	Repo models.DetailRepositoryInterface
}

// List devuelve todos los detalles de un item específico (GET /api/items/{item_id}/details)
func (h *DetailHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	itemIDStr := r.PathValue("item_id")
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		http.Error(w, "item_id inválido", http.StatusBadRequest)
		return
	}

	details, err := h.Repo.GetByItemID(itemID)
	if err != nil {
		http.Error(w, "error obteniendo detalles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

// Get obtiene un detalle específico por su ID (GET /api/items/{item_id}/details/{id})
func (h *DetailHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	detail, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "detalle no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(detail)
}

// Create crea un nuevo detalle asociado a un item (POST /api/items/{item_id}/details)
func (h *DetailHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	itemIDStr := r.PathValue("item_id")
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		http.Error(w, "item_id inválido", http.StatusBadRequest)
		return
	}

	var detail models.Detail
	if err := json.NewDecoder(r.Body).Decode(&detail); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	// Asegurar que el detalle pertenezca al item de la URL
	detail.ItemID = itemID

	if err := h.Repo.Create(&detail); err != nil {
		http.Error(w, "error creando detalle", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(detail)
}
