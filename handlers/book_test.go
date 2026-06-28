package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type mockBookRepo struct {
	books []models.Book
	err   error
}

func (m *mockBookRepo) GetAll() ([]models.Book, error) {

	if m.err != nil {
		return nil, m.err
	}

	return m.books, nil
}

func (m *mockBookRepo) GetByID(id int) (models.Book, error) {

	if m.err != nil {
		return models.Book{}, m.err
	}

	for _, b := range m.books {

		if int(b.ID) == id {
			return b, nil
		}
	}

	return models.Book{}, errors.New("not found")
}

func (m *mockBookRepo) Create(
	book *models.Book,
) error {

	if m.err != nil {
		return m.err
	}

	book.ID = 1

	m.books = append(
		m.books,
		*book,
	)

	return nil
}

func (m *mockBookRepo) Update(
	_ *models.Book,
) error {

	if m.err != nil {
		return m.err
	}

	return nil
}

func (m *mockBookRepo) Delete(
	_ int,
) error {

	if m.err != nil {
		return m.err
	}

	return nil
}

func TestBooksGET(t *testing.T) {

	repo :=
		&mockBookRepo{
			books: []models.Book{
				{
					ID:     1,
					Title:  "Libro test",
					Author: "Autor",
					ISBN:   "123",
					Image:  "img",
				},
			},
		}

	handler :=
		BookHandler{
			Repo: repo,
		}

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Books(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"código esperado %d recibido %d",
			http.StatusOK,
			rec.Code,
		)
	}

}

func TestBooksPOST(t *testing.T) {

	repo :=
		&mockBookRepo{}

	handler :=
		BookHandler{
			Repo: repo,
		}

	body :=
		`
{
"title":"Nuevo",
"author":"Autor",
"ISBN":"999",
"image":"img"
}
`

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/books",
			bytes.NewBufferString(body),
		)

	rec :=
		httptest.NewRecorder()

	handler.Books(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"POST fallo %d",
			rec.Code,
		)
	}

}

func TestBookGETByID(t *testing.T) {

	repo :=
		&mockBookRepo{
			books: []models.Book{
				{
					ID:    1,
					Title: "Buscar",
				},
			},
		}

	handler :=
		BookHandler{
			Repo: repo,
		}

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books/1",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Book(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"GET ID fallo %d",
			rec.Code,
		)
	}

}

func TestBookPUT(t *testing.T) {

	repo :=
		&mockBookRepo{}

	handler :=
		BookHandler{
			Repo: repo,
		}

	body :=
		`
{
"title":"Editado",
"author":"Autor"
}
`

	req :=
		httptest.NewRequest(
			http.MethodPut,
			"/api/books/1",
			bytes.NewBufferString(body),
		)

	rec :=
		httptest.NewRecorder()

	handler.Book(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"PUT fallo %d",
			rec.Code,
		)
	}

}

func TestBookDELETE(t *testing.T) {

	repo :=
		&mockBookRepo{}

	handler :=
		BookHandler{
			Repo: repo,
		}

	req :=
		httptest.NewRequest(
			http.MethodDelete,
			"/api/books/1",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Book(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"DELETE fallo %d",
			rec.Code,
		)
	}

}

func TestBookInvalidID(t *testing.T) {

	handler :=
		BookHandler{
			Repo: &mockBookRepo{},
		}

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books/hola",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Book(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"esperaba 400 recibió %d",
			rec.Code,
		)
	}

}

func TestBooksMethodNotAllowed(t *testing.T) {

	handler :=
		BookHandler{
			Repo: &mockBookRepo{},
		}

	req :=
		httptest.NewRequest(
			http.MethodDelete,
			"/api/books",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Books(
		rec,
		req,
	)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"esperaba 405 recibió %d",
			rec.Code,
		)
	}

}

func TestBookRepositoryError(t *testing.T) {

	repo :=
		&mockBookRepo{
			err: errors.New("db error"),
		}

	handler :=
		BookHandler{
			Repo: repo,
		}

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	handler.Books(
		rec,
		req,
	)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf(
			"esperaba error 500",
		)
	}

}
