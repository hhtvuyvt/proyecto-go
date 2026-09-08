package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Open abre la base de datos, crea el esquema y garantiza la existencia del usuario administrador.
func Open(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	// Comprueba que la conexión realmente funciona.
	if err := database.Ping(); err != nil {
		_ = database.Close()
		return nil, err
	}

	if err := CreateSchema(database); err != nil {
		_ = database.Close()
		return nil, err
	}

	if err := EnsureAdminUser(database); err != nil {
		_ = database.Close()
		return nil, err
	}

	return database, nil
}
