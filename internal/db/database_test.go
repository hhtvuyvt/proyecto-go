package db

import (
	"testing"
)

func TestOpen(
	t *testing.T,
) {

	t.Setenv(
		"ADMIN_USERNAME",
		"admin",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"admin123",
	)

	database, err :=
		Open(
			":memory:",
		)

	if err != nil {

		t.Fatalf(
			"Open devolvió error: %v",
			err,
		)

	}

	if database == nil {

		t.Fatal(
			"Open devolvió una base de datos nil",
		)

	}

	defer func() {

		if err := database.Close(); err != nil {

			t.Fatalf(
				"error cerrando la base de datos: %v",
				err,
			)

		}

	}()

	if err := database.Ping(); err != nil {

		t.Fatalf(
			"la conexión no responde: %v",
			err,
		)

	}

}

func TestOpenWithoutAdminUsername(
	t *testing.T,
) {

	t.Setenv(
		"ADMIN_USERNAME",
		"",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"admin123",
	)

	database, err :=
		Open(
			":memory:",
		)

	if err == nil {

		if database != nil {

			_ = database.Close()

		}

		t.Fatal(
			"se esperaba un error",
		)

	}

}

func TestOpenWithoutAdminPassword(
	t *testing.T,
) {

	t.Setenv(
		"ADMIN_USERNAME",
		"admin",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"",
	)

	database, err :=
		Open(
			":memory:",
		)

	if err == nil {

		if database != nil {

			_ = database.Close()

		}

		t.Fatal(
			"se esperaba un error",
		)

	}

}
