package db

import (
	"database/sql"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// EnsureAdminUser garantiza
// la existencia del usuario administrador.
//
// Si ya existe,
// no realiza ninguna acción.
func EnsureAdminUser(
	database *sql.DB,
) error {

	username :=
		os.Getenv(
			"ADMIN_USERNAME",
		)

	password :=
		os.Getenv(
			"ADMIN_PASSWORD",
		)

	if username == "" {

		return errors.New(
			"ADMIN_USERNAME no configurado",
		)

	}

	if password == "" {

		return errors.New(
			"ADMIN_PASSWORD no configurado",
		)

	}

	repo :=
		models.UserRepository{
			DB: database,
		}

	// Si ya existe,
	// no hay nada que hacer.
	_, err :=
		repo.GetByUsername(
			username,
		)

	if err == nil {

		return nil

	}

	hash, err :=
		bcrypt.GenerateFromPassword(

			[]byte(
				password,
			),

			bcrypt.DefaultCost,
		)

	if err != nil {

		return err

	}

	admin :=
		models.User{

			Username: username,

			PasswordHash: string(hash),
		}

	return repo.Create(
		&admin,
	)
}
