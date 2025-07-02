package models

import (
	"database/sql"
)

type BookRepository struct {
	DB *sql.DB
}

func (r BookRepository) GetAll() ([]Book, error) {
	rows, err := r.DB.Query("SELECT id, title, author, isbn, image FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Image); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r BookRepository) Create(b *Book) error {
	res, err := r.DB.Exec(
		"INSERT INTO books(title, author, isbn, image) VALUES (?, ?, ?, ?)",
		b.Title, b.Author, b.ISBN, b.Image,
	)
	if err != nil {
		return err
	}
	b.ID, err = res.LastInsertId()
	return err
}

func (r BookRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
