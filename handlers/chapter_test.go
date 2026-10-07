package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type fakeChapterRepo struct{}

func (fakeChapterRepo) Create(ch *models.Chapter) error {
	return nil
}

func (fakeChapterRepo) GetByBookID(bookID int64) ([]models.Chapter, error) {
	return nil, nil
}

func (fakeChapterRepo) GetByID(chapterID int64) (*models.Chapter, error) {
	return nil, nil
}

func TestCreateChapterHandlerRejectsInvalidBookID(t *testing.T) {
	handler := &ChapterHandler{ChapterRepo: fakeChapterRepo{}}

	req := httptest.NewRequest(http.MethodPost, "/api/books/not-a-number/chapters", strings.NewReader(`{"chapter_number":1,"title":"Test","content":"Contenido"}`))
	req.SetPathValue("book_id", "not-a-number")
	rec := httptest.NewRecorder()

	handler.CreateChapterHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, obtuvo %d", rec.Code)
	}
}
