package db

import (
	"database/sql"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type SQLiteChapterRepository struct {
	DB *sql.DB
}

func (r *SQLiteChapterRepository) GetByBookID(bookID int64) ([]models.Chapter, error) {
	query := "SELECT id, book_id, chapter_number, title, content, created_at FROM chapters WHERE book_id = ? ORDER BY chapter_number ASC"
	rows, err := r.DB.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []models.Chapter
	for rows.Next() {
		var c models.Chapter
		if err := rows.Scan(&c.ID, &c.BookID, &c.ChapterNumber, &c.Title, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		chapters = append(chapters, c)
	}
	return chapters, nil
}

func (r *SQLiteChapterRepository) GetByID(chapterID int64) (*models.Chapter, error) {
	query := "SELECT id, book_id, chapter_number, title, content, created_at FROM chapters WHERE id = ?"
	row := r.DB.QueryRow(query, chapterID)

	var c models.Chapter
	if err := row.Scan(&c.ID, &c.BookID, &c.ChapterNumber, &c.Title, &c.Content, &c.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (r *SQLiteChapterRepository) Create(c *models.Chapter) error {
	query := "INSERT INTO chapters (book_id, chapter_number, title, content) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, c.BookID, c.ChapterNumber, c.Title, c.Content)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}
