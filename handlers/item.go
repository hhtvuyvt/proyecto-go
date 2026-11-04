package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// ItemHandler maneja las peticiones HTTP relacionadas con Items (antes Books).
type ItemHandler struct {
	Repo models.ItemRepositoryInterface
}

// List devuelve todos los items (GET /api/items)
func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	items, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "error obteniendo items", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Create crea un nuevo item (POST /api/items)
func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&item); err != nil {
		http.Error(w, "error creando item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Get obtiene un item por su ID (GET /api/items/{id})
func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id") // Go 1.22+
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	item, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "item no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Update actualiza un item existente (PUT /api/items/{id})
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	item.ID = id // Asegurar que el ID de la URL prevalezca sobre el del JSON

	if err := h.Repo.Update(&item); err != nil {
		http.Error(w, "error actualizando item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete elimina un item por su ID (DELETE /api/items/{id})
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "error eliminando item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
