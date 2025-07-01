package repository

import (
	"database/sql"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// BookRepository encapsula todas las consultas SQL de Book.
type BookRepository struct {
	DB *sql.DB
}

func (r BookRepository) List() ([]models.Book, error) {
	rows, err := r.DB.Query(`SELECT id, title, author, isbn, image FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Image); err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, nil
}

func (r BookRepository) Create(b *models.Book) error {
	res, err := r.DB.Exec(
		`INSERT INTO books(title, author, isbn, image) VALUES(?, ?, ?, ?)`,
		b.Title, b.Author, b.ISBN, b.Image,
	)
	if err != nil {
		return err
	}
	b.ID, err = res.LastInsertId()
	return err
}

func (r BookRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM books WHERE id = ?`, id)
	return err
}
