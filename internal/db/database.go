package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Open abre la base de datos,
// crea el esquema
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

	// Comprueba que la conexión realmente funciona.
	if err :=
		database.Ping(); err != nil {

		if closeErr :=
			database.Close(); closeErr != nil {

			return nil, closeErr

		}

		return nil, err

	}

	if err :=
		CreateSchema(
			database,
		); err != nil {

		if closeErr :=
			database.Close(); closeErr != nil {

			return nil, closeErr

		}

		return nil, err

	}

	if err :=
		EnsureAdminUser(
			database,
		); err != nil {

		if closeErr :=
			database.Close(); closeErr != nil {

			return nil, closeErr

		}

		return nil, err

	}

	return database, nil

}
