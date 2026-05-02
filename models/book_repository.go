package models

import (
	"database/sql"
)

// BookRepository maneja el acceso a la base de datos.
type BookRepository struct {
	DB *sql.DB
}

// GetAll devuelve todos los libros.
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

// GetByID obtiene un libro por ID.
func (r BookRepository) GetByID(id int) (Book, error) {
	var b Book

	err := r.DB.QueryRow(
		"SELECT id, title, author, isbn, image FROM books WHERE id = ?",
		id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Image)

	return b, err
}

// Create inserta un nuevo libro.
func (r BookRepository) Create(book *Book) error {
	result, err := r.DB.Exec(
		"INSERT INTO books (title, author, isbn, image) VALUES (?, ?, ?, ?)",
		book.Title, book.Author, book.ISBN, book.Image,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.ID = id
	return nil
}

// Update actualiza un libro existente.
func (r BookRepository) Update(book *Book) error {
	_, err := r.DB.Exec(
		"UPDATE books SET title=?, author=?, isbn=?, image=? WHERE id=?",
		book.Title, book.Author, book.ISBN, book.Image, book.ID,
	)
	return err
}

// Delete elimina un libro por ID.
func (r BookRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}