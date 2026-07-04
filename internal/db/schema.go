package db

import "database/sql"

// CreateSchema crea todas las tablas
// necesarias para la aplicación.
func CreateSchema(
	db *sql.DB,
) error {

	const books = `
CREATE TABLE IF NOT EXISTS books(

	id INTEGER PRIMARY KEY AUTOINCREMENT,

	title TEXT NOT NULL,

	author TEXT NOT NULL,

	isbn TEXT NOT NULL,

	image TEXT NOT NULL

);`

	if _, err :=
		db.Exec(
			books,
		); err != nil {

		return err

	}

	const users = `
CREATE TABLE IF NOT EXISTS users(

	id INTEGER PRIMARY KEY AUTOINCREMENT,

	username TEXT NOT NULL UNIQUE,

	password_hash TEXT NOT NULL,

	created_at DATETIME NOT NULL
	DEFAULT CURRENT_TIMESTAMP

);`

	if _, err :=
		db.Exec(
			users,
		); err != nil {

		return err

	}

	return nil

}
