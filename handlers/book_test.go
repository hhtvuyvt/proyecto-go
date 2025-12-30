package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type mockRepo struct{}

func (m mockRepo) Create(b *models.Book) error {
	b.ID = 1
	return nil
}
func (m mockRepo) Delete(id int) error { return nil }
func (m mockRepo) Update(id int, b *models.Book) error {
	return nil
}
func (m mockRepo) GetPaginated(s string, l, o int) ([]models.Book, error) {
	return []models.Book{}, nil
}
func (m mockRepo) Count(s string) (int, error) { return 0, nil }

func TestGetBooksPublic(t *testing.T) {
	h := BookHandler{Repo: mockRepo{}}

	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	rec := httptest.NewRecorder()

	h.Books(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d", rec.Code)
	}
}

func TestCreateBookWithoutToken(t *testing.T) {
	h := BookHandler{Repo: mockRepo{}}

	book := models.Book{
		Title:  "Test",
		Author: "Autor",
		ISBN:   "123",
		Image:  "img",
	}

	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/api/books/1", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	h.Book(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401 sin token, obtuve %d", rec.Code)
	}
}
