package models

import "database/sql"

// / BookRepository administra el acceso a los libros en la base de datos.
type BookRepository struct {
	DB *sql.DB
}

// / Create agrega un nuevo libro a la base de datos.
// / El campo ID del libro se actualiza con el valor generado automáticamente.
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

// / Delete elimina un libro de la base de datos según su ID.
// / Si no existe, la operación no afecta a ningún registro.
func (r BookRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}

// / GetPaginated devuelve una lista de libros con soporte para búsqueda y paginación.
// / - search: filtro por título (LIKE)
// / - limit: cantidad máxima de resultados
// / - offset: desde qué posición comenzar
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

// / Count retorna la cantidad total de libros que cumplen con un criterio de búsqueda.
// / Se utiliza para paginación (conocer el total de páginas disponibles).
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
