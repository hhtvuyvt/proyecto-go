package models

import (
	"database/sql"
	"time"
)

// Detail representa un detalle asociado a un Item (ej: variante de producto, reserva de habitación).
// Antes era "Chapter" asociado a "Book".
type Detail struct {
	ID        int64     `json:"id"`
	ItemID    int64     `json:"item_id"` // Relación con el Item (antes book_id)
	Name      string    `json:"name"`    // Nombre del detalle (antes title)
	Content   string    `json:"content"` // Descripción o contenido (antes content/body)
	CreatedAt time.Time `json:"created_at"`
}

// DetailRepositoryInterface define el contrato para acceder a detalles.
type DetailRepositoryInterface interface {
	GetByItemID(itemID int64) ([]Detail, error)
	GetByID(id int64) (Detail, error)
	Create(detail *Detail) error
	Update(detail *Detail) error
	Delete(id int64) error
}

// DetailRepository implementa DetailRepositoryInterface usando SQLite.
type DetailRepository struct {
	DB *sql.DB
}

// Comprobación en tiempo de compilación.
var _ DetailRepositoryInterface = (*DetailRepository)(nil)

// GetByItemID obtiene todos los detalles de un Item específico.
func (r *DetailRepository) GetByItemID(itemID int64) ([]Detail, error) {
	rows, err := r.DB.Query(`
		SELECT id, item_id, name, content, created_at
		FROM details
		WHERE item_id = ?
	`, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []Detail
	for rows.Next() {
		var d Detail
		if err := rows.Scan(&d.ID, &d.ItemID, &d.Name, &d.Content, &d.CreatedAt); err != nil {
			return nil, err
		}
		details = append(details, d)
	}
	return details, rows.Err()
}

// GetByID obtiene un detalle específico por su ID.
func (r *DetailRepository) GetByID(id int64) (Detail, error) {
	var d Detail
	err := r.DB.QueryRow(`
		SELECT id, item_id, name, content, created_at
		FROM details
		WHERE id = ?
	`, id).Scan(&d.ID, &d.ItemID, &d.Name, &d.Content, &d.CreatedAt)
	return d, err
}

// Create inserta un nuevo detalle.
func (r *DetailRepository) Create(detail *Detail) error {
	result, err := r.DB.Exec(`
		INSERT INTO details (item_id, name, content)
		VALUES (?, ?, ?)
	`, detail.ItemID, detail.Name, detail.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	detail.ID = id
	return nil
}

// Update actualiza un detalle existente.
func (r *DetailRepository) Update(detail *Detail) error {
	_, err := r.DB.Exec(`
		UPDATE details
		SET name = ?, content = ?
		WHERE id = ?
	`, detail.Name, detail.Content, detail.ID)
	return err
}

// Delete elimina un detalle por su ID.
func (r *DetailRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM details WHERE id = ?`, id)
	return err
}
