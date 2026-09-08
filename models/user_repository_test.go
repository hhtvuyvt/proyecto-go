package models

import (
	"database/sql"
	"errors"
	"testing"

	// Usamos el driver limpio de SQLite (asegúrate de que coincida con el de tu proyecto, ej: modernc.org/sqlite o mattn/go-sqlite3)
	_ "modernc.org/sqlite"
)

// setupTestDB es una función auxiliar que prepara la base de datos en la memoria RAM
// y crea la tabla de usuarios idéntica a la de producción para poder testear.
func setupUserTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// Abrimos la conexión en memoria
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Error al abrir base de datos en memoria: %v", err)
	}

	// Creamos la tabla 'users' necesaria para los métodos del repositorio
	query := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(query); err != nil {
		if err := db.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
		t.Fatalf("Error al crear la tabla de pruebas: %v", err)
	}

	return db
}

// TestUserRepository_Create_And_Get verifica que se pueda insertar un usuario
// correctamente y luego recuperarlo tanto por ID como por Username.
func TestUserRepository_Create_And_Get(t *testing.T) {
	db := setupUserTestDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}()

	repo := &UserRepository{DB: db}

	// 1. Preparar el usuario de prueba
	mockUser := &User{
		Username:     "camilo_writer",
		PasswordHash: "$2a$10$encryptedpasswordhashhere",
	}

	// 2. Probar el método Create
	err := repo.Create(mockUser)
	if err != nil {
		t.Fatalf("No se esperaba un error al crear el usuario: %v", err)
	}

	// Verificar que SQLite le asignó un ID autoincremental (debe ser mayor a 0)
	if mockUser.ID == 0 {
		t.Error("Se esperaba que el ID del usuario fuera asignado y diferente de cero")
	}

	// 3. Probar el método GetByUsername
	foundByUsername, err := repo.GetByUsername("camilo_writer")
	if err != nil {
		t.Fatalf("No se esperaba error al buscar por username: %v", err)
	}

	if foundByUsername.ID != mockUser.ID || foundByUsername.Username != mockUser.Username {
		t.Errorf("El usuario encontrado por username (%+v) no coincide con el original (%+v)", foundByUsername, mockUser)
	}

	// 4. Probar el método GetByID
	foundByID, err := repo.GetByID(mockUser.ID)
	if err != nil {
		t.Fatalf("No se esperaba error al buscar por ID: %v", err)
	}

	if foundByID.Username != mockUser.Username || foundByID.PasswordHash != mockUser.PasswordHash {
		t.Errorf("El usuario encontrado por ID (%+v) no coincide con el original (%+v)", foundByID, mockUser)
	}
}

// TestUserRepository_Get_NotFound verifica el comportamiento del repositorio
// cuando se busca un usuario o un ID que no existen en la base de datos.
func TestUserRepository_Get_NotFound(t *testing.T) {
	db := setupUserTestDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("error cerrando la base de datos: %v", err)
		}
	}()

	repo := &UserRepository{DB: db}

	// Probar GetByUsername con un usuario inexistente
	_, err := repo.GetByUsername("usuario_fantasma")
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Se esperaba el error sql.ErrNoRows, pero se obtuvo: %v", err)
	}

	// Probar GetByID con un ID inexistente
	_, err = repo.GetByID(999)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Se esperaba el error sql.ErrNoRows, pero se obtuvo: %v", err)
	}
}
