package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hhtvuyvt/proyecto-go/internal/db"
	"github.com/hhtvuyvt/proyecto-go/models"
)

func createTestRouter(t *testing.T) http.Handler {

	t.Helper()

	testDB := db.Open(":memory:")

	repo := models.BookRepository{
		DB: testDB,
	}

	return Router(
		RouterConfig{
			BookRepo: repo,
			JWTKey:   []byte("test-secret"),
		},
	)
}

func TestRouterRedirect(t *testing.T) {

	router := createTestRouter(t)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	if rec.Code != http.StatusTemporaryRedirect {
		t.Fatalf(
			"esperaba redirect %d, obtuvo %d",
			http.StatusTemporaryRedirect,
			rec.Code,
		)
	}

}

func TestRouterStatic(t *testing.T) {

	router := createTestRouter(t)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/static/index.html",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	if rec.Code == http.StatusNotFound {
		t.Fatal(
			"ruta static no configurada",
		)
	}

}

func TestRouterLoginRoute(t *testing.T) {

	router := createTestRouter(t)

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/login",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	if rec.Code == http.StatusNotFound {
		t.Fatal(
			"ruta /api/login no existe",
		)
	}

}

func TestRouterBooksPublicRoute(t *testing.T) {

	router := createTestRouter(t)

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/books",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	if rec.Code == http.StatusNotFound {
		t.Fatal(
			"ruta /api/books no existe",
		)
	}

}

func TestRouterProtectedBooks(t *testing.T) {

	router := createTestRouter(t)

	req :=
		httptest.NewRequest(
			http.MethodDelete,
			"/api/books/1",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	// Sin JWT debe bloquear
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf(
			"esperaba 401 sin token, obtuvo %d",
			rec.Code,
		)
	}

}
