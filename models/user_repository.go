package models

import (
	"database/sql"
	"time"
)

// User representa un usuario del sistema.
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // Nunca debe enviarse al frontend.
	CreatedAt    time.Time `json:"created_at"`
}

// UserRepositoryInterface define el contrato para acceder a usuarios.
type UserRepositoryInterface interface {
	GetByUsername(username string) (User, error)
	GetByID(id int64) (User, error)
	Create(user *User) error
}

// UserRepository implementa UserRepositoryInterface usando SQLite.
type UserRepository struct {
	DB *sql.DB
}

// Comprobación en tiempo de compilación.
var _ UserRepositoryInterface = (*UserRepository)(nil)

// GetByUsername obtiene un usuario usando su nombre.
func (r *UserRepository) GetByUsername(username string) (User, error) {
	var user User
	err := r.DB.QueryRow(`
		SELECT id, username, password_hash, created_at
		FROM users
		WHERE username = ?
	`, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	return user, err
}

// GetByID obtiene un usuario usando su ID.
func (r *UserRepository) GetByID(id int64) (User, error) {
	var user User
	err := r.DB.QueryRow(`
		SELECT id, username, password_hash, created_at
		FROM users
		WHERE id = ?
	`, id).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	return user, err
}

// Create inserta un nuevo usuario en la base de datos y le asigna el ID generado.
func (r *UserRepository) Create(user *User) error {
	result, err := r.DB.Exec(`
		SELECT id, username, password_hash, created_at
		FROM users
		WHERE id = ?
	`, user.Username, user.PasswordHash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}
