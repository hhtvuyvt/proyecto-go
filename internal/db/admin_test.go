package db

import (
	"database/sql"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/models"

	_ "github.com/mattn/go-sqlite3"
)

// ===================================
// Utilidades
// ===================================

func createTestDB(
	t *testing.T,
) *sql.DB {

	t.Helper()

	database, err :=
		sql.Open(
			"sqlite3",
			":memory:",
		)

	if err != nil {

		t.Fatal(err)

	}

	if err :=
		CreateSchema(
			database,
		); err != nil {

		t.Fatal(err)

	}

	return database

}

// ===================================
// Tests
// ===================================

func TestEnsureAdminUserCreatesAdmin(
	t *testing.T,
) {

	database :=
		createTestDB(t)

	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}(database)

	t.Setenv(
		"ADMIN_USERNAME",
		"admin",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"admin123",
	)

	err :=
		EnsureAdminUser(
			database,
		)

	if err != nil {

		t.Fatalf(
			"EnsureAdminUser devolvió error: %v",
			err,
		)

	}

	repo :=
		models.UserRepository{
			DB: database,
		}

	user, err :=
		repo.GetByUsername(
			"admin",
		)

	if err != nil {

		t.Fatalf(
			"no se creó el administrador: %v",
			err,
		)

	}

	if user.Username != "admin" {

		t.Fatalf(
			"usuario incorrecto: %s",
			user.Username,
		)

	}

	if user.PasswordHash == "" {

		t.Fatal(
			"el hash quedó vacío",
		)

	}

	if err :=
		bcrypt.CompareHashAndPassword(

			[]byte(
				user.PasswordHash,
			),

			[]byte(
				"admin123",
			),
		); err != nil {

		t.Fatal(
			"la contraseña almacenada no coincide",
		)

	}

}

func TestEnsureAdminUserAlreadyExists(
	t *testing.T,
) {

	database :=
		createTestDB(t)

	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}(database)

	t.Setenv(
		"ADMIN_USERNAME",
		"admin",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"admin123",
	)

	if err :=
		EnsureAdminUser(
			database,
		); err != nil {

		t.Fatal(err)

	}

	if err :=
		EnsureAdminUser(
			database,
		); err != nil {

		t.Fatalf(
			"segunda llamada devolvió error: %v",
			err,
		)

	}

}

func TestEnsureAdminUserWithoutUsername(
	t *testing.T,
) {

	database :=
		createTestDB(t)

	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}(database)

	t.Setenv(
		"ADMIN_USERNAME",
		"",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"admin123",
	)

	err :=
		EnsureAdminUser(
			database,
		)

	if err == nil {

		t.Fatal(
			"se esperaba un error",
		)

	}

}

func TestEnsureAdminUserWithoutPassword(
	t *testing.T,
) {

	database :=
		createTestDB(t)

	defer func(database *sql.DB) {
		if err := database.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}(database)

	t.Setenv(
		"ADMIN_USERNAME",
		"admin",
	)

	t.Setenv(
		"ADMIN_PASSWORD",
		"",
	)

	err :=
		EnsureAdminUser(
			database,
		)

	if err == nil {

		t.Fatal(
			"se esperaba un error",
		)

	}

	var total int

	err = database.QueryRow(
		"SELECT COUNT(*) FROM users",
	).Scan(&total)

	if err != nil {

		t.Fatal(err)

	}

	if total != 0 {

		t.Fatalf(
			"se esperaba 0 usuario, hay %d",
			total,
		)

	}

}
