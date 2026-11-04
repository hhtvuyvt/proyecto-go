package db

import "database/sql"

// CreateSchema crea todas las tablas necesarias para la aplicación.
// Actualizado para modelo SaaS (Items y Details en lugar de Books y Chapters).
func CreateSchema(db *sql.DB) error {
	const schema = `
	-- Tabla de usuarios (se mantiene igual, es la base del SaaS)
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

	-- Tabla de Items (reemplaza a 'books'. Ej: Productos en tienda, Habitaciones en hotel)
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,        -- Antes: title
		category TEXT NOT NULL,    -- Antes: author
		code TEXT NOT NULL,        -- Antes: isbn (SKU, código de habitación, etc.)
		image TEXT NOT NULL
	);

	-- Tabla de Detalles (reemplaza a 'chapters'. Ej: Variantes, Reservas, Especificaciones)
	CREATE TABLE IF NOT EXISTS details (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id INTEGER NOT NULL,  -- Antes: book_id
		name TEXT NOT NULL,        -- Antes: title (del capítulo)
		content TEXT NOT NULL,     -- Antes: content
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
	);
	`

	if _, err := db.Exec(schema); err != nil {
		return err
	}

	return nil
}
