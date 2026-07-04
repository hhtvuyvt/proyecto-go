package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Open abre la base de datos SQLite,
// crea el esquema necesario
// y garantiza la existencia
// del usuario administrador.
func Open(
	path string,
) (*sql.DB, error) {

	database, err :=
		sql.Open(
			"sqlite3",
			path,
		)

	if err != nil {

		return nil, err

	}

	if err :=
		CreateSchema(
			database,
		); err != nil {

		return nil, err

	}

	if err :=
		EnsureAdminUser(
			database,
		); err != nil {

		return nil, err

	}

	return database, nil

}
