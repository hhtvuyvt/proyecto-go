package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

// mockRepo implementa la interfaz para pruebas sin DB real.
type mockRepo struct{}

func (m mockRepo) GetAll() ([]models.Book, error) {
	return []models.Book{
		{ID: 1, Title: "Mock Book", Author: "Autor", ISBN: "123"},
	}, nil
}

func (m mockRepo) GetByID(id int) (models.Book, error) {
	return models.Book{ID: int64(id), Title: "Mock Book"}, nil
}

func (m mockRepo) Create(b *models.Book) error {
	b.ID = 1
	return nil
}

func (m mockRepo) Update(b *models.Book) error {
	return nil
}

func (m mockRepo) Delete(id int) error {
	return nil
}

// TestBooksGET verifica que GET /api/books responde correctamente.
func TestBooksGET(t *testing.T) {
	handler := BookHandler{Repo: mockRepo{}}

	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	w := httptest.NewRecorder()

	handler.Books(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("esperado 200, obtenido %d", w.Code)
	}
}