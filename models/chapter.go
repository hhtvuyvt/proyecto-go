package models

import "time"

type Chapter struct {
	ID            int64     `json:"id"`
	BookID        int64     `json:"book_id"`
	ChapterNumber int       `json:"chapter_number"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
}

type ChapterRepositoryInterface interface {
	GetByBookID(bookID int64) ([]Chapter, error)
	GetByID(chapterID int64) (*Chapter, error)
	Create(chapter *Chapter) error
}
