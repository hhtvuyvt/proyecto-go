package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/models"
)

type mockBookRepo struct {
	books []models.Book
}

func (m *mockBookRepo) GetAll() ([]models.Book, error) {

	return m.books, nil

}

func (m *mockBookRepo) GetByID(
	id int,
) (models.Book, error) {

	for _, b := range m.books {

		if int(b.ID) == id {

			return b, nil

		}

	}

	return models.Book{}, errors.New("no encontrado")

}

func (m *mockBookRepo) Create(
	book *models.Book,
) error {

	book.ID =
		int64(len(m.books) + 1)

	m.books =
		append(
			m.books,
			*book,
		)

	return nil

}

func (m *mockBookRepo) Update(
	book *models.Book,
) error {

	for i, b := range m.books {

		if b.ID == book.ID {

			m.books[i] = *book

			return nil

		}

	}

	return errors.New("no encontrado")

}

func (m *mockBookRepo) Delete(
	id int,
) error {

	for i, b := range m.books {

		if int(b.ID) == id {

			m.books =
				append(
					m.books[:i],
					m.books[i+1:]...,
				)

			return nil

		}

	}

	return errors.New("no encontrado")

}

func TestBooksGET(t *testing.T) {

	repo :=
		&mockBookRepo{
			books: []models.Book{
				{
					ID:     1,
					Title:  "Libro prueba",
					Author: "Autor",
					ISBN:   "123",
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
			"esperado 200 recibido %d",
			rec.Code,
		)

	}

	var books []models.Book

	err :=
		json.NewDecoder(
			rec.Body,
		).Decode(&books)

	if err != nil {

		t.Fatal(err)

	}

	if len(books) != 1 {

		t.Fatalf(
			"esperaba 1 libro recibió %d",
			len(books),
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
		`{
			"title":"Nuevo",
			"author":"Autor",
			"ISBN":"999"
		}`

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
			"esperado 200 recibido %d",
			rec.Code,
		)

	}

	if len(repo.books) != 1 {

		t.Fatal(
			"el libro no fue creado",
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
			"esperado 200 recibido %d",
			rec.Code,
		)

	}

}

func TestBookPUT(t *testing.T) {

	repo :=
		&mockBookRepo{
			books: []models.Book{
				{
					ID:    1,
					Title: "Viejo",
				},
			},
		}

	handler :=
		BookHandler{
			Repo: repo,
		}

	body :=
		`{
			"title":"Nuevo",
			"author":"Autor",
			"ISBN":"555"
		}`

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
			"esperado 200 recibido %d",
			rec.Code,
		)

	}

	if repo.books[0].Title != "Nuevo" {

		t.Fatal(
			"no actualizo",
		)

	}

}

func TestBookDELETE(t *testing.T) {

	repo :=
		&mockBookRepo{
			books: []models.Book{
				{
					ID:    1,
					Title: "Eliminar",
				},
			},
		}

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
			"esperado 200 recibido %d",
			rec.Code,
		)

	}

	if len(repo.books) != 0 {

		t.Fatal(
			"no elimino",
		)

	}

}
