package models

import "database/sql"

type BookRepository struct {
	DB *sql.DB
}

func (r BookRepository) Create(b *Book) error {
	res, err := r.DB.Exec(
		"INSERT INTO books(title, author, isbn, image) VALUES (?, ?, ?, ?)",
		b.Title, b.Author, b.ISBN, b.Image)
	if err != nil {
		return err
	}
	b.ID, _ = res.LastInsertId()
	return nil
}

func (r BookRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

func (r BookRepository) GetPaginated(search string, limit, offset int) ([]Book, error) {
	query := "SELECT id, title, author, isbn, image FROM books"
	args := []interface{}{}

	if search != "" {
		query += " WHERE title LIKE ?"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
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

func (r BookRepository) Count(search string) (int, error) {
	query := "SELECT COUNT(*) FROM books"
	args := []interface{}{}

	if search != "" {
		query += " WHERE title LIKE ?"
		args = append(args, "%"+search+"%")
	}

	var count int
	err := r.DB.QueryRow(query, args...).Scan(&count)
	return count, err
}
