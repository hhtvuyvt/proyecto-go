package db

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateSchema(
	t *testing.T,
) {

	database, err :=
		sql.Open(
			"sqlite3",
			":memory:",
		)

	if err != nil {

		t.Fatal(err)

	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)

	if err :=
		CreateSchema(
			database,
		); err != nil {

		t.Fatalf(
			"CreateSchema devolvió error: %v",
			err,
		)

	}

	// =========================
	// Verificar tabla books
	// =========================

	var name string

	err =
		database.QueryRow(

			`SELECT name
			 FROM sqlite_master
			 WHERE type='table'
			 AND name='books'`,
		).Scan(&name)

	if err != nil {

		t.Fatal(
			"no existe la tabla books",
		)

	}

	// =========================
	// Verificar tabla users
	// =========================

	err =
		database.QueryRow(

			`SELECT name
			 FROM sqlite_master
			 WHERE type='table'
			 AND name='users'`,
		).Scan(&name)

	if err != nil {

		t.Fatal(
			"no existe la tabla users",
		)

	}

	// =========================
	// Verificar índice
	// =========================

	err =
		database.QueryRow(

			`SELECT name
			 FROM sqlite_master
			 WHERE type='index'
			 AND name='idx_users_username'`,
		).Scan(&name)

	if err != nil {

		t.Fatal(
			"no existe el índice",
		)

	}

}

func TestCreateSchemaTwice(
	t *testing.T,
) {

	database, err :=
		sql.Open(
			"sqlite3",
			":memory:",
		)

	if err != nil {

		t.Fatal(err)

	}

	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {

		}
	}(database)

	if err :=
		CreateSchema(
			database,
		); err != nil {

		t.Fatal(err)

	}

	// Debe poder ejecutarse nuevamente
	// gracias al IF NOT EXISTS.

	if err :=
		CreateSchema(
			database,
		); err != nil {

		t.Fatalf(
			"la segunda ejecución falló: %v",
			err,
		)

	}

}
