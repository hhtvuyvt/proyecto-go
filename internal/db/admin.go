package db

import (
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// EnsureAdminUser garantiza la existencia del usuario administrador.
// Si las variables de entorno no están configuradas, no hace nada
// (ideal para producción SaaS donde el registro es público).
func EnsureAdminUser(database *sql.DB) error {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	// Si no están configurados, omitimos la creación silenciosamente.
	// Esto evita que la app falle en producción si no se desea un admin por defecto.
	if username == "" || password == "" {
		log.Println("⚠️ ADMIN_USERNAME o ADMIN_PASSWORD no configurados. Omitiendo creación de admin.")
		return nil
	}

	repo := models.UserRepository{DB: database}

	// Si ya existe, no hay nada que hacer.
	_, err := repo.GetByUsername(username)
	if err == nil {
		return nil // El usuario ya existe, todo bien.
	}

	// Generar hash de la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := models.User{
		Username:     username,
		PasswordHash: string(hash),
	}

	err = repo.Create(&admin)
	if err != nil {
		return err
	}

	log.Println("✅ Usuario administrador creado exitosamente.")
	return nil
}
