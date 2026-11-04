package models

import (
	"database/sql"
)

// ItemRepositoryInterface define el contrato del repositorio.
type ItemRepositoryInterface interface {
	GetAll() ([]Item, error)
	GetByID(id int64) (Item, error)
	Create(item *Item) error
	Update(item *Item) error
	Delete(id int64) error
}

// ItemRepository implementa ItemRepositoryInterface usando SQL.
type ItemRepository struct {
	DB *sql.DB
}

// Verificación en tiempo de compilación.
var _ ItemRepositoryInterface = (*ItemRepository)(nil)

// GetAll devuelve todos los items.
func (r *ItemRepository) GetAll() ([]Item, error) {
	// CAMBIO: 'books' por 'items', y nombres de columnas actualizados
	rows, err := r.DB.Query("SELECT id, name, category, code, image FROM items")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Name, &i.Category, &i.Code, &i.Image); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// GetByID obtiene un item por ID.
func (r *ItemRepository) GetByID(id int64) (Item, error) {
	var i Item
	err := r.DB.QueryRow(
		"SELECT id, name, category, code, image FROM items WHERE id = ?",
		id,
	).Scan(&i.ID, &i.Name, &i.Category, &i.Code, &i.Image)

	return i, err
}

// Create inserta un nuevo item.
func (r *ItemRepository) Create(item *Item) error {
	result, err := r.DB.Exec(
		"INSERT INTO items (name, category, code, image) VALUES (?, ?, ?, ?)",
		item.Name, item.Category, item.Code, item.Image,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.ID = id
	return nil
}

// Update actualiza un item existente.
func (r *ItemRepository) Update(item *Item) error {
	_, err := r.DB.Exec(
		"UPDATE items SET name=?, category=?, code=?, image=? WHERE id=?",
		item.Name, item.Category, item.Code, item.Image, item.ID,
	)
	return err
}

// Delete elimina un item por ID.
func (r *ItemRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}
