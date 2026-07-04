package db

import (
	"database/sql"
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// EnsureAdminUser garantiza
// la existencia del administrador
// inicial.
//
// Si ya existe,
// no realiza ninguna acción.
func EnsureAdminUser(
	db *sql.DB,
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

	var count int

	err :=
		db.QueryRow(
			`
SELECT COUNT(*)
FROM users
WHERE username = ?
`,
			username,
		).Scan(
			&count,
		)

	if err != nil {

		return err

	}

	if count > 0 {

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

	_, err =
		db.Exec(
			`
INSERT INTO users(

	username,

	password_hash

)

VALUES(

	?,

	?

)
`,
			username,

			string(hash),
		)

	return err

}
